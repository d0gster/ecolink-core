package handlers

import (
	"ecolink-core/internal/middleware"

	"github.com/gin-gonic/gin"
)

type CSRFHandler struct{}

func NewCSRFHandler() *CSRFHandler {
	return &CSRFHandler{}
}

// GetCSRFToken generates and returns a CSRF token for the authenticated user
func (h *CSRFHandler) GetCSRFToken(c *gin.Context) {
	middleware.GenerateCSRFToken(c)
}
