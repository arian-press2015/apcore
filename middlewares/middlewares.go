package middlewares

import (
	"apcore/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMiddlewares),
	fx.Provide(NewJWTAuthMiddleware),
	fx.Provide(NewCustomerAccessMiddleware),
	fx.Provide(NewAuthorizationMiddleware),
)

type Middlewares struct {
	logger *logger.Logger
}

func NewMiddlewares(logger *logger.Logger) *Middlewares {
	return &Middlewares{logger: logger}
}

func (m *Middlewares) SetupMiddlewares(router *gin.Engine) {
	router.Use(TrackIdMiddleware())
	router.Use(LocaleMiddleware())
	router.Use(RecoveryMiddleware(m.logger))
	router.Use(ResponseHandlerMiddleware())
	// router.Use(authz.NewAuthorizer(acl.Enforcer))
}
