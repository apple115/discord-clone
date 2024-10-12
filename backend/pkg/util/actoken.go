package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

func Setup() {
	jwtSecret = []byte("go-discord-clone")
}

type AccessClaims struct {
	UserId uint `json:"userId"`
	jwt.StandardClaims
}

func GenerateAccessToken(userId uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * time.Minute)

	claims := AccessClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-discord-clone",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseAccessToken(token string) (*AccessClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AccessClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
