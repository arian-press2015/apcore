package routes

import (
	"apcore/controllers"

	"github.com/gin-gonic/gin"
)

func PingRoutes(router *gin.Engine, ctrl *controllers.PingController) {
	router.GET("/ping", ctrl.Ping)
}
