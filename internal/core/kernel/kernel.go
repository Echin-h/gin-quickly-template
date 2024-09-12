package kernel

import (
	"context"
	"gin-quickly-template/internal/core/logx"
	"github.com/gin-gonic/gin"
	"github.com/juanjiTech/inject/v2"
	"net/http"
	"sync"
)

const (
	MYSQL = "mysql"
	REDIS = "redis"
	PGSQL = "pgsql"
)

type Engine struct {
	Gin *gin.Engine

	Ctx        context.Context
	CancelFunc context.CancelFunc

	inject.Injector

	HttpServer *http.Server

	CurrentIpList []string

	modules   map[string]Module
	modulesMu sync.Mutex
}

func New() *Engine {
	return &Engine{
		modules:       make(map[string]Module),
		CurrentIpList: make([]string, 0),
		Injector:      inject.New(),
	}
}

func (e *Engine) StartModule() error {
	hub := Hub{
		Injector: e.Injector,
	}

	for _, mod := range e.modules {
		h4m := hub
		h4m.Log = logx.NameSpace("module." + mod.Name())
		if err := mod.PreInit(&h4m); err != nil {
			h4m.Log.Error(err, "module.PreInit", mod.Name())
			panic(err)
		}
	}

	for _, mod := range e.modules {
		h4m := hub
		h4m.Log = logx.NameSpace("module." + mod.Name())
		if err := mod.Init(&h4m); err != nil {
			h4m.Log.Error(err, "module.Init", mod.Name())
			panic(err)
		}
	}

	for _, mod := range e.modules {
		h4m := hub
		h4m.Log = logx.NameSpace("module." + mod.Name())
		if err := mod.PostInit(&h4m); err != nil {
			h4m.Log.Error(err, "module.PostInit", mod.Name())
			panic(err)
		}
	}

	for _, mod := range e.modules {
		h4m := hub
		h4m.Log = logx.NameSpace("module." + mod.Name())
		if err := mod.Load(&h4m); err != nil {
			h4m.Log.Error(err, "module.Load", mod.Name())
			panic(err)
		}
	}

	for _, mod := range e.modules {
		h4m := hub
		h4m.Log = logx.NameSpace("module." + mod.Name())
		if err := mod.Start(&h4m); err != nil {
			h4m.Log.Error(err, "module.Start", mod.Name())
			panic(err)
		}
	}

	return nil
}

func (e *Engine) Stop() error {
	wg := sync.WaitGroup{}
	wg.Add(len(e.modules))
	for _, mod := range e.modules {
		err := mod.Stop(&wg, e.Ctx)
		if err != nil {
			return err
		}
	}
	wg.Wait()

	return nil
}

func (e *Engine) RegMod(mods ...Module) {
	e.modulesMu.Lock()
	defer e.modulesMu.Unlock()
	for _, mod := range mods {
		if mod.Name() == "" {
			panic("name of module can't be empty")
		}
		if _, ok := e.modules[mod.Name()]; ok {
			panic("module " + mod.Name() + " already exists")
		}
		e.modules[mod.Name()] = mod
	}
}

func (e *Engine) Serve() {}
