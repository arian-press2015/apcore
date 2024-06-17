package routes

import (
	"apcore/controllers"

	"github.com/gin-gonic/gin"
)

func PingRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}
