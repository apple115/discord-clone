package api

import (
	"discord-clone/models"
	"discord-clone/pkg/app"
	"discord-clone/pkg/e"
	"discord-clone/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Username string `json:"username" validate:"required,min=3,max=50"`  // 用户名，必填，最小长度3，最大长度50
	Email    string `json:"email" validate:"required,email"`            // 邮箱，必填，必须是有效的邮箱格式
	Password string `json:"password" validate:"required,min=6,max=100"` // 密码，必填，最小长度6，最大长度100
}

// 注册请求
func Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form RegisterForm
	)
	httpCode, errCode := app.BindAndValidate(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//TODO密码强度检查

	// 检查用户名唯一性
	exist, err := models.ExistUsername(form.Username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_USER, nil)
		return
	}
	if exist {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USER, nil)
		return
	}

	//检查邮箱唯一性
	exist, err = models.ExistEmail(form.Email)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_EMAIL, nil)
		return
	}
	if exist {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_EMAIL, nil)
		return
	}

	//生成哈希密码,存储用户
	passwrordhash, err := util.HashPassword(form.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_HASHPASSWORD, nil)
		return
	}
	data := map[string]interface{}{
		"username":     form.Username,
		"email":        form.Email,
		"passwordhash": passwrordhash,
	}
	err = models.AddUser(data)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_USER, nil)
		return
	}

	//TODO 注册成功通知
	//返回结果
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
