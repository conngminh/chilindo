package token

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtKey = []byte("superSecretKey")

type IJwtMiddleware interface {
	GenerateJWT(email string, userid uint) (tokenString string, err error)
	ExtractToken(tokenString string) (*Claims, error)
}

type Claims struct {
	Email  string
	UserId uint
	jwt.StandardClaims
}

func (j *Claims) GenerateJWT(email string, userid uint) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email:  email,
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ExtractToken(signedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
