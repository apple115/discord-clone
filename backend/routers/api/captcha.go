package api

import (
	"discord-clone/pkg/app"
	"discord-clone/pkg/captdata"
	"discord-clone/pkg/e"
	"discord-clone/pkg/gredis"
	"discord-clone/pkg/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wenlng/go-captcha/v2/click"
)

func GetClickBasicCaptData(c *gin.Context) {
	var capt click.Captcha
	if c.Query("type") == "light" {
		capt = captdata.LightTextCapt

	} else {
		capt = captdata.TextCapt
	}

	captData, err := capt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	dotData := captData.GetData() // 获取点阵数据
	if dotData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "gen captcha data failed",
		})
		return
	}

	var masterImageBase64, thumbImageBase64 string

	masterImageBase64 = captData.GetMasterImage().ToBase64()
	thumbImageBase64 = captData.GetThumbImage().ToBase64()

	dotsByte, _ := json.Marshal(dotData)
	// 保存点阵数据，五分钟
	fmt.Println("dot>>>>", string(dotsByte))
	key := util.StringToMD5(string(dotsByte))
	gredis.Set(key, dotData, 60*50)
	c.JSON(http.StatusOK, gin.H{
		"code":         0,
		"captcha_key":  key,
		"image_base64": masterImageBase64,
		"thumb_base64": thumbImageBase64,
	})
}

// VerifyCaptcha .
func VerifyCaptcha(c *gin.Context) {
	var appG = app.Gin{C: c}
	captchaKey := c.PostForm("captcha_key")
	userDots := c.PostForm("user_dots")

	// 从缓存中读取验证码数据
	storedDots, err := gredis.Get(captchaKey)
		log.Println(string(storedDots))

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CAPTCHA_KEY_EMPTY, nil)
		return
	}

	// 解析存储的验证码数据
	var storedDotsMap map[int]*click.Dot
	err = json.Unmarshal(storedDots, &storedDotsMap)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CAPTCHA_STORE, err.Error())
		return
	}

	// 解析用户提交的坐标数据
	dots := strings.Split(userDots, ",")

	// 验证用户提交的坐标是否与存储的验证码数据匹配
	chkRet := captdata.VerifyDots(dots, storedDotsMap)
	if chkRet {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	} else {
		appG.Response(http.StatusOK, e.ERROR, nil)
	}
}
