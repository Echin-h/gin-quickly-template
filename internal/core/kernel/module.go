package kernel

import (
	"context"
	"github.com/juanjiTech/inject/v2"
	"go.uber.org/zap"
	"sync"
)

type Hub struct {
	inject.Injector
	Log *zap.SugaredLogger
}

type Module interface {
	Name() string
	PreInit(*Hub) error
	Init(*Hub) error
	PostInit(*Hub) error
	Load(*Hub) error
	Start(*Hub) error
	Stop(wg *sync.WaitGroup, ctx context.Context) error
	mustEmbedUnimplementedModule()
}

// check if UnimplementedModule implements Module
var _ Module = (*UnimplementedModule)(nil)

// UnimplementedModule provides a default implementation for the Module interface
type UnimplementedModule struct{}

// Name should be defined
func (u *UnimplementedModule) Name() string { panic("name of module should be defined") }

func (u *UnimplementedModule) PreInit(*Hub) error { return nil }

func (u *UnimplementedModule) Init(*Hub) error { return nil }

func (u *UnimplementedModule) PostInit(*Hub) error { return nil }

func (u *UnimplementedModule) Load(*Hub) error { return nil }

func (u *UnimplementedModule) Start(*Hub) error { return nil }

func (u *UnimplementedModule) Stop(wg *sync.WaitGroup, _ context.Context) error {
	defer wg.Done()
	return nil
}

func (u *UnimplementedModule) mustEmbedUnimplementedModule() {}
