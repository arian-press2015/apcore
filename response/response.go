package response

import "github.com/gin-gonic/gin"

type Response struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
	TrackId    string      `json:"trackId"`
	StatusCode int         `json:"-"`
}

func NewResponse(data interface{}, message string, pagination interface{}, statusCode int) *Response {
	return &Response{
		Data:       data,
		Message:    message,
		Pagination: pagination,
		StatusCode: statusCode,
	}
}

func Success(c *gin.Context, data interface{}, message string, pagination interface{}, statusCode int) {
	resp := NewResponse(data, message, pagination, statusCode)
	c.Set("response", resp)
}

func Error(c *gin.Context, data interface{}, message string, statusCode int) {
	resp := NewResponse(data, message, nil, statusCode)
	c.Set("response", resp)
}
