package sentryx

import (
	"gin-quickly-template/config"
	"gin-quickly-template/internal/core/logx"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"time"
)

// TODO: this file isn't tested yet, so just do it

func NewSentry() {
	if !config.GetConfig().Sentry.Enable || config.GetConfig().Sentry.Dsn == "" {
		return
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.GetConfig().Sentry.Dsn,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		logx.NameSpace("sentryx").Error(err)
	}
}

// SentryMiddleware is a middleware specified for gin
func SentryMiddleware() gin.HandlerFunc {
	return sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: true,
		Timeout:         2 * time.Second,
	})
}

func Capture(ctx *gin.Context, key string, value string, msg ...string) {
	hub := sentrygin.GetHubFromContext(ctx)
	if hub == nil {
		sentry.CurrentHub().Clone()
		sentry.SetHubOnContext(ctx, hub)
	}
	hub.WithScope(func(scope *sentry.Scope) {
		scope.SetExtra(key, value)
		hub.CaptureMessage(msg[0])
	})
}
