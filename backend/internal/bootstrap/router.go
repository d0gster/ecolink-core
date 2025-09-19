package bootstrap

import (
	"ecolink-core/internal/auth/delivery/http"
	"ecolink-core/internal/auth/repository"
	"ecolink-core/internal/auth/usecase"
	"ecolink-core/internal/config"
	"ecolink-core/internal/handlers"
	"ecolink-core/internal/middleware"
	"ecolink-core/internal/security"
	"ecolink-core/internal/services"
	"ecolink-core/internal/validation"
	"ecolink-core/pkg/database"

	"github.com/gin-gonic/gin"
)

// Application holds all application dependencies
type Application struct {
	Config *config.Config
	DB     database.Database
	Router *gin.Engine
}

// NewApplication creates and wires all application dependencies
func NewApplication(cfg *config.Config, db database.Database) *Application {
	router := setupRouter(cfg, db)
	return &Application{
		Config: cfg,
		DB:     db,
		Router: router,
	}
}

// setupRouter configures the HTTP router with all middleware and routes
func setupRouter(cfg *config.Config, db database.Database) *gin.Engine {
	// Initialize services
	linkService := services.NewLinkService(db, cfg.BaseURL)
	userService := services.NewUserService(db)

	// Initialize JWT service
	jwtService := security.NewJWTService(
		cfg.Security.JWTSecret,
		"ecolink",
		"ecolink-users",
	)

	// Initialize auth services
	tokenService := usecase.NewJWTTokenService(
		cfg.Security.JWTSecret,
		"ecolink",
	)

	googleConfig := usecase.GoogleConfig{
		ClientID:     cfg.GoogleAuth.ClientID,
		ClientSecret: cfg.GoogleAuth.ClientSecret,
		RedirectURI:  cfg.FrontendURL + "/auth/callback/google",
	}

	// Initialize user repository
	userRepo := repository.NewInMemoryUserRepository()

	// Initialize auth service
	authService := usecase.NewAuthService(userRepo, tokenService, googleConfig)

	// Initialize handlers
	linkHandler := handlers.NewLinkHandler(linkService)
	userHandler := handlers.NewUserHandler(userService)
	csrfHandler := handlers.NewCSRFHandler()

	// Initialize validator
	validator := validation.NewValidator()

	// Legacy auth handler (to be replaced)
	authHandler := handlers.NewAuthHandler(cfg, userService, jwtService)

	// New auth handler (when repository is implemented)
	newAuthHandler := http.NewAuthHandler(*authService, validator)

	// Setup router
	r := gin.Default()

	// Global middleware
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.CORS([]string{cfg.FrontendURL}))

	// Public routes
	r.GET("/:code", linkHandler.RedirectLink)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "ecolink-core"})
	})

	// Auth routes (public)
	auth := r.Group("/auth")
	{
		auth.POST("/google/callback", newAuthHandler.GoogleCallback)
		auth.POST("/logout", authHandler.Logout)
		auth.POST("/register", newAuthHandler.Register)
		auth.POST("/login", newAuthHandler.Login)
		auth.GET("/google", newAuthHandler.GoogleLogin)
	}

	// Protected API routes
	api := r.Group("/api/v1")
	api.Use(middleware.RequireAuth(tokenService))
	{
		// CSRF token endpoint
		api.GET("/csrf-token", csrfHandler.GetCSRFToken)

		// User endpoints
		api.GET("/me", newAuthHandler.GetCurrentUser)
		api.GET("/profile", userHandler.GetProfile)

		// Link endpoints (with CSRF protection)
		protected := api.Group("")
		protected.Use(middleware.CSRFMiddleware(cfg.Security.CSRFSecret))
		{
			protected.POST("/links", linkHandler.CreateLink)
			protected.DELETE("/links/:code", linkHandler.DeleteLink)
		}

		// Read-only endpoints (no CSRF needed)
		api.GET("/links", linkHandler.GetUserLinks)
	}

	return r
}

// NewRouter maintains backward compatibility (deprecated)
func NewRouter(cfg *config.Config, db database.Database) *gin.Engine {
	return setupRouter(cfg, db)
}
