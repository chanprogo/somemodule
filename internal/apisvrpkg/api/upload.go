package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/chanprogo/somemodule/app"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/common"
	"github.com/chanprogo/somemodule/pkg/conf/iconf"
	"github.com/chanprogo/somemodule/pkg/constant"
	"github.com/chanprogo/somemodule/pkg/log/logging"
	filepkg "github.com/chanprogo/somemodule/pkg/util/file"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, constant.RESPONSE_CODE_SYSTEM, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}

	imageName := filepkg.GetImageName(image.Filename)

	fullPath := common.GetImageFullPath()
	savePath := common.GetImagePath()

	src := fullPath + imageName

	if !filepkg.CheckImageExt(imageName, iconf.AppSetting.ImageAllowExts) || !filepkg.CheckImageSize(file, iconf.AppSetting.ImageMaxSize) {
		appG.Response(http.StatusBadRequest, constant.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = filepkg.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, constant.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, constant.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constant.RESPONSE_CODE_OK, map[string]string{
		"image_url":      common.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
