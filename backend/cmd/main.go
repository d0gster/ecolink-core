package main

import (
	"ecolink-core/internal/config"
	"ecolink-core/internal/handlers"
	"ecolink-core/internal/services"
	"ecolink-core/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	
	// Inicializa banco em memória para demo
	db := database.NewMemoryDB()
	
	// Inicializa services
	linkService := services.NewLinkService(db, cfg.BaseURL)
	
	// Inicializa handlers
	linkHandler := handlers.NewLinkHandler(linkService)
	
	// Setup router
	r := gin.Default()
	
	// CORS middleware para desenvolvimento
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, X-User-ID")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})
	
	// Routes
	api := r.Group("/api/v1")
	{
		api.POST("/links", linkHandler.CreateLink)
		api.GET("/links", linkHandler.GetUserLinks)
	}
	
	// Redirect route (sem /api prefix)
	r.GET("/:code", linkHandler.RedirectLink)
	
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "ecolink-core"})
	})
	
	log.Printf("🌱 EcoLink Core rodando na porta %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}