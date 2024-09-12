package ping

import (
	"context"
	"fmt"
	"gin-quickly-template/internal/app/ping/dao"
	"gin-quickly-template/internal/app/ping/router"
	"gin-quickly-template/internal/core/kernel"
	"gin-quickly-template/pkg/colorful"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"sync"
)

var _ kernel.Module = (*Ping)(nil)

type Ping struct {
	kernel.UnimplementedModule
}

func (u *Ping) Name() string {
	return "ping"
}

func (u *Ping) PreInit(hub *kernel.Hub) error {
	// pgsql as the same and replace
	var mysql *gorm.DB
	err := hub.Load(&mysql)
	if err != nil {
		return err
	}

	var rds *redis.Client
	err = hub.Load(&rds)
	if err != nil {
		return err
	}

	err = dao.Init(mysql, rds)
	if err != nil {
		return err
	}

	return nil
}

func (u *Ping) Init(hub *kernel.Hub) error { return nil }

func (u *Ping) PostInit(hub *kernel.Hub) error { return nil }

func (u *Ping) Load(hub *kernel.Hub) error {
	var ginE *gin.Engine
	err := hub.Load(&ginE)
	if err != nil {
		return err
	}
	router.InitRouter(ginE)
	return nil
}

func (u *Ping) Start(hub *kernel.Hub) error {
	fmt.Println(colorful.Yellow("module Ping start success ..."))
	return nil
}

func (u *Ping) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
