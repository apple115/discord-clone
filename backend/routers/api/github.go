package api

import (
	"context"
	"discord-clone/pkg/app"
	"discord-clone/pkg/e"
	"discord-clone/pkg/oauth2/github"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// 重定位
func GitHubLogin(c *gin.Context) {
	oauthConfig := github.GetClient()
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// 重定向会应用程序
func GitHubCallback(c *gin.Context) {
	code := c.Query("code")
	appG := app.Gin{C: c}
	oauthConfig := github.GetClient()
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	// 使用访问令牌获取用户信息
	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	// log.Println(userInfo)
	//TODO: 创建或更新用户记录

	//TODO: 生成内部访问令牌和刷新令牌

	//TODO: 存储刷新令牌到 Redis

	//TODO: 返回内部访问令牌和刷新令牌
	username, _ := userInfo["login"].(string)
	email, _ := userInfo["email"].(string)
	data := map[string]interface{}{
		"Username": username,
		"email":    email,
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)

}
