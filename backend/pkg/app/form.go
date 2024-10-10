package app

import (
	"discord-clone/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(c *gin.Context, form interface{}) (int, int) {
	// 绑定请求数据到结构体
	if err := c.ShouldBind(form); err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	// 获取验证器实例
	validate := validator.New()

	// 验证结构体
	if err := validate.Struct(form); err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
