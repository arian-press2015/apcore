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

type RoleController struct {
	service services.RoleService
}

func NewRoleController(service services.RoleService) *RoleController {
	return &RoleController{service}
}

func (ctrl *RoleController) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.CreateRole(&role); err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, role, "success", nil, http.StatusOK)
}

func (ctrl *RoleController) GetRoles(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	roles, err := ctrl.service.GetRoles(offset, limit)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, roles, messages.MsgSuccessful, nil, http.StatusOK)
}
