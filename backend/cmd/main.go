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
	// Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
	}
	
	cfg := config.Load()
	
	// Inicializa banco de dados
	var db database.Database
	if cfg.Database.Type == "firestore" && cfg.Database.ProjectID != "" {
		firestoreDB, err := database.NewFirestoreDB(cfg.Database.ProjectID, cfg.Database.Credentials)
		if err != nil {
			log.Printf("Erro ao conectar Firestore: %v. Usando memória.", err)
			db = database.NewMemoryDB()
		} else {
			db = firestoreDB
			log.Println("✅ Firestore conectado")
		}
	} else {
		db = database.NewMemoryDB()
		log.Println("⚠️  Usando banco em memória")
	}
	
	// Inicializa services
	linkService := services.NewLinkService(db, cfg.BaseURL)
	userService := services.NewUserService(db)
	
	// Inicializa handlers
	linkHandler := handlers.NewLinkHandler(linkService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(
		os.Getenv("GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		userService,
	)
	
	// Setup router
	r := gin.Default()
	
	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-User-ID")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})
	
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
		log.Println("✅ Auth0 middleware ativo")
	} else {
		log.Println("⚠️  Auth0 não configurado - rotas desprotegidas")
	}
	
	api.POST("/links", linkHandler.CreateLink)
	api.GET("/links", linkHandler.GetUserLinks)
	api.DELETE("/links/:code", linkHandler.DeleteLink)
	api.GET("/profile", userHandler.GetProfile)
	
	log.Printf("🌱 EcoLink Core rodando na porta %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}