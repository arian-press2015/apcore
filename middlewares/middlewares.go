package middlewares

import (
	"apcore/config"
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
	cfg    *config.Config
}

func NewMiddlewares(logger *logger.Logger, cfg *config.Config) *Middlewares {
	return &Middlewares{logger: logger, cfg: cfg}
}

func (m *Middlewares) SetupMiddlewares(router *gin.Engine) {
	router.Use(CorsMiddleware(m.cfg))
	router.Use(TrackIdMiddleware())
	router.Use(LocaleMiddleware())
	router.Use(RecoveryMiddleware(m.logger))
	router.Use(ResponseHandlerMiddleware())
	// router.Use(authz.NewAuthorizer(acl.Enforcer))
}
