package middlewares

import "github.com/gin-gonic/gin"

func SetupMiddlewares(router *gin.Engine) {
	router.Use(TrackIdMiddleware())
	router.Use(LocaleMiddleware())
	router.Use(RecoveryMiddleware())
	router.Use(ResponseHandlerMiddleware())
	// router.Use(authz.NewAuthorizer(acl.Enforcer))
}
