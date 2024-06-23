package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRouter),
	fx.Provide(NewRoutes),
)

func NewRouter() *gin.Engine {
	return gin.Default()
}

type Routes struct {
	controllers       *controllers.Controllers
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware
}

func NewRoutes(controllers *controllers.Controllers, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) *Routes {
	return &Routes{
		controllers:       controllers,
		jwtAuthMiddleware: jwtAuthMiddleware,
	}
}

func (r *Routes) SetupRoutes(router *gin.Engine) {
	PingRoutes(router, r.controllers.PingController)
	AuthRoutes(router, r.controllers.AuthController, r.jwtAuthMiddleware)
	UsersRoutes(router, r.controllers.UserController, r.jwtAuthMiddleware)
	RolesRoutes(router, r.controllers.RoleController, r.jwtAuthMiddleware)
	AdminAuthRoutes(router, r.controllers.AdminAuthController, r.jwtAuthMiddleware)
	SwaggerRoutes(router)

	router.NoMethod(func(c *gin.Context) {
		response.Error(c, nil, messages.MsgMethodNotAllowed, http.StatusMethodNotAllowed)
	})

	router.NoRoute(func(c *gin.Context) {
		response.Error(c, nil, messages.MsgNotFound, http.StatusNotFound)
	})
}
