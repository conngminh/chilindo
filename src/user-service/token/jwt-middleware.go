package token

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtKey = []byte("superSecretKey")

type IJwtMiddleware interface {
	GenerateJWT(email string, uid uint) (tokenString string, err error)
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

func ExtractedToken(signedToken string) *Claims {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil
	}
	return claims
}

//func ValidateToken(signedToken string) (uint, error) {
//	claims := ExtractedToken(signedToken)
//	if claims.ExpiresAt < time.Now().Local().Unix() {
//		log.Println("Token is expire")
//		return 0, nil
//	}
//	return claims.UserId, nil
//}
