package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TrackIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		trackId := uuid.New().String()

		c.Set("trackId", trackId)

		c.Next()
	}
}
