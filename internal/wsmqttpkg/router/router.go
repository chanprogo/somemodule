package router

import (
	"net/http"

	"github.com/chanprogo/somemodule/app/middleware"

	. "github.com/chanprogo/somemodule/internal/wsmqttpkg/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Static("/public", "./public")
	router.Use(middleware.ShowRequest())

	new(MelodyController).Router(router)

	user := router.Group("/user")
	user.GET("/upload", GetUpload)

}

func GetUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "/safe/upload.html", gin.H{
		"title": "下载页面",
	})
}
