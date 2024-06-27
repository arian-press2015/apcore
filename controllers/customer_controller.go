package controllers

import (
	"apcore/messages"
	"apcore/response"
	"apcore/services"
	"apcore/utils/parsers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service services.CustomerService
}

func NewCustomerController(service services.CustomerService) *CustomerController {
	return &CustomerController{service}
}

func (ctrl *CustomerController) GetCustomers(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	customers, err := ctrl.service.GetCustomers(offset, limit)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, customers, messages.MsgSuccessful, nil, http.StatusOK)
}
