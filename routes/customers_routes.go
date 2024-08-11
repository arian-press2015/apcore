package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func CustomersRoutes(router *gin.Engine, ctrl *controllers.CustomerController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	customers := router.Group("/customers")
	customers.POST("", ctrl.CreateCustomer)
	customers.GET("", ctrl.GetCustomers)
	customers.GET("album", ctrl.GetAlbum)
	customers.POST("album", ctrl.AddToAlbum)
	customers.DELETE("album/:imageName", ctrl.DeleteFromAlbum)
	customers.GET(":slug", ctrl.GetCustomerBySlug)
	customers.PUT(":slug", ctrl.UpdateCustomer)
}
