package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/response"
)

func (h *Handler) SignIn(c *gin.Context) {
	var request memberships.SignInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := h.service.SignIn(request)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response := response.Response{
		Status:  http.StatusOK,
		Message: "Login successfully",
		Data: memberships.SignInResponse{
			AccessToken: accessToken,
		},
	}

	c.JSON(http.StatusOK, response)
}
