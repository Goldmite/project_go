package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// AuthHandler contains authentication-related handlers
type AuthHandler struct {
	// You can add dependencies here (e.g., database connections)
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// LoginRedirect provides the Auth0 login URL for frontend redirection
func (h *AuthHandler) LoginRedirect(c *gin.Context) {
	domain := "https://" + os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
	audience := os.Getenv("AUTH0_AUDIENCE")
	redirectURI := os.Getenv("AUTH0_CALLBACK_URL")

	loginURL := fmt.Sprintf(
		"%s/authorize?response_type=code&client_id=%s&redirect_uri=%s&audience=%s&scope=openid profile email",
		domain,
		clientID,
		redirectURI,
		audience,
	)

	c.JSON(http.StatusOK, gin.H{
		"login_url": loginURL,
	})
}

// UserProfile returns the user profile from the JWT claims
func (h *AuthHandler) UserProfile(c *gin.Context) {
	// This assumes you've stored the validated claims in context
	claims, exists := c.Get("jwt_claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No user information available",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profile": claims,
	})
}

// ProtectedPing is an example protected endpoint
func (h *AuthHandler) ProtectedPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "You have accessed a protected endpoint!",
		"status":  "success",
	})
}
