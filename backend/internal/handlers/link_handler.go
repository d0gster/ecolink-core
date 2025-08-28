package handlers

import (
	"ecolink-core/internal/models"
	"ecolink-core/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	linkService *services.LinkService
}

func NewLinkHandler(linkService *services.LinkService) *LinkHandler {
	return &LinkHandler{
		linkService: linkService,
	}
}

func (h *LinkHandler) CreateLink(c *gin.Context) {
	var req models.CreateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Para demo, usa userID fixo. Na versão completa, vem do JWT
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		userID = "demo-user"
	}
	
	response, err := h.linkService.CreateLink(req.URL, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, response)
}

func (h *LinkHandler) RedirectLink(c *gin.Context) {
	shortCode := c.Param("code")
	
	originalURL, err := h.linkService.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}
	
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func (h *LinkHandler) GetUserLinks(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		userID = "demo-user"
	}
	
	links, err := h.linkService.GetUserLinks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"links": links})
}