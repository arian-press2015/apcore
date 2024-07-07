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
	customers.GET(":name", ctrl.GetCustomerByName)
	customers.PUT(":name", ctrl.UpdateCustomer)
}
