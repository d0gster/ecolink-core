package http

import (
	"ecolink-core/internal/auth/usecase"
	"ecolink-core/internal/errors"
	"ecolink-core/internal/validation"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService usecase.AuthService
	validator   *validation.Validator
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type GoogleCallbackRequest struct {
	Code        string `json:"code" validate:"required"`
	State       string `json:"state" validate:"required"`
	RedirectURI string `json:"redirect_uri" validate:"required"`
}

func NewAuthHandler(authService usecase.AuthService, validator *validation.Validator) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator,
	}
}

// Register handles user registration with local credentials
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("Invalid request format"))
		return
	}

	// Validate input
	if err := h.validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
		return
	}

	// Create user
	user, err := h.authService.Register(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, errors.NewBusinessError("Registration failed", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login handles user authentication with local credentials
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("Invalid request format"))
		return
	}

	// Validate input
	if err := h.validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
		return
	}

	// Authenticate user
	token, err := h.authService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.NewAuthError("Authentication failed"))
		return
	}

	// Set secure HTTP-only cookie
	h.setAuthCookie(c, token.Token, token.ExpiresAt)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Login successful",
		"expires_at": token.ExpiresAt,
	})
}

// GoogleLogin initiates Google OAuth flow
func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	redirectURL, state := h.authService.HandleGoogleLogin(c.Request.Context())

	// Auto-detect HTTPS for secure flag
	secure := c.Request.TLS != nil

	// Set SameSite based on HTTPS
	sameSite := "Lax"
	if secure {
		sameSite = "Strict"
	}

	// Store state in secure cookie for CSRF protection
	c.SetCookie("oauth_state", state, 600, "/", "", secure, true) // 10 minutes

	// Set SameSite manually
	existingCookie := c.Writer.Header().Get("Set-Cookie")
	if existingCookie != "" {
		c.Writer.Header().Set("Set-Cookie", existingCookie+"; SameSite="+sameSite)
	}

	c.JSON(http.StatusOK, gin.H{
		"redirect_url": redirectURL,
		"state":        state,
	})
}

// GoogleCallback handles Google OAuth callback
func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	var req GoogleCallbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("Invalid request format"))
		return
	}

	// Validate input
	if err := h.validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
		return
	}

	// Verify state parameter (CSRF protection) - REQUIRED for security
	if req.State == "" {
		c.JSON(http.StatusBadRequest, errors.NewSecurityError("State parameter is required"))
		return
	}

	storedState, err := c.Cookie("oauth_state")
	if err != nil || storedState != req.State {
		c.JSON(http.StatusBadRequest, errors.NewSecurityError("Invalid state parameter"))
		return
	}

	// Auto-detect HTTPS for secure flag when clearing state cookie
	secure := c.Request.TLS != nil

	// Set SameSite based on HTTPS
	sameSite := "Lax"
	if secure {
		sameSite = "Strict"
	}

	// Clear state cookie with proper security attributes
	c.SetCookie("oauth_state", "", -1, "/", "", secure, true)

	// Set SameSite manually
	existingCookie := c.Writer.Header().Get("Set-Cookie")
	if existingCookie != "" {
		c.Writer.Header().Set("Set-Cookie", existingCookie+"; SameSite="+sameSite)
	}

	// Process OAuth callback
	token, err := h.authService.HandleGoogleCallback(c.Request.Context(), req.Code, req.State)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewAuthError("OAuth authentication failed"))
		return
	}

	// Set secure HTTP-only cookie
	h.setAuthCookie(c, token.Token, token.ExpiresAt)

	// Get user info for response
	user, err := h.authService.GetUserByID(c.Request.Context(), token.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewBusinessError("Failed to get user info", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"state": req.State,
	})
}

// GetCurrentUser retrieves the currently authenticated user's information.
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, errors.NewAuthError("User ID not found in context"))
		return
	}

	user, err := h.authService.GetUserByID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, errors.NewBusinessError("User not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}

// Logout clears the authentication cookie with auto-detection
func (h *AuthHandler) Logout(c *gin.Context) {
	// Auto-detect HTTPS for secure flag
	secure := c.Request.TLS != nil

	// Set SameSite based on HTTPS
	sameSite := "Lax"
	if secure {
		sameSite = "Strict"
	}

	// Clear auth cookie with proper security attributes
	c.SetCookie("ecolink_token", "", -1, "/", "", secure, true)

	// Set SameSite manually
	existingCookie := c.Writer.Header().Get("Set-Cookie")
	if existingCookie != "" {
		c.Writer.Header().Set("Set-Cookie", existingCookie+"; SameSite="+sameSite)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

// setAuthCookie sets a secure HTTP-only authentication cookie
func (h *AuthHandler) setAuthCookie(c *gin.Context, token string, expiresAt time.Time) {
	maxAge := int(time.Until(expiresAt).Seconds())

	// Auto-detect HTTPS for secure flag
	secure := c.Request.TLS != nil

	// Set secure cookie with proper attributes
	c.SetCookie(
		"ecolink_token", // name
		token,           // value
		maxAge,          // maxAge
		"/",             // path
		"",              // domain (empty for same-origin)
		secure,          // secure flag based on HTTPS
		true,            // httpOnly
	)

	// Set SameSite attribute manually via header
	sameSite := "Lax"
	if secure {
		sameSite = "Strict" // Stricter for HTTPS
	}

	// Get existing Set-Cookie header and append SameSite
	existingCookie := c.Writer.Header().Get("Set-Cookie")
	if existingCookie != "" {
		c.Writer.Header().Set("Set-Cookie", existingCookie+"; SameSite="+sameSite)
	}
}
