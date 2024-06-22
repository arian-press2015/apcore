package middlewares

import (
	"apcore/logger"
	"apcore/messages"
	"apcore/response"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				stackTrace := debug.Stack()
				logger.Error("Recovered from panic",
					zap.Any("error", err),
					zap.String("request", c.Request.URL.Path),
					zap.ByteString("stack", stackTrace),
				)
				trackID, _ := c.Get("trackId")

				if resp, exists := c.Get("response"); exists {
					responseInstance := resp.(*response.Response)
					responseInstance.TrackId = trackID.(string)
					c.JSON(responseInstance.StatusCode, responseInstance)
				} else {
					errorResponse := response.NewResponse(nil, messages.MsgInternalServerError, nil, http.StatusInternalServerError)
					errorResponse.TrackId = trackID.(string)
					c.JSON(http.StatusInternalServerError, errorResponse)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
