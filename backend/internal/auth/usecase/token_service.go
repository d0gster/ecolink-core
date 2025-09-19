package usecase

import (
	"crypto/rand"
	"ecolink-core/internal/auth/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(userID, email string) (*domain.AuthToken, error)
	ValidateToken(tokenString string) (*domain.TokenClaims, error)
}

type JWTTokenService struct {
	secretKey []byte
	issuer    string
}

func NewJWTTokenService(secretKey, issuer string) TokenService {
	return &JWTTokenService{
		secretKey: []byte(secretKey),
		issuer:    issuer,
	}
}

// GenerateToken creates a new JWT token with secure claims
func (s *JWTTokenService) GenerateToken(userID, email string) (*domain.AuthToken, error) {
	now := time.Now()
	expiresAt := now.Add(24 * time.Hour) // 24 hour expiration

	// Create secure random JTI (JWT ID) to prevent replay attacks
	jti := make([]byte, 16)
	if _, err := rand.Read(jti); err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"sub":    userID,
		"email":  email,
		"iss":    s.issuer,
		"iat":    now.Unix(),
		"exp":    expiresAt.Unix(),
		"jti":    string(jti), // Unique token ID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return nil, err
	}

	return &domain.AuthToken{
		Token:     tokenString,
		ExpiresAt: expiresAt,
		UserID:    userID,
	}, nil
}

// ValidateToken parses and validates a JWT token
func (s *JWTTokenService) ValidateToken(tokenString string) (*domain.TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Extract and validate claims
	userID, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid user ID claim")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid email claim")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid expiration claim")
	}

	iat, ok := claims["iat"].(float64)
	if !ok {
		return nil, errors.New("invalid issued at claim")
	}

	return &domain.TokenClaims{
		UserID: userID,
		Email:  email,
		Exp:    int64(exp),
		Iat:    int64(iat),
	}, nil
}