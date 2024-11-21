package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
type service interface {
	SignUp(request memberships.SignUpRequest) error
}

type Handler struct {
	e       *gin.Engine
	service service
}

func NewHandler(e *gin.Engine, service service) *Handler {
	return &Handler{e: e, service: service}
}

func (h *Handler) RegisterRoute() {
	route := h.e.Group("/memberships")
	route.POST("/sign-up", h.SignUp)
}
