package middlewares

import (
	"apcore/messages"
	"apcore/response"
	"apcore/services"
	"apcore/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleware struct {
	jwtService  *jwt.JWTService
	userService services.UserService
}

func NewJWTAuthMiddleware(jwtService *jwt.JWTService, userService services.UserService) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{jwtService, userService}
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

		user, err := jam.userService.GetUserByPhone(claims.Phone)
		if err != nil {
			response.Error(c, nil, "User not found", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("userID", user.BaseModel.ID)
		c.Next()
	}
}
