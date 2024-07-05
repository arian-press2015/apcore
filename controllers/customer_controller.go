package controllers

import (
	"apcore/messages"
	"apcore/models"
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

type CreateCustomerBody struct {
	Name      string  `json:"name"`
	Details   string  `json:"details"`
	Phone     string  `json:"phone"`
}

func (ctrl *CustomerController) CreateCustomer(c *gin.Context) {
	var input CreateCustomerBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	newCustomer := &models.Customer{
		Name:       input.Name,
		Details:    input.Details,
		Phone:      input.Phone,
		IsActive:   false,
		IsDisabled: false,
	}

	if err := ctrl.service.CreateCustomer(newCustomer); err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, newCustomer, messages.MsgSuccessful, nil, http.StatusCreated)
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

func (ctrl *CustomerController) GetCustomerByName(c *gin.Context) {
	name := c.Param("name")

	customer, err := ctrl.service.GetCustomerByName(name)
	if err != nil {
		response.Error(c, nil, messages.MsgNotFound, http.StatusNotFound)
		return
	}

	response.Success(c, customer, messages.MsgSuccessful, nil, http.StatusOK)
}
