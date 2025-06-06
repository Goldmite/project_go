package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

var (
	// ErrNoTokenFound is returned when no JWT is found in the request
	ErrNoTokenFound = errors.New("no token found")
)

// CustomClaims contains custom data we want from the token
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate implements validator.CustomClaims
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a Gin middleware that validates JWT tokens
func EnsureValidToken() gin.HandlerFunc {
	issuerURL := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
	audience := os.Getenv("AUTH0_AUDIENCE")
	keyFunc := func(ctx context.Context) (interface{}, error) {
		cert := getPemCert(ctx)
		if cert == "" {
			return nil, errors.New("failed to fetch PEM certificate")
		}
		// Return the certificate as interface{} instead of []byte
		return cert, nil
	}

	jwtValidator, err := validator.New(
		keyFunc,
		validator.RS256,
		issuerURL,
		[]string{audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
	)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		// Get the token from the request
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization header format",
			})
			return
		}

		token := parts[1]

		// Validate the token
		_, err := jwtValidator.ValidateToken(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid token",
				"details": err.Error(),
			})
			return
		}

		// Token is valid, continue
		c.Next()
	}
}

// getPemCert fetches the PEM certificate from Auth0 (implementation similar to before)
func getPemCert(ctx context.Context) string {
	// Implement your actual certificate fetching logic here
	// This is a simplified example - in production you should:
	// 1. Cache the certificate
	// 2. Handle errors properly
	// 3. Verify the certificate chain

	domain := os.Getenv("AUTH0_DOMAIN")
	if domain == "" {
		return ""
	}

	// In a real implementation, you would:
	// 1. Fetch the JWKS from https://{domain}/.well-known/jwks.json
	// 2. Parse the JWKS to get the correct key
	// 3. Convert it to PEM format

	// For now, return an empty string (implementation left as exercise)
	return ""
}
