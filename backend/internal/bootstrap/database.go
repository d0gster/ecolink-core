package bootstrap

import (
	"ecolink-core/internal/config"
	"ecolink-core/pkg/database"
	"fmt"
	"log"
)

// NewDBConnection initializes database connection based on configuration
func NewDBConnection(cfg *config.Config) (database.Database, error) {
	switch cfg.Database.Type {
	case "firestore":
		if cfg.Database.ProjectID == "" {
			log.Println("⚠️  Firestore ProjectID not configured, falling back to memory")
			return database.NewMemoryDB(), nil
		}
		
		db, err := database.NewFirestoreDB(cfg.Database.ProjectID, cfg.Database.Credentials)
		if err != nil {
			log.Printf("❌ Error connecting to Firestore: %v. Using memory database", err)
			return database.NewMemoryDB(), nil
		}
		
		log.Println("✅ Firestore database connected")
		return db, nil
		
	case "memory":
		log.Println("⚠️  Using in-memory database (data will not persist)")
		return database.NewMemoryDB(), nil
		
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}
}