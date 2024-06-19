package middlewares

import (
	"apcore/messages"
	"apcore/models"
	"apcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FeatureFlagMiddleware(db *gorm.DB, featureName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var feature models.Feature
		result := db.Where("name = ?", featureName).First(&feature)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				response.Error(c, nil, messages.MsgFeatureNotFound, http.StatusNotFound)
				return
			}
			response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
			return
		}

		if !feature.Enabled {
			response.Error(c, nil, messages.MsgFeatureDisabled, http.StatusForbidden)
			return
		}

		c.Next()
	}
}
