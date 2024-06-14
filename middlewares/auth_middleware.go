package middlewares

import (
	"apcore/response"
	"apcore/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStringRaw := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(tokenStringRaw, "Bearer ")

		if tokenString == "" {
			response.Error(c, nil, "Authorization token not provided", http.StatusUnauthorized)
			c.Abort()
			return
		}

		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			response.Error(c, nil, "Internal Server Error", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
