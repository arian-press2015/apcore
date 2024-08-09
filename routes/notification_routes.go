package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(router *gin.Engine, ctrl *controllers.NotificationController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	notifications := router.Group("/notifications")
	notifications.GET("", ctrl.GetNotifications)
	notifications.PUT("markAsRead", ctrl.MarkAllAsRead)
	notifications.PUT(":uuid/markAsRead", ctrl.MarkAsRead)
}
