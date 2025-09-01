package middleware

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func AuthMiddleware(domain, audience string) gin.HandlerFunc {
	issuerURL, _ := url.Parse("https://" + domain + "/")
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, _ := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
	)

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}

		token, err := jwtValidator.ValidateToken(context.Background(), tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.(*validator.ValidatedClaims).CustomClaims.(*CustomClaims)
		c.Set("user_id", claims.Sub)
		c.Set("user_email", claims.Email)
		c.Set("user_name", claims.Name)
		c.Next()
	}
}