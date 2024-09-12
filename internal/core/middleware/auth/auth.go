package auth

import (
	"gin-quickly-template/internal/core/middleware/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" || strings.Index(token, "Bearer") != 0 {
			response.UnAuthorization(c)
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		entity, err := ParseToken(token)
		if err != nil {
			response.UnAuthorization(c)
			return
		}
		c.Set("user", entity.Info)
	}
}
