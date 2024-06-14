package middlewares

import (
	"apcore/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)

				trackId, _ := c.Get("trackId")
				errorResponse := response.NewResponse(nil, "Internal Server Error", nil, http.StatusInternalServerError)
				errorResponse.TrackId = trackId.(string)
				c.JSON(http.StatusInternalServerError, errorResponse)
				c.Abort()
			}
		}()
		c.Next()
	}
}
