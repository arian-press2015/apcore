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
	customers.GET(":slug", ctrl.GetCustomerBySlug)
	customers.PUT(":slug", ctrl.UpdateCustomer)

	// album
	customers.GET(":slug/album", ctrl.GetAlbum)
	customers.POST(":slug/album", ctrl.AddToAlbum)
	customers.DELETE(":slug/album/:imageName", ctrl.DeleteFromAlbum)

	// menu
	customers.GET(":slug/menu", ctrl.GetMenu)
	customers.POST(":slug/menu", ctrl.CreateMenu)
	customers.PUT(":slug/menu", ctrl.UpdateMenu)

	// categories and products
	customers.GET(":slug/menu/categories/:categorySlug", ctrl.GetCategoryProducts)
	customers.GET(":slug/products", ctrl.GetProducts)
	customers.GET(":slug/products/:productSlug", ctrl.GetProductBySlug)
}
