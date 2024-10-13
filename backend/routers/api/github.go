package api

import (
	"context"
	"discord-clone/models"
	"discord-clone/pkg/app"
	"discord-clone/pkg/e"
	"discord-clone/pkg/oauth2/github"
	auth_services "discord-clone/service/auth_service"
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

	username, _ := userInfo["login"].(string)
	email, _ := userInfo["email"].(string)

	//TODO: 如果没有这个 创建或更新用户记录
	exits, err := models.ExistEmail(email)
	if !exits {
		data := map[string]interface{}{
			"Username":     username,
			"PasswrodHash": nil,
			"Email":        email,
		}

		err = models.AddUser(data)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER, nil)
			return
		}
	}
	User, err := models.GetUserByEmail(email)
	//生成内部访问令牌和刷新令牌
	AccessToken, RefreshToken, err := auth_services.GenerateToken(User.ID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GEN_TOKEN, nil)
		return
	}
	//返回内部访问令牌和刷新令牌
	appG.Response(http.StatusOK, e.SUCCESS, models.AuthResponse{AccessToken: AccessToken, RefreshToken: RefreshToken})
}
