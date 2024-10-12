package auth_services

import (
	"discord-clone/pkg/gredis"
	"discord-clone/pkg/util"
	"fmt"
	"time"
)

// 返回双token
func GenerateToken(userId uint) (string, string, error) {
	AccessToken, err := util.GenerateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	RefreshToken, err := util.GenerateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}
	//TODO: 保存RefreshToken的MD5到redis
	key := "user:rftMD5:" + util.EncodeMD5(RefreshToken)
	time := 60 * 60 * 24 * 7 // 7天
	err = gredis.Set(key, "", time)
	if err != nil {
		return "", "", err
	}
	// 返回token
	return AccessToken, RefreshToken, nil
}

func GenerateAccessToken(userId uint) (string, error) {
	accesstoken, err := util.GenerateAccessToken(userId)
	if err != nil {
		return "", err
	}
	return accesstoken, nil
}

func VerifyRefreshToken(refreshToken string) (uint, error) {
	claims, err := util.ParseRefreshToken(refreshToken)
	if err != nil {
		return 0, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return 0, fmt.Errorf("refresh token is expired")
	}
	//验证redis中是否存在
	key := "user:rftMD5:" + util.EncodeMD5(refreshToken)
	_, err = gredis.Get(key)
	if err != nil {
		return 0, fmt.Errorf("refresh token is not exist redis")
	}
	//如果存在,删除redis中的refreshToken
	_, err = gredis.Delete(key)
	if err != nil {
		return 0, fmt.Errorf("refrshtoken delete error")
	}
	return claims.UserId, nil
}

func VerifyAccessToken(accessToken string) (uint, error) {
	claims, err := util.ParseAccessToken(accessToken)
	if err != nil {
		return 0, err
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return 0, fmt.Errorf("access token is expired")
	}
	return claims.UserId, nil
}
