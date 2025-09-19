package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateShortCode creates a secure short code using SHA-256 and crypto/rand
func GenerateShortCode(url string) string {
	// Generate cryptographically secure random bytes
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		// Fallback to deterministic but still secure approach
		return generateFallbackCode(url)
	}

	// Combine URL with random bytes and hash with SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hasher.Write(randomBytes)
	hash := hasher.Sum(nil)

	// Convert first 8 bytes to base62 for 6-character code
	return encodeBase62(hash[:8])
}

func generateFallbackCode(url string) string {
	// Secure fallback using only URL hash
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hash := hasher.Sum(nil)
	return encodeBase62(hash[:8])
}

func encodeBase62(input []byte) string {
	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[int(input[i%len(input)])%len(charset)]
	}
	return string(result)
}

// ValidateShortCode ensures the code meets security requirements
func ValidateShortCode(code string) error {
	if len(code) != 6 {
		return fmt.Errorf("invalid code length: expected 6, got %d", len(code))
	}
	
	for _, char := range code {
		valid := false
		for _, validChar := range charset {
			if char == validChar {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid character in code: %c", char)
		}
	}
	
	return nil
}
