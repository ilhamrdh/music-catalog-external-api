package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/response"
)

func (h *Handler) SignUp(c *gin.Context) {
	var request memberships.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.SignUp(request)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response := response.Response{
		Status:  http.StatusCreated,
		Message: "Register successfully",
	}
	c.JSON(http.StatusCreated, response)
}
