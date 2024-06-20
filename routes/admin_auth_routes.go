package routes

import (
	"apcore/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	auth := router.Group("/admin")
	auth.POST("/signin", controllers.AdminLogin)
}
