package router

import (
	"github.com/defeng-hub/Go-Storage/controller"
	"github.com/defeng-hub/Go-Storage/middleware"
	"github.com/gin-gonic/gin"
)

func setWebRouter(router *gin.Engine) {
	// Always available
	router.GET("/", middleware.ShowIndexPermissionCheck(), controller.GetIndexPage)
	router.GET("/login", controller.GetLoginPage)
	router.GET("/image", controller.GetImagePage)
	router.GET("/video", controller.GetVideoPage)
	router.GET("/help", controller.GetHelpPage)
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)

	// Download files
	fileDownloadAuth := router.Group("/")
	{
		fileDownloadAuth.GET("/upload/:file", controller.DownloadFile)
	}

	basicAuth := router.Group("/")
	basicAuth.Use(middleware.WebAuth())
	{
		basicAuth.GET("/manage", controller.GetManagePage)
	}
}
