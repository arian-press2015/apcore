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
				trackID, _ := c.Get("trackId")

				if resp, exists := c.Get("response"); exists {
					responseInstance := resp.(*response.Response)
					responseInstance.TrackId = trackID.(string)
					c.JSON(responseInstance.StatusCode, responseInstance)
				} else {
					errorResponse := response.NewResponse(nil, "Internal Server Error", nil, http.StatusInternalServerError)
					errorResponse.TrackId = trackID.(string)
					c.JSON(http.StatusInternalServerError, errorResponse)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}