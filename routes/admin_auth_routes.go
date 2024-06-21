package routes

import (
	"apcore/controllers"
	"apcore/repositories"
	"apcore/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminAuthRoutes(router *gin.Engine, db *gorm.DB) {
	adminRepository := repositories.NewAdminRepository(db)
	adminService := services.NewAdminService(adminRepository)
	adminController := controllers.NewAdminAuthController(adminService)

	auth := router.Group("/admin")
	auth.POST("/signin", adminController.AdminLogin)
}
