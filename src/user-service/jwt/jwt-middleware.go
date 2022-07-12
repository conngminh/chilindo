package token

import (
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

var jwtKey = []byte("supersecretkey")

type IJwtMiddleware interface {
	GenerateJWT(email string, username string, id int) (tokenString string, err error)
	ExtractToken(tokenString string) *JWTClaim
	ValidateToken(string string) (int, error)
}
type JWTClaim struct {
	Username string
	Email    string
	Id       int
	jwt.StandardClaims
}

func (j *JWTClaim) GenerateJWT(email string, username string, id int) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Id:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func (j *JWTClaim) ExtractToken(tokenString string) *JWTClaim {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil
	}
	return claims
}
func (j *JWTClaim) ValidateToken(signedToken string) (int, error) {
	claims := j.ExtractToken(signedToken)
	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Println("Token is expire")
		return 0, nil
	}
	return claims.Id, nil
}
