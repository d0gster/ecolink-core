package usecase

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"ecolink-core/internal/auth/domain"
	"ecolink-core/internal/auth/repository"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo     repository.UserRepository
	tokenService TokenService
	googleConfig GoogleConfig
}

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
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

func NewAuthService(userRepo repository.UserRepository, tokenService TokenService, googleConfig GoogleConfig) *AuthService {
	return &AuthService{
		UserRepo:     userRepo,
		tokenService: tokenService,
		googleConfig: googleConfig,
	}
}

// Register creates a new user with local credentials
func (s *AuthService) Register(ctx context.Context, name, email, password string) (*domain.User, error) {
	// Check if user already exists
	existingUser, _, err := s.UserRepo.FindByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password using bcrypt (fixes MD5 vulnerability)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user and credential entities
	user := &domain.User{
		ID:    s.generateSecureID(),
		Email: email,
		Name:  name,
	}

	credential := &domain.Credential{
		UserID:       user.ID,
		PasswordHash: string(hashedPassword),
	}

	// Persist user
	if err := s.UserRepo.CreateUser(ctx, user, credential); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// Login validates credentials and generates a token
func (s *AuthService) Login(ctx context.Context, email, password string) (*domain.AuthToken, error) {
	// Retrieve user and credentials
	user, credential, err := s.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(credential.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.tokenService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

// HandleGoogleLogin initiates Google OAuth flow
func (s *AuthService) HandleGoogleLogin(ctx context.Context) (redirectURL string, state string) {
	state = s.generateSecureState()

	params := url.Values{}
	params.Add("client_id", s.googleConfig.ClientID)
	params.Add("redirect_uri", s.googleConfig.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", "email profile")
	params.Add("state", state)

	redirectURL = "https://accounts.google.com/o/oauth2/v2/auth?" + params.Encode()
	return redirectURL, state
}

// HandleGoogleCallback processes Google OAuth callback
func (s *AuthService) HandleGoogleCallback(ctx context.Context, code, state string) (*domain.AuthToken, error) {
	// Exchange code for token
	tokenResp, err := s.exchangeCodeForToken(code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	// Get user info from Google
	userInfo, err := s.getUserInfo(tokenResp.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	// Find or create user
	user, err := s.UserRepo.FindByProviderID(ctx, "google", userInfo.ID)
	if err != nil {
		// Create new user from social profile
		newUser := &domain.User{
			ID:      s.generateSecureID(),
			Email:   userInfo.Email,
			Name:    userInfo.Name,
			Picture: userInfo.Picture,
		}

		socialProfile := &domain.SocialProfile{
			UserID:     newUser.ID,
			Provider:   "google",
			ProviderID: userInfo.ID,
		}

		user, err = s.UserRepo.CreateUserFromSocial(ctx, newUser, socialProfile)
		if err != nil {
			return nil, fmt.Errorf("failed to create user from social: %w", err)
		}
	}

	// Generate application JWT token
	token, err := s.tokenService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (s *AuthService) exchangeCodeForToken(code string) (*GoogleTokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", s.googleConfig.ClientID)
	data.Set("client_secret", s.googleConfig.ClientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", s.googleConfig.RedirectURI)

	resp, err := http.Post("https://oauth2.googleapis.com/token",
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token exchange failed with status: %d", resp.StatusCode)
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

func (s *AuthService) getUserInfo(accessToken string) (*GoogleUserInfo, error) {
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
		return nil, fmt.Errorf("failed to get user info with status: %d", resp.StatusCode)
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

// generateSecureID creates a cryptographically secure ID (fixes MD5 vulnerability)
func (s *AuthService) generateSecureID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	hash := sha256.Sum256(bytes)
	return base64.URLEncoding.EncodeToString(hash[:16])
}

// GetUserByID retrieves a user by their ID
func (s *AuthService) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	return s.UserRepo.FindByID(ctx, userID)
}

// generateSecureState creates a secure state parameter for OAuth
func (s *AuthService) generateSecureState() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}
