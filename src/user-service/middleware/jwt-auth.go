package middleware

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

//AuthorizeJWT validates thr token user given, return 401 if not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		token := strings.TrimPrefix(tokenString, "Bearer ")
		useId, err := jwt.ValidateToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			log.Println("Check ValidateToken", err)
			return
		}
		context.Set(config.UserId, useId)
		context.Next()
	}
}
