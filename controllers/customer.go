package controllers

import (
	"apcore/dto"
	"apcore/messages"
	"apcore/models"
	"apcore/response"
	"apcore/services"
	"apcore/utils/parsers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	cservice  services.CustomerService
	aservice  services.AlbumService
	mservice  services.MenuService
	c2service services.CategoryService
	pservice  services.ProductService
}

func NewCustomerController(cservice services.CustomerService, aservice services.AlbumService, mservice services.MenuService, c2service services.CategoryService, pservice services.ProductService) *CustomerController {
	return &CustomerController{cservice, aservice, mservice, c2service, pservice}
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

func (ctrl *CustomerController) GetMenu(c *gin.Context) {
	slug := c.Param("slug")

	customer, err := ctrl.cservice.GetCustomerBySlug(slug)
	if err != nil {
		response.Error(c, nil, "Customer not found", http.StatusNotFound)
		return
	}

	menu, err := ctrl.mservice.GetMenu(customer.ID)
	if err != nil {
		response.Error(c, nil, messages.MsgNotFound, http.StatusNotFound)
		return
	}

	response.Success(c, menu, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *CustomerController) CreateMenu(c *gin.Context) {
	var input dto.CreateMenuBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	slug := c.Param("slug")
	customer, err := ctrl.cservice.GetCustomerBySlug(slug)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	err = ctrl.mservice.CreateMenu(customer.ID, input)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, nil, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *CustomerController) UpdateMenu(c *gin.Context) {
	var input dto.UpdateMenuBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	slug := c.Param("slug")
	customer, err := ctrl.cservice.GetCustomerBySlug(slug)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	err = ctrl.mservice.UpdateMenu(customer.ID, input)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, nil, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *CustomerController) GetCategoryProducts(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	customerSlug := c.Param("slug")
	customer, err := ctrl.cservice.GetCustomerBySlug(customerSlug)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	categorySlug := c.Param("categorySlug")
	category, err := ctrl.c2service.GetCategoryBySlug(categorySlug, customer.ID)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	fmt.Println("later, i need to add this to GetProducts conditions struct", category.ID.String())
	products, err := ctrl.pservice.GetProducts(offset, limit)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	count, err := ctrl.pservice.GetProductCount()
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	pagination := &response.Pagination{
		Offset: offset,
		Limit:  limit,
		Count:  count,
	}

	response.Success(c, products, messages.MsgSuccessful, pagination, http.StatusOK)
}

func (ctrl *CustomerController) GetProducts(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	slug := c.Param("slug")
	customer, err := ctrl.cservice.GetCustomerBySlug(slug)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	fmt.Println("later, i need to add this, and price filter to GetProducts conditions struct", customer.ID.String())
	products, err := ctrl.pservice.GetProducts(offset, limit)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	count, err := ctrl.pservice.GetProductCount()
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	pagination := &response.Pagination{
		Offset: offset,
		Limit:  limit,
		Count:  count,
	}

	response.Success(c, products, messages.MsgSuccessful, pagination, http.StatusOK)
}

func (ctrl *CustomerController) GetProductBySlug(c *gin.Context) {
	customerSlug := c.Param("slug")
	customer, err := ctrl.cservice.GetCustomerBySlug(customerSlug)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	productSlug := c.Param("productSlug")
	product, err := ctrl.pservice.GetProductBySlug(productSlug, customer.ID)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, product, messages.MsgSuccessful, nil, http.StatusOK)
}
