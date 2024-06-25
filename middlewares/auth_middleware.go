package middlewares

import (
	"apcore/messages"
	"apcore/response"
	"apcore/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleware struct {
	jwtService *jwt.JWTService
}

func NewJWTAuthMiddleware(jwtService *jwt.JWTService) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{jwtService}
}

func (jam *JWTAuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStringRaw := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(tokenStringRaw, "Bearer ")

		if tokenString == "" {
			response.Error(c, nil, messages.MsgNoAuthHeader, http.StatusUnauthorized)
			c.Abort()
			return
		}

		claims, err := jam.jwtService.VerifyJWT(tokenString)
		if err != nil {
			response.Error(c, nil, messages.MsgUnauthorized, http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("username", claims.Phone)
		c.Next()
	}
}
