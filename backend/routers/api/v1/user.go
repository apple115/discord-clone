package v1

import (
	"discord-clone/models"
	"discord-clone/pkg/app"
	"discord-clone/pkg/e"
	"discord-clone/pkg/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadUserImage(c *gin.Context) {
	var appG = app.Gin{C: c}
	UserId, exist := c.Get("user_id")
	if !exist {
		appG.Response(http.StatusBadRequest, e.ERROR_CHECK_EXIST_USER, nil)
		return
	}
	//存在这个ID
	exist, err := models.ExistUserId(UserId.(uint))
	if !exist {
		appG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_USER, nil)
		return
	}
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}
	err = upload.CheckImage(fullPath)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	if err := c.SaveUploadedFile(image, src); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	image_url := upload.GetImageFullUrl(imageName)
	//mysql 存储图片url和用户关联
	data := map[string]interface{}{
		"ProfilePictureUrl": savePath + imageName,
	}
	err = models.AddUserPicture(UserId.(uint), data)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_USER, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      image_url,
		"image_save_url": savePath + imageName,
	})
}
