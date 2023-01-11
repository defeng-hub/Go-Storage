package router

import (
	"github.com/defeng-hub/Go-Storage/controller"
	"github.com/defeng-hub/Go-Storage/middleware"
	"github.com/gin-gonic/gin"
)

func setWebRouter(router *gin.Engine) {
	router.Use(middleware.GlobalWebRateLimit())
	// Always available
	router.GET("/", controller.GetIndexPage)
	router.GET("/login", controller.GetLoginPage)
	router.GET("/image", controller.GetImagePage)
	router.GET("/video", controller.GetVideoPage)
	router.GET("/help", controller.GetHelpPage)
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)

	// Download files
	fileDownloadAuth := router.Group("/")
	fileDownloadAuth.Use(middleware.FileDownloadPermissionCheck())
	{
		fileDownloadAuth.GET("/upload/:file", controller.DownloadFile)
		//fileDownloadAuth.GET("/upload/images/:file", controller.DownloadFile)
		//fileDownloadAuth.GET("/explorer", controller.GetExplorerPageOrFile)
	}

	basicAuth := router.Group("/")
	basicAuth.Use(middleware.WebAuth())
	{
		basicAuth.GET("/manage", controller.GetManagePage)
	}
}
