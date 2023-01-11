package router

import (
	"github.com/defeng-hub/Go-Storage/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	setWebRouter(router)
	setApiRouter(router)
	router.NoRoute(controller.Get404Page)
}
