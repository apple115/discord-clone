package v1

import (
	"discord-clone/models"
	"discord-clone/pkg/app"
	"discord-clone/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateChannelForm struct {
	UserID      string `json:"userID" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// POST /api/channels
// @Router /api/channels
// 创建某个频道
func CreateChannel(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form CreateChannelForm
	)
	httpCode, errCode := app.BindAndValidate(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//如果频道名存在
	exist, err := models.ExitChannelByName(form.Name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_CHANNEL_NAME_FAIL, nil)
		return
	}
	if exist {
		appG.Response(http.StatusOK, e.ERROR_EXIST_CHANNEL_NAME, nil)
		return
	}

	data := map[string]interface{}{
		"name":        form.Name,
		"description": form.Description,
		"userID":      form.UserID,
	}
	err = models.AddChannel(data)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CHANNEL_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// DELETE /api/channels/{channelID}
// @Router /api/channels/{channelID}
// 删除某个频道
func DeleteChannel(c *gin.Context) {
	appG := app.Gin{C: c}
	channelID := c.Param("channelID")

	err := models.DeleteChannel(channelID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_CHANNEL_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// Get
// @Router /channels/{channelID}/messages
// 得到某个频道的所有消息
func GetChannelMessageByID(c *gin.Context) {
	appG := app.Gin{C: c}
	channelID := c.Param("channelID")

	// 从 MongoDB 获取消息
	data, err := models.GetChannelMessages(channelID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CHANNEL_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// 获取所有频道
// @Router /api/channels
func GetChannels(c *gin.Context) {
	appG := app.Gin{C: c}
	data, err := models.GetChannel()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CHANNEL_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// 获取特定频道的详细信息
// @Router /api/channels/{channelID}
func GetChannelByID(c *gin.Context) {
	appG := app.Gin{C: c}
	channelID := c.Param("channelID")

	data, err := models.GetChannelByID(channelID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CHANNEL_BY_ID_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// PUT /api/channels/{channelID}
// @Router /api/channels/{channelID}
// 更新频道信息
func UpdateChannel(c *gin.Context) {
	var (
		form  CreateChannelForm
		appG  = app.Gin{C: c}
		err   Error
		exist bool
	)
	channelID := c.Param("channelID")
	httpCode, errCode := app.BindAndValidate(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//是否存在这个频道
	exist, err = models.ExitChannel(channelID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CHANNEL_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CHANNEL, nil)
		return
	}

	//如果更新的频道名存在，返回错误
	exist, err = models.ExitChannelByName(form.Name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_CHANNEL_NAME_FAIL, nil)
		return
	}
	if exist {
		appG.Response(http.StatusOK, e.ERROR_EXIST_CHANNEL_NAME, nil)
		return
	}

	data := map[string]interface{}{
		"name":        form.Name,
		"description": form.Description,
	}
	// 更新频道信息
	err = models.EditChannel(channelID, data)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_CHANNEL_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// DELETE /api/channels/{channelID}
// @Router /api/channels/{channelID}
// 删除频道
func DeleteChannelByID(c *gin.Context) {
	channelID := c.Param("channelID")
	appG := app.Gin{C: c}

	//是否存在这个频道
	exist, err = models.ExitChannel(channelID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CHANNEL_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CHANNEL, nil)
		return
	}

	// 删除频道
	err := models.DeleteChannel(channelID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_CHANNEL_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
