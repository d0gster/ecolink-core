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

	// Get user ID from middleware context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	response, err := h.linkService.CreateLink(req.URL, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *LinkHandler) RedirectLink(c *gin.Context) {
	shortCode := c.Param("code")

	originalURL, err := h.linkService.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found", "code": shortCode})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func (h *LinkHandler) GetUserLinks(c *gin.Context) {
	// Get user ID from middleware context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	links, err := h.linkService.GetUserLinks(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"links": links})
}

func (h *LinkHandler) DeleteLink(c *gin.Context) {
	code := c.Param("code")

	// Get user ID from middleware context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	err := h.linkService.DeleteLink(code, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link deleted successfully"})
}
