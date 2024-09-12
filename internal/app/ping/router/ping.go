package router

import (
	"gin-quickly-template/internal/app/ping/handler"
	"github.com/gin-gonic/gin"
)

// Don't forget to add mod to kernel.RegisterModule

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", handler.GetPing)
	}
}
