package main

import (
	"context"
	"ecolink-core/internal/bootstrap"
	"ecolink-core/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	// 2. Establish database connection
	db, err := bootstrap.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// 3. Initialize application with dependency injection
	app := bootstrap.NewApplication(cfg, db)
	router := app.Router

	// 4. Configure HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 5. Start server in goroutine
	go func() {
		log.Printf("🌱 EcoLink Core starting on port %s", cfg.Port)
		log.Printf("🔒 Security: JWT=%t, CSRF=%t", len(cfg.Security.JWTSecret) >= 32, len(cfg.Security.CSRFSecret) >= 32)
		log.Printf("🗄️  Database: %s", cfg.Database.Type)
		log.Printf("🌐 Frontend: %s", cfg.FrontendURL)
		
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed to start: %v", err)
		}
	}()

	// 6. Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down server...")

	// 7. Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("❌ Server forced to shutdown: %v", err)
	} else {
		log.Println("✅ Server exited gracefully")
	}
}
