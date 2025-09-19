package integration

import (
	"bytes"
	"ecolink-core/internal/bootstrap"
	"ecolink-core/internal/config"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthFlow(t *testing.T) {
	// Setup test environment
	os.Setenv("JWT_SECRET", "test-secret-key-32-characters-long")
	os.Setenv("GOOGLE_CLIENT_ID", "test-client-id")
	os.Setenv("GOOGLE_CLIENT_SECRET", "test-client-secret")
	
	cfg, err := config.Load()
	assert.NoError(t, err)
	
	db, err := bootstrap.NewDBConnection(cfg)
	assert.NoError(t, err)
	
	router := bootstrap.NewRouter(cfg, db)

	t.Run("Health Check", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/health", nil)
		router.ServeHTTP(w, req)
		
		assert.Equal(t, 200, w.Code)
		
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ok", response["status"])
	})

	t.Run("CSRF Token Endpoint Requires Auth", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/csrf-token", nil)
		router.ServeHTTP(w, req)
		
		assert.Equal(t, 401, w.Code)
	})

	t.Run("Protected Endpoint Requires Auth", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/me", nil)
		router.ServeHTTP(w, req)
		
		assert.Equal(t, 401, w.Code)
	})

	t.Run("CSRF Protection on POST", func(t *testing.T) {
		// Test that POST requests without CSRF token are rejected
		linkData := map[string]string{
			"url": "https://example.com",
		}
		jsonData, _ := json.Marshal(linkData)
		
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/links", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		
		// Should be 401 (no auth) or 403 (no CSRF)
		assert.True(t, w.Code == 401 || w.Code == 403)
	})
}