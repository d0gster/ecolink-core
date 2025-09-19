package errors

import (
	"fmt"
	"net/http"
)

// APIError represents a structured API error response
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Status  int    `json:"-"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Error constructors for different error types

func NewValidationError(message string) *APIError {
	return &APIError{
		Code:    "VALIDATION_ERROR",
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func NewAuthError(message string) *APIError {
	return &APIError{
		Code:    "AUTHENTICATION_ERROR",
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *APIError {
	return &APIError{
		Code:    "AUTHORIZATION_ERROR",
		Message: message,
		Status:  http.StatusForbidden,
	}
}

func NewNotFoundError(resource string) *APIError {
	return &APIError{
		Code:    "RESOURCE_NOT_FOUND",
		Message: fmt.Sprintf("%s not found", resource),
		Status:  http.StatusNotFound,
	}
}

func NewBusinessError(message, details string) *APIError {
	return &APIError{
		Code:    "BUSINESS_ERROR",
		Message: message,
		Details: details,
		Status:  http.StatusConflict,
	}
}

func NewSecurityError(message string) *APIError {
	return &APIError{
		Code:    "SECURITY_ERROR",
		Message: message,
		Status:  http.StatusForbidden,
	}
}

func NewInternalError() *APIError {
	return &APIError{
		Code:    "INTERNAL_ERROR",
		Message: "An internal error occurred",
		Status:  http.StatusInternalServerError,
	}
}

func NewRateLimitError() *APIError {
	return &APIError{
		Code:    "RATE_LIMIT_EXCEEDED",
		Message: "Rate limit exceeded, please try again later",
		Status:  http.StatusTooManyRequests,
	}
}