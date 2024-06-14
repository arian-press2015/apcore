package middlewares

import (
	"apcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		resp, exists := c.Get("response")
		if exists {
			responseInstance := resp.(*response.Response)

			trackId, _ := c.Get("trackId")
			responseInstance.TrackId = trackId.(string)

			c.JSON(responseInstance.StatusCode, responseInstance)
		} else {
			trackId, _ := c.Get("trackId")
			errorResponse := response.NewResponse(nil, "Internal Server Error", nil, http.StatusInternalServerError)
			errorResponse.TrackId = trackId.(string)
			c.JSON(http.StatusInternalServerError, errorResponse)
		}
	}
}
