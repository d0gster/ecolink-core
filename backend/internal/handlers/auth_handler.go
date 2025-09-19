package handlers

import (
	"ecolink-core/internal/config"
	"ecolink-core/internal/security"
	"ecolink-core/internal/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.UserService
	jwtService  *security.JWTService
	cfg         *config.Config
}

type GoogleTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

type GoogleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func NewAuthHandler(cfg *config.Config, userService *services.UserService, jwtService *security.JWTService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		jwtService:  jwtService,
		cfg:         cfg,
	}
}

func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	var req struct {
		Code        string `json:"code" binding:"required"`
		RedirectURI string `json:"redirect_uri" binding:"required"`
		State       string `json:"state,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Exchange code for token
	token, err := h.exchangeCodeForToken(req.Code, req.RedirectURI)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	// Get user information
	userInfo, err := h.getUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
		return
	}

	// Debug: log user info
	fmt.Printf("Google User Info: ID=%s, Name=%s, Email=%s, Picture=%s\n",
		userInfo.ID, userInfo.Name, userInfo.Email, userInfo.Picture)

	// Register/update user in system
	user, err := h.userService.CreateOrUpdateUser(
		userInfo.ID,
		userInfo.Name,
		userInfo.Email,
		userInfo.Picture,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT token
	jwtToken, err := h.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set cookie attributes from config and environment
	domain := h.cfg.Cookie.Domain
	secure := h.cfg.Cookie.Secure

	// Parse SameSite option
	sameSite := h.getSameSiteMode()

	// Use gin's SetCookie (which doesn't accept SameSite) so we set header directly for SameSite
	cookie := &http.Cookie{
		Name:     "ecolink_token",
		Value:    jwtToken,
		Path:     "/",
		Domain:   domain,
		MaxAge:   3600 * 24 * 30,
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"state": req.State,
	})
}

func (h *AuthHandler) exchangeCodeForToken(code, redirectURI string) (*GoogleTokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", h.cfg.GoogleAuth.ClientID)
	data.Set("client_secret", h.cfg.GoogleAuth.ClientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", redirectURI)

	resp, err := http.Post("https://oauth2.googleapis.com/token",
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to exchange code for token: %s", string(bodyBytes))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResp GoogleTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func (h *AuthHandler) getUserInfo(accessToken string) (*GoogleUserInfo, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get user info: %s", string(bodyBytes))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo GoogleUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	user, err := h.userService.GetUser(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// Use config values for cookie clearing
	domain := h.cfg.Cookie.Domain
	secure := h.cfg.Cookie.Secure

	cookie := &http.Cookie{
		Name:     "ecolink_token",
		Value:    "",
		Path:     "/",
		Domain:   domain,
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   secure,
		SameSite: h.getSameSiteMode(),
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (h *AuthHandler) getSameSiteMode() http.SameSite {
	switch h.cfg.Cookie.SameSite {
	case "strict":
		return http.SameSiteStrictMode
	case "none":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteLaxMode
	}
}
