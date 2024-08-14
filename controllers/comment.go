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

type CommentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) *CommentController {
	return &CommentController{service}
}

func (ctrl *CommentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.CreateComment(&comment); err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, comment, "success", nil, http.StatusOK)
}

func (ctrl *CommentController) GetComments(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	comments, err := ctrl.service.GetComments(offset, limit)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	count, err := ctrl.service.GetCommentCount()
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	pagination := &response.Pagination{
		Offset: offset,
		Limit:  limit,
		Count:  count,
	}

	response.Success(c, comments, messages.MsgSuccessful, pagination, http.StatusOK)
}
