package server

import (
	"context"
	"errors"
	"fmt"
	"gin-quickly-template/config"
	"gin-quickly-template/internal/app/ping"
	"gin-quickly-template/internal/core/database/mysql"
	"gin-quickly-template/internal/core/database/pgsql"
	"gin-quickly-template/internal/core/database/redis"
	"gin-quickly-template/internal/core/kernel"
	"gin-quickly-template/internal/core/logx"
	"gin-quickly-template/pkg/colorful"
	"gin-quickly-template/pkg/ip"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var log = logx.NameSpace("cmd.server")

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "server",
		Short:   "Set Application config info",
		Example: "main server -c ./config.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			// init kernel
			fmt.Println(colorful.Yellow("init kernel..."))
			k := kernel.New()
			k.Ctx, k.CancelFunc = ctx, cancel

			msl := mysql.InitMysql()
			if msl == nil {
				panic(errors.New("mysql init failed"))
			}
			k.Injector.Map(&msl)

			pg := pgsql.InitPostgres()
			if pg == nil {
				panic(errors.New("postgres init failed"))
			}
			k.Injector.Map(&pg)

			rds := redis.InitCache()
			if rds == nil {
				panic(errors.New("redis init failed"))
			}
			k.Injector.Map(&rds)

			// init sentry
			//if config.GetConfig().Sentry.Enable {
			//fmt.Println(colorful.Yellow("init sentry..."))
			//sentryx.NewSentry()
			//}

			// init tracer
			//if config.GetConfig().OTel.Enable {
			//fmt.Println(colorful.Yellow("init tracer..."))
			//tracex.Init()
			//}

			k.Gin = gin.New()
			k.Gin.Use(gin.Recovery(),
				gin.Logger(),
				cors.Default(),
				//sentryx.SentryMiddleware()
				//tracer.Trace(),
			)
			k.Injector.Map(&k.Gin)

			// !!!IMPORTANT: register module here
			k.RegMod(&ping.Ping{})

			err := k.StartModule()
			if err != nil {
				panic(err)
			}

			port := config.GetConfig().Port
			k.HttpServer = &http.Server{
				Addr:    ":" + port,
				Handler: k.Gin,
			}

			// start http server
			go func() {
				if err := k.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					fmt.Printf(colorful.Green("listen: %s\n"), port)
					fmt.Printf(colorful.Yellow("Server run failed: %s\n"), err)
				}
			}()
			println(colorful.Green("Server run at:"))
			println(fmt.Sprintf("-  Local:   http://localhost:%s", port))

			// store local ip
			localHosts := ip.GetLocalHost()
			k.CurrentIpList = make([]string, 0, len(localHosts))
			for _, host := range localHosts {
				k.CurrentIpList = append(k.CurrentIpList, host)
				println(fmt.Sprintf("-  Network: http://%s:%s", host, port))
			}

			// graceful shutdown
			quit := make(chan os.Signal)
			signal.Notify(quit, os.Interrupt)
			<-quit
			println(colorful.Blue("Shutting down server..."))

			err = k.Stop()
			if err != nil {
				panic(err)
			}

			// cancelFunc ...
			newCtx, cancelFunc := context.WithTimeout(k.Ctx, 5*time.Second)
			defer cancelFunc()
			defer k.CancelFunc()

			// shutdown http server
			if err := k.HttpServer.Shutdown(newCtx); err != nil {
				fmt.Println(colorful.Yellow("Server forced to shutdown: " + err.Error()))
			}

			fmt.Println(colorful.Green("Server exiting Correctly"))
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "", "config file path")
}

func setup() {
	// init config
	fmt.Println(colorful.Yellow("init config..."))
	config.LoadConfig(configYml)

	// init logx
	fmt.Println(colorful.Yellow("init logx..."))
	if config.GetConfig().MODE == "" || config.GetConfig().MODE == "debug" {
		logx.Init(zapcore.DebugLevel)
	} else {
		logx.Init(zapcore.InfoLevel)
	}
	defer func() {
		if err := recover(); err != nil {
			_ = log.Sync()
		}
	}()

	// init gin
	fmt.Println(colorful.Yellow("init gin..."))
	if config.GetConfig().MODE == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

}
