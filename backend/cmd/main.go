package main

import (
	"ecolink-core/internal/config"
	"ecolink-core/internal/handlers"
	"ecolink-core/internal/middleware"
	"ecolink-core/internal/services"
	"ecolink-core/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system variables")
	}

	cfg := config.Load()

	// Initialize database
	var db database.Database
	if cfg.Database.Type == "firestore" && cfg.Database.ProjectID != "" {
		firestoreDB, err := database.NewFirestoreDB(cfg.Database.ProjectID, cfg.Database.Credentials)
		if err != nil {
			log.Printf("Error connecting to Firestore: %v. Using memory.", err)
			db = database.NewMemoryDB()
		} else {
			db = firestoreDB
			log.Println("✅ Firestore connected")
		}
	} else {
		db = database.NewMemoryDB()
		log.Println("⚠️  Using in-memory database")
	}

	// Initialize services
	linkService := services.NewLinkService(db, cfg.BaseURL)
	userService := services.NewUserService(db)

	// Initialize handlers
	linkHandler := handlers.NewLinkHandler(linkService)
	userHandler := handlers.NewUserHandler(userService)

	googleClientID, ok := os.LookupEnv("GOOGLE_CLIENT_ID")
	if !ok || googleClientID == "" {
		log.Fatal("GOOGLE_CLIENT_ID not set")
	}

	googleClientSecret, ok := os.LookupEnv("GOOGLE_CLIENT_SECRET")
	if !ok || googleClientSecret == "" {
		log.Fatal("GOOGLE_CLIENT_SECRET not set")
	}

	authHandler := handlers.NewAuthHandler(
		googleClientID,
		googleClientSecret,
		userService,
	)

	// Setup router
	r := gin.Default()

	// CORS middleware
	r.Use(CORSMiddleware(cfg.FrontendURL))

	// Public routes
	r.GET("/:code", linkHandler.RedirectLink)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "ecolink-core"})
	})

	// Auth routes
	r.POST("/auth/google/callback", authHandler.GoogleCallback)

	// Protected API routes
	api := r.Group("/api/v1")
	if cfg.Auth0Domain != "" && cfg.Auth0Audience != "" {
		api.Use(middleware.AuthMiddleware(cfg.Auth0Domain, cfg.Auth0Audience))
		log.Println("✅ Auth0 middleware active")
	} else {
		log.Println("⚠️  Auth0 not configured - unprotected routes")
	}

	api.POST("/links", linkHandler.CreateLink)
	api.GET("/links", linkHandler.GetUserLinks)
	api.DELETE("/links/:code", linkHandler.DeleteLink)
	api.GET("/profile", userHandler.GetProfile)

	log.Printf("🌱 EcoLink Core running on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}

// CORSMiddleware handles Cross-Origin Resource Sharing
func CORSMiddleware(frontendURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", frontendURL)
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-User-ID")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
