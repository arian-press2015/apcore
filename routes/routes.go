package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	PingRoutes(router)
	AuthRoutes(router, db)
	UsersRoutes(router, db)
	RolesRoutes(router, db)
	AdminAuthRoutes(router, db)
	SwaggerRoutes(router)
}
