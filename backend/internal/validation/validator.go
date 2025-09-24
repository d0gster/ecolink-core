package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()
	
	// Register custom validators
	v.RegisterValidation("shortcode", validateShortCode)
	v.RegisterValidation("url_safe", validateURLSafe)
	
	return &Validator{validate: v}
}

func (v *Validator) Validate(s interface{}) error {
	if err := v.validate.Struct(s); err != nil {
		return v.formatValidationError(err)
	}
	return nil
}

func (v *Validator) formatValidationError(err error) error {
	var messages []string
	
	for _, err := range err.(validator.ValidationErrors) {
		message := v.getErrorMessage(err)
		messages = append(messages, message)
	}
	
	return fmt.Errorf("%s", strings.Join(messages, "; "))
}

func (v *Validator) getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", fe.Field(), fe.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", fe.Field())
	case "shortcode":
		return fmt.Sprintf("%s must be a valid short code (6 alphanumeric characters)", fe.Field())
	case "url_safe":
		return fmt.Sprintf("%s contains unsafe characters", fe.Field())
	default:
		return fmt.Sprintf("%s is invalid", fe.Field())
	}
}

// Custom validators

// validateShortCode ensures short codes are exactly 6 alphanumeric characters
func validateShortCode(fl validator.FieldLevel) bool {
	code := fl.Field().String()
	if len(code) != 6 {
		return false
	}
	
	// Only allow alphanumeric characters (prevents path traversal)
	// Using compiled regex for better performance
	alphanumericPattern := regexp.MustCompile(`^[a-zA-Z0-9]{6}$`)
	return alphanumericPattern.MatchString(code)
}

// validateURLSafe ensures strings don't contain dangerous characters
func validateURLSafe(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	
	// Prevent common injection patterns
	dangerous := []string{
		"<script", "</script>", "javascript:", "data:",
		"../", "./", "\\", "<", ">", "\"", "'",
	}
	
	lowerValue := strings.ToLower(value)
	for _, pattern := range dangerous {
		if strings.Contains(lowerValue, pattern) {
			return false
		}
	}
	
	return true
}

// Compiled regex patterns for better performance
var (
	htmlTagPattern = regexp.MustCompile(`<[^>]*>`)
)

// SanitizeInput removes potentially dangerous characters from user input
func SanitizeInput(input string) string {
	// Remove HTML tags first using compiled regex
	input = htmlTagPattern.ReplaceAllString(input, "")
	
	// Escape HTML entities in correct order (& first to avoid double encoding)
	input = strings.ReplaceAll(input, "&", "&amp;")
	input = strings.ReplaceAll(input, "<", "&lt;")
	input = strings.ReplaceAll(input, ">", "&gt;")
	input = strings.ReplaceAll(input, "\"", "&quot;")
	input = strings.ReplaceAll(input, "'", "&#x27;")
	
	return strings.TrimSpace(input)
}