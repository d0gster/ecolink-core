package unit

import (
	"ecolink-core/internal/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_ShortCode(t *testing.T) {
	validator := validation.NewValidator()

	type TestStruct struct {
		Code string `validate:"shortcode"`
	}

	tests := []struct {
		name    string
		code    string
		wantErr bool
	}{
		{"valid code", "abc123", false},
		{"valid uppercase", "ABC123", false},
		{"valid mixed", "AbC123", false},
		{"too short", "abc12", true},
		{"too long", "abc1234", true},
		{"invalid chars", "abc-12", true},
		{"with spaces", "abc 12", true},
		{"with symbols", "abc@12", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TestStruct{Code: tt.code}
			err := validator.Validate(s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidator_URLSafe(t *testing.T) {
	validator := validation.NewValidator()

	type TestStruct struct {
		Value string `validate:"url_safe"`
	}

	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"safe string", "hello-world", false},
		{"alphanumeric", "abc123", false},
		{"script tag", "<script>alert(1)</script>", true},
		{"javascript", "javascript:alert(1)", true},
		{"path traversal", "../etc/passwd", true},
		{"html entities", "&lt;script&gt;", false}, // Already encoded, should be safe
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TestStruct{Value: tt.value}
			err := validator.Validate(s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"clean input", "hello world", "hello world"},
		{"html tags", "<script>alert(1)</script>", "alert(1)"},
		{"special chars", "<>&\"'", "&amp;&quot;&#x27;"},
		{"mixed content", "Hello <b>world</b>!", "Hello world!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validation.SanitizeInput(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}