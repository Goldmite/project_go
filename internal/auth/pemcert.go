package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	certCache   string
	cacheMutex  sync.RWMutex
	lastFetched time.Time
)

// getPemCert fetches the current public key from Auth0's JWKS endpoint
// and returns it in PEM format. Implements caching to avoid repeated requests.
func getPemCert(ctx context.Context) string {
	domain := os.Getenv("AUTH0_DOMAIN")
	if domain == "" {
		return ""
	}

	// Check cache first (with read lock)
	cacheMutex.RLock()
	cachedCert := certCache
	lastFetch := lastFetched
	cacheMutex.RUnlock()

	// Return cached cert if it's fresh (less than 1 hour old)
	if cachedCert != "" && time.Since(lastFetch) < time.Hour {
		return cachedCert
	}

	// Fetch fresh certificate
	newCert, err := fetchJWKSCertificate(ctx, domain)
	if err != nil {
		// If we have a cached cert, use it even if stale
		if cachedCert != "" {
			return cachedCert
		}
		return ""
	}

	// Update cache (with write lock)
	cacheMutex.Lock()
	certCache = newCert
	lastFetched = time.Now()
	cacheMutex.Unlock()

	return newCert
}

// fetchJWKSCertificate fetches the JWKS and converts the first RSA key to PEM format
func fetchJWKSCertificate(ctx context.Context, domain string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://"+domain+"/.well-known/jwks.json", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var jwks struct {
		Keys []struct {
			Kty string   `json:"kty"`
			Use string   `json:"use"`
			Kid string   `json:"kid"`
			X5c []string `json:"x5c"`
			N   string   `json:"n"`
			E   string   `json:"e"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return "", fmt.Errorf("failed to decode JWKS: %w", err)
	}

	if len(jwks.Keys) == 0 {
		return "", errors.New("no keys found in JWKS")
	}

	// Prefer keys with x5c certificate chain (more reliable)
	for _, key := range jwks.Keys {
		if key.Kty == "RSA" && len(key.X5c) > 0 {
			return formatX5CAsPEM(key.X5c[0]), nil
		}
	}

	return "", errors.New("no valid RSA key found in JWKS")
}

// formatX5CAsPEM formats a certificate from the x5c field as PEM
func formatX5CAsPEM(x5c string) string {
	return fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----\n", x5c)
}
