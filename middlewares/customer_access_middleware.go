package middlewares

import (
	"apcore/messages"
	"apcore/repositories"
	"apcore/response"
	"apcore/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerAccessMiddleware struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerAccessMiddleware(customerRepository repositories.CustomerRepository) *CustomerAccessMiddleware {
	return &CustomerAccessMiddleware{customerRepository}
}

func (cam *CustomerAccessMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDString := c.GetString("userID")
		customerIDString := c.Param("customerID")

		userID, err := utils.UUIDParser(userIDString)
		if err != nil {
			response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
			c.Abort()
			return
		}

		customerID, err := utils.UUIDParser(customerIDString)
		if err != nil {
			response.Error(c, nil, messages.MsgUnauthorized, http.StatusBadRequest)
			c.Abort()
			return
		}

		hasAccess, _ := cam.customerRepository.CheckUserHasAccessToCustomer(userID, customerID)
		if !hasAccess {
			response.Error(c, nil, messages.MsgUnauthorized, http.StatusBadRequest)
			c.Abort()
			return
		}
		c.Next()
	}
}
