package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	var request memberships.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	err := h.service.SignUp(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
