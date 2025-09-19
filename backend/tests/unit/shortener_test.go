package unit

import (
	"ecolink-core/pkg/utils"
	"testing"
)

func TestGenerateShortCode(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"Valid HTTP URL", "http://example.com"},
		{"Valid HTTPS URL", "https://example.com"},
		{"Long URL", "https://example.com/very/long/path/with/many/segments"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := utils.GenerateShortCode(tt.url)
			
			// Test code length
			if len(code) != 6 {
				t.Errorf("Expected code length 6, got %d", len(code))
			}
			
			// Test code validation
			if err := utils.ValidateShortCode(code); err != nil {
				t.Errorf("Generated code failed validation: %v", err)
			}
		})
	}
}

func TestGenerateShortCodeUniqueness(t *testing.T) {
	url := "https://example.com"
	codes := make(map[string]bool)
	
	// Generate 1000 codes and check for uniqueness
	for i := 0; i < 1000; i++ {
		code := utils.GenerateShortCode(url)
		if codes[code] {
			t.Errorf("Duplicate code generated: %s", code)
		}
		codes[code] = true
	}
}

func TestValidateShortCode(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		wantErr bool
	}{
		{"Valid code", "abc123", false},
		{"Valid uppercase", "ABC123", false},
		{"Empty code", "", true},
		{"Too short", "abc12", true},
		{"Too long", "abc1234", true},
		{"Invalid characters", "abc-12", true},
		{"Special characters", "abc@12", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ValidateShortCode(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateShortCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}