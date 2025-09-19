package integration

import (
	"bytes"
	"ecolink-core/internal/middleware"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCSRFProtection(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GET request bypasses CSRF", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.CSRFMiddleware("test-secret"))
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		req := httptest.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("POST without CSRF token fails", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.CSRFMiddleware("test-secret"))
		router.POST("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		body := bytes.NewBufferString(`{"test": "data"}`)
		req := httptest.NewRequest("POST", "/test", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, 403, w.Code)
	})

	t.Run("POST with valid CSRF token succeeds", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.CSRFMiddleware("test-secret"))
		router.GET("/csrf", middleware.GenerateCSRFToken)
		router.POST("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// Get CSRF token
		req := httptest.NewRequest("GET", "/csrf", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		require.Equal(t, 200, w.Code)

		var tokenResp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &tokenResp)
		require.NoError(t, err)

		token := tokenResp["csrf_token"]
		require.NotEmpty(t, token)

		// Extract cookie
		cookies := w.Result().Cookies()
		var csrfCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "csrf_token" {
				csrfCookie = cookie
				break
			}
		}
		require.NotNil(t, csrfCookie)

		// Make POST request with CSRF token
		body := bytes.NewBufferString(`{"test": "data"}`)
		req = httptest.NewRequest("POST", "/test", body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-CSRF-Token", token)
		req.AddCookie(csrfCookie)

		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}