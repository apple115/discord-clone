package api

import (
	"discord-clone/models"
	"discord-clone/pkg/app"
	"discord-clone/pkg/e"
	auth_services "discord-clone/service/auth_service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type RefreshTokenForm struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=6,max=256"`
}

func GetAuth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AuthForm
	)
	httpCode, errCode := app.BindAndValidate(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	a := auth_services.Auth{Email: form.Email, Password: form.Password}
	//验证用户信息
	check, err := a.Check()
	if !check {
		appG.Response(httpCode, e.ERROR_CHECK_EXIST_USER, nil)
		return
	}
	if err != nil {
		appG.Response(httpCode, e.ERROR_CHECK_EXIST_USER, nil)
		return
	}
	//读取用户信息
	UserPublic, err := a.GetUserPublic()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER, nil)
		return
	}
	//生成token
	accesstoken, RefreshToken, err := auth_services.GenerateToken(UserPublic.ID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GEN_TOKEN, nil)
		return
	}
	//返回token
	appG.Response(http.StatusOK, e.SUCCESS, models.AuthResponse{AccessToken: accesstoken, RefreshToken: RefreshToken})
}

// 刷新access token
func RefreshToken(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form RefreshTokenForm
	)
	httpCode, errCode := app.BindAndValidate(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//验证refresh token
	userID, err := auth_services.VerifyRefreshToken(form.RefreshToken)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_VERIFY_RFRESH_TOKEN, err)
		return
	}
	log.Println(userID)
	//生成token
	AccessToken, RefreshToken, err := auth_services.GenerateToken(userID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GEN_TOKEN, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, models.AuthResponse{AccessToken: AccessToken, RefreshToken: RefreshToken})
}
