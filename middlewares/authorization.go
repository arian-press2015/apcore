package middlewares

import (
	"apcore/messages"
	"apcore/repositories"
	"apcore/response"
	"apcore/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorizationMiddleware struct {
	permissionRepository repositories.PermissionRepository
}

func NewAuthorizationMiddleware(permissionRepository repositories.PermissionRepository) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{permissionRepository}
}

func (am *AuthorizationMiddleware) Middleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDString := c.GetString("userID")

		userID, err := utils.UUIDParser(userIDString)
		if err != nil {
			response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
			c.Abort()
			return
		}

		hasPermission, _ := am.permissionRepository.HasUserSomePermission(userID,requiredPermission) 

		if !hasPermission {
			response.Error(c, nil, messages.MsgUnauthorized, http.StatusBadRequest)
			c.Abort()
			return
		}
		c.Next()
	}
}
