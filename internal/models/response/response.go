package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"errors,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
}

func ErrorResponse(c *gin.Context, status int, err string) {
	c.JSON(status, Response{
		Status: status,
		Error:  err,
	})
}
