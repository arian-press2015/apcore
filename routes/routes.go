package routes

import (
	_ "apcore/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine) {
	PingRoutes(router)
	AuthRoutes(router)
	UsersRoutes(router)
	RolesRoutes(router)
	AdminRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
