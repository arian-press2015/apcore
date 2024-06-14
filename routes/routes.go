package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	HealthCheckRoutes(router)
	AuthRoutes(router)
	UsersRoutes(router)
	RolesRoutes(router)
	AdminRoutes(router)
}
