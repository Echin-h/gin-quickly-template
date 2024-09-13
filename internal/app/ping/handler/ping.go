package handler

import (
	"gin-quickly-template/internal/core/middleware/response"
	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	response.Success(c, "pong")
}
