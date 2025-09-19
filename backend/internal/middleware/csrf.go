package middleware

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CSRFTokenLength = 32
	CSRFCookieName  = "csrf_token"
	CSRFHeaderName  = "X-CSRF-Token"
)

// CSRFMiddleware implements Double Submit Cookie pattern for CSRF protection
func CSRFMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip CSRF for safe methods
		if c.Request.Method == "GET" || c.Request.Method == "HEAD" || c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		// Get CSRF token from cookie
		cookieToken, err := c.Cookie(CSRFCookieName)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "CSRF token missing"})
			c.Abort()
			return
		}

		// Get CSRF token from header
		headerToken := c.GetHeader(CSRFHeaderName)
		if headerToken == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "CSRF token required in header"})
			c.Abort()
			return
		}

		// Validate tokens match using constant-time comparison
		if !validateCSRFToken(cookieToken, headerToken) {
			c.JSON(http.StatusForbidden, gin.H{"error": "CSRF token mismatch"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// GenerateCSRFToken creates a new CSRF token and sets it in a cookie
func GenerateCSRFToken(c *gin.Context) {
	token := generateSecureToken()
	
	// Set CSRF token in cookie
	c.SetCookie(
		CSRFCookieName,
		token,
		int(time.Hour.Seconds()), // 1 hour expiration
		"/",
		"",
		false, // Not HTTPS-only for development
		false, // Not HTTP-only so frontend can read it
	)
	
	c.JSON(http.StatusOK, gin.H{
		"csrf_token": token,
	})
}

// validateCSRFToken performs constant-time comparison of CSRF tokens
func validateCSRFToken(cookieToken, headerToken string) bool {
	return subtle.ConstantTimeCompare([]byte(cookieToken), []byte(headerToken)) == 1
}

// generateSecureToken creates a cryptographically secure random token
func generateSecureToken() string {
	bytes := make([]byte, CSRFTokenLength)
	if _, err := rand.Read(bytes); err != nil {
		panic("Failed to generate secure random token")
	}
	return base64.URLEncoding.EncodeToString(bytes)
}