package handler

import (
	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	response.Success(c, "pong")
}
