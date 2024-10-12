package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RefreshClaims struct {
	UserId uint `json:"userId"`
	jwt.StandardClaims
}

func GenerateRefreshToken(userId uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 7 * time.Hour)
	claims := RefreshClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-discord-clone",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseRefreshToken(token string) (*RefreshClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*RefreshClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
