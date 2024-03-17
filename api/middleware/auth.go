package middleware

import (
	"github.com/gin-gonic/gin"
	"go-api-server/internal/util"
	"log"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("ACCESS_TOKEN")
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		tokenDetail, err := util.ParseJwtToken(accessToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Println(tokenDetail)
		c.Next()
	}
}
