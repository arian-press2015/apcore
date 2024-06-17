package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	PingRoutes(router)
	AuthRoutes(router)
	UsersRoutes(router)
	RolesRoutes(router)
}
