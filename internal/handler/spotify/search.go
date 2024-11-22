package spotify

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/response"
)

func (h *Handler) Search(c *gin.Context) {
	ctx := c.Request.Context()

	query := c.Query("query")
	pageSizeStr := c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 1
	}

	res, err := h.service.Search(ctx, query, pageSize, pageIndex)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
