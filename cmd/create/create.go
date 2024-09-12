package create

import (
	"bytes"
	"errors"
	"gin-quickly-template/pkg/colorful"
	"gin-quickly-template/pkg/fs"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
	"text/template"
)

var (
	appName  string
	dir      string
	force    bool
	StartCmd = &cobra.Command{
		Use:     "create",
		Short:   "create a new app",
		Example: "app create -n users",
		Run: func(cmd *cobra.Command, args []string) {
			err := load()
			if err != nil {
				colorful.Red("load error:" + err.Error())
				os.Exit(1)
			}
			colorful.Yellow("Module " + appName + " generate success under " + dir)
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "", "create a new app with provided name")
	StartCmd.PersistentFlags().StringVarP(&dir, "path", "p", "internal/app", "new file will generate under provided path")
	StartCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Force generate the app")
}

func load() error {
	if appName == "" {
		return errors.New("app name should not be empty, use -n")
	}

	router := path.Join(dir, appName, "router")
	handlerMain := path.Join(dir, appName, "handler")
	dto := path.Join(dir, appName, "dto")
	dao := path.Join(dir, appName, "dao")
	service := path.Join(dir, appName, "service")
	model := path.Join(dir, appName, "model")
	orm := path.Join(dir, appName, "dao")
	cache := path.Join(dir, appName, "dao")
	init := path.Join(dir, appName)

	_ = fs.IsNotExistMkDir(init)
	_ = fs.IsNotExistMkDir(router)
	_ = fs.IsNotExistMkDir(handlerMain)
	_ = fs.IsNotExistMkDir(dto)
	_ = fs.IsNotExistMkDir(service)
	_ = fs.IsNotExistMkDir(model)
	_ = fs.IsNotExistMkDir(dao)

	m := map[string]string{}
	m["appName"] = strings.ToLower(appName[:1]) + appName[1:]
	m["AppName"] = strings.ToUpper(appName[:1]) + appName[1:]

	router += "/" + m["appName"] + ".go"
	service += "/" + "service.go"
	model += "/" + m["appName"] + ".go"
	handlerMain += "/" + m["appName"] + ".go"
	dto += "/" + m["appName"] + ".go"
	init += "/" + "init.go"
	dao += "/" + m["appName"] + ".go"
	orm += "/" + "orm.go"
	cache += "/" + "cache.go"

	if rt, err := template.ParseFiles("template/router.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, router)
	}

	if rt, err := template.ParseFiles("template/handler.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, handlerMain)
	}

	if rt, err := template.ParseFiles("template/dto.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, dto)
	}
	if rt, err := template.ParseFiles("template/service.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, service)
	}
	if rt, err := template.ParseFiles("template/model.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, model)
	}

	if rt, err := template.ParseFiles("template/init.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, init)
	}

	if rt, err := template.ParseFiles("template/dao.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, dao)
	}

	if rt, err := template.ParseFiles("template/orm.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, orm)
	}

	if rt, err := template.ParseFiles("template/cache.template"); err != nil {
		return err
	} else {
		var b bytes.Buffer
		err = rt.Execute(&b, m)
		fs.FileCreate(b, cache)
	}

	return nil
}
