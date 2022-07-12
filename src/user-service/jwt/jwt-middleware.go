package jwt

import (
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

var jwtKey = []byte("superSecretKey")

type JWTClaims struct {
	UserName string
	UserID   int
	jwt.StandardClaims
}

func GenerateJWT(username string, userid int) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaims{
		UserName: username,
		UserID:   userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ExtractedToken(signedToken string) *JWTClaims {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil
	}
	return claims
}

func ValidateToken(signedToken string) (int, error) {
	claims := ExtractedToken(signedToken)
	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Println("Token is expire")
		return 0, nil
	}
	return claims.UserID, nil
}
