package middleware

import (
	"ecolink-core/internal/auth/usecase"
	"ecolink-core/internal/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireAuth middleware validates JWT tokens and injects user context
func RequireAuth(tokenService usecase.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from cookie or Authorization header
		token := extractToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, errors.NewAuthError("Authentication required"))
			c.Abort()
			return
		}

		// Validate token
		claims, err := tokenService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errors.NewAuthError("Invalid or expired token"))
			c.Abort()
			return
		}

		// Inject user context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Next()
	}
}

// extractToken gets JWT from cookie or Authorization header
func extractToken(c *gin.Context) string {
	// Try cookie first (primary method)
	if token, err := c.Cookie("ecolink_token"); err == nil && token != "" {
		return token
	}

	// Fallback to Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	return ""
}

// SecurityHeaders applies OWASP recommended security headers
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevent MIME type sniffing
		c.Header("X-Content-Type-Options", "nosniff")
		
		// Prevent clickjacking
		c.Header("X-Frame-Options", "DENY")
		
		// XSS protection (legacy but still useful)
		c.Header("X-XSS-Protection", "1; mode=block")
		
		// Control referrer information
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		
		// Content Security Policy
		c.Header("Content-Security-Policy", 
			"default-src 'self'; "+
			"script-src 'self' 'unsafe-inline'; "+
			"style-src 'self' 'unsafe-inline'; "+
			"img-src 'self' data: https:; "+
			"font-src 'self' https:; "+
			"connect-src 'self' https://accounts.google.com https://oauth2.googleapis.com")
		
		// HSTS for HTTPS (only in production)
		if c.Request.TLS != nil {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		}
		
		// Permissions policy
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		
		// Cache control for sensitive endpoints
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
			c.Header("Pragma", "no-cache")
			c.Header("Expires", "0")
		}
		
		// Remove server information
		c.Header("Server", "")
		
		c.Next()
	}
}

// CORS configures Cross-Origin Resource Sharing with security considerations
func CORS(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// Validate origin against allowed list
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				break
			}
		}
		
		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
		} else {
			// Don't set CORS headers for disallowed origins
			c.Header("Access-Control-Allow-Origin", "")
			c.Header("Access-Control-Allow-Credentials", "false")
		}
		
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-CSRF-Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Max-Age", "86400") // 24 hours

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
