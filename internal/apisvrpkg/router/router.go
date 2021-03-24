package router

import (
	"net/http"

	"github.com/chanprogo/somemodule/internal/apisvrpkg/api"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/common"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/controller"
	"github.com/chanprogo/somemodule/internal/apisvrpkg/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.StaticFS("/export", http.Dir(service.GetExcelFullPath()))

	r.StaticFS("/upload/images", http.Dir(common.GetImageFullPath()))
	r.POST("/upload", api.UploadImage)

	new(controller.TagController).Router(r)

}
