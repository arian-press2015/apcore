package middlewares

import (
	"apcore/messages"
	"apcore/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ResponseHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// responnse handler prevents serving static files, so I bypass it for swagger route
		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			return
		}

		locale := c.GetString("Locale")
		resp, exists := c.Get("response")
		if exists {
			responseInstance := resp.(*response.Response)

			trackId, _ := c.Get("trackId")
			responseInstance.TrackId = trackId.(string)
			responseInstance.Message = messages.TranslateMessage(responseInstance.Message, locale)

			c.JSON(responseInstance.StatusCode, responseInstance)
		} else {
			trackId, _ := c.Get("trackId")
			errorResponse := response.NewResponse(nil, messages.MsgInternalServerError, nil, http.StatusInternalServerError)
			errorResponse.TrackId = trackId.(string)
			c.JSON(http.StatusInternalServerError, errorResponse)
		}
	}
}
