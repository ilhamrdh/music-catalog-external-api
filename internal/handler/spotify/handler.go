package spotify

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/middleware"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/spotify"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=spotify
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchResponse, error)
}

type Handler struct {
	engine  *gin.Engine
	service service
}

func NewHandler(engine *gin.Engine, service service) *Handler {
	return &Handler{engine: engine, service: service}
}

func (h *Handler) SpotifyRoute() {
	route := h.engine.Group("/tracks")
	route.Use(middleware.AuthMiddleware())
	route.GET("/search", h.Search)
}
