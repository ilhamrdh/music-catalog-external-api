package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
)

func (h *Handler) SignIn(c *gin.Context) {
	var request memberships.SignInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	accessToken, err := h.service.SignIn(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberships.SignInResponse{
		AccessToken: accessToken,
	})

}
