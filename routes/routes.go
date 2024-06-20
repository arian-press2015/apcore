package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	PingRoutes(router)
	AuthRoutes(router)
	UsersRoutes(router, db)
	RolesRoutes(router)
	AdminRoutes(router)
	SwaggerRoutes(router)
}
