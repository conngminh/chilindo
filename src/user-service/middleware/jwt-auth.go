package middleware

import (
	"chilindo/src/user-service/service"
	"chilindo/src/user-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
)

//AuthorizeJWT validates thr token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenSigned := strings.TrimPrefix(authHeader, "bearer ")
		token, err := jwtService.ValidateToken(tokenSigned)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]:", claims["user_id"])
			log.Println("Claim[issuer]:", claims["issuer"])
		} else {
			log.Println(err)
			response := utils.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
