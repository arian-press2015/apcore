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
	cservice services.CustomerService
	aservice services.AlbumService
}

func NewCustomerController(cservice services.CustomerService, aservice services.AlbumService) *CustomerController {
	return &CustomerController{cservice, aservice}
}

type CreateCustomerBody struct {
	Name    string `json:"name"`
	Details string `json:"details"`
	Phone   string `json:"phone"`
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

	if err := ctrl.cservice.CreateCustomer(newCustomer); err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, newCustomer, messages.MsgSuccessful, nil, http.StatusCreated)
}

func (ctrl *CustomerController) GetCustomers(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	customers, err := ctrl.cservice.GetCustomers(offset, limit)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	count, err := ctrl.cservice.GetCustomerCount()
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	pagination := &response.Pagination{
		Offset: offset,
		Limit:  limit,
		Count:  count,
	}

	response.Success(c, customers, messages.MsgSuccessful, pagination, http.StatusOK)
}

func (ctrl *CustomerController) GetCustomerBySlug(c *gin.Context) {
	slug := c.Param("slug")

	customer, err := ctrl.cservice.GetCustomerBySlug(slug)
	if err != nil {
		response.Error(c, nil, messages.MsgNotFound, http.StatusNotFound)
		return
	}

	response.Success(c, customer, messages.MsgSuccessful, nil, http.StatusOK)
}

type UpdateCustomerBody struct {
	Name    string `json:"name"`
	Details string `json:"details"`
	Phone   string `json:"phone"`
}

func (ctrl *CustomerController) UpdateCustomer(c *gin.Context) {
	var input UpdateCustomerBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	slug := c.Param("slug")

	existingCustomer, err := ctrl.cservice.GetCustomerBySlug(slug)
	if err != nil {
		response.Error(c, nil, "Customer not found", http.StatusNotFound)
		return
	}

	existingCustomer.Name = input.Name
	existingCustomer.Details = input.Details
	existingCustomer.Phone = input.Phone

	if err := ctrl.cservice.UpdateCustomer(existingCustomer); err != nil {
		response.Error(c, nil, "Failed to update customer", http.StatusInternalServerError)
		return
	}

	response.Success(c, existingCustomer, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *CustomerController) GetAlbum(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	slug := c.Param("slug")
	owner, err := ctrl.cservice.GetCustomerBySlug(slug)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	album, err := ctrl.aservice.GetAlbum(offset, limit, owner.ID)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	count, err := ctrl.aservice.GetAlbumCount(owner.ID)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	pagination := &response.Pagination{
		Offset: offset,
		Limit:  limit,
		Count:  count,
	}

	response.Success(c, album, messages.MsgSuccessful, pagination, http.StatusOK)
}

type CreateAlbumBody struct {
	Name string `json:"name"`
}

func (ctrl *CustomerController) AddToAlbum(c *gin.Context) {
	var input CreateCustomerBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	slug := c.Param("slug")
	owner, err := ctrl.cservice.GetCustomerBySlug(slug)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	newAlbum := &models.CustomerAlbum{
		Name:    input.Name,
		OwnerId: owner.ID,
		Address: "/images/folan.jpg",
	}

	if err := ctrl.aservice.AddToAlbum(newAlbum); err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, newAlbum, messages.MsgSuccessful, nil, http.StatusCreated)
}

func (ctrl *CustomerController) DeleteFromAlbum(c *gin.Context) {
	slug := c.Param("slug")
	imageName := c.Param("imageName")

	owner, err := ctrl.cservice.GetCustomerBySlug(slug)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	err = ctrl.aservice.DeleteFromAlbum(imageName, owner.ID)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, nil, messages.MsgSuccessful, nil, http.StatusOK)
}
