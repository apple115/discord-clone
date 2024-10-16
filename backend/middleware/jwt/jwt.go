package jwt

import (
	"discord-clone/pkg/e"
	"discord-clone/pkg/util"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		accesstoken := c.Query("accesstoken")
		if accesstoken == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseAccessToken(accesstoken)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_VERIFY_ACCESS_TOKEN
				}
			} else {
				c.Set("user_id", claims.UserId)
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
