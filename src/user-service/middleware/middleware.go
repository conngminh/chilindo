package middleware

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type IMiddleWare interface {
	MiddleWare() gin.HandlerFunc
}

type SMiddleWare struct {
	tokenController *token.Claims
}

func (s *SMiddleWare) MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSONP(http.StatusUnauthorized, gin.H{
				"Message": "Request doest not contain token",
			})
			log.Println("MiddleWare: Error to get token in")
			c.Abort()
			return
		}
		tokenResult := strings.TrimPrefix(tokenString, "Bearer ")
		claim := token.ExtractedToken(tokenResult)
		c.Set(config.UserId, claim.Id)
		if claim.ExpiresAt < time.Now().Local().Unix() {
			c.JSONP(http.StatusUnauthorized, gin.H{
				"Message": "Token is Expired",
			})
			log.Println("MiddleWare: Error token is Expired")
			c.Abort()
			return
		}
		c.Next()
	}
}
