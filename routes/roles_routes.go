package routes

import (
	"apcore/controllers"
	"apcore/middlewares"
	"apcore/repositories"
	"apcore/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RolesRoutes(router *gin.Engine, db *gorm.DB) {
	roleRepository := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepository)
	roleController := controllers.NewRoleController(roleService)

	roles := router.Group("/")
	roles.Use(middlewares.JWTAuthMiddleware())
	roles.POST("/roles", roleController.CreateRole)
	roles.GET("/roles", roleController.GetRoles)
}
