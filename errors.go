package commet

import "fmt"

// CommetError is the base error type for all SDK errors.
type CommetError struct {
	Message    string
	Code       string
	StatusCode int
	Details    any
}

func (e *CommetError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("commet: %s (code=%s, status=%d)", e.Message, e.Code, e.StatusCode)
	}
	return fmt.Sprintf("commet: %s (status=%d)", e.Message, e.StatusCode)
}

// ValidationError represents a validation failure from the API.
type ValidationError struct {
	CommetError
	ValidationErrors map[string][]string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("commet: validation error: %s", e.Message)
}
