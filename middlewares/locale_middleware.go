package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LocaleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("Locale")
		if locale == "" {
			locale = "en"
		}
		c.Set("Locale", locale)
		c.Next()
	}
}