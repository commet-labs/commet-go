package commet

import "fmt"

type CommetError struct {
	Message    string `json:"message"`
	Code       string `json:"code,omitempty"`
	Type       string `json:"type,omitempty"`
	StatusCode int    `json:"status_code"`
	Param      string `json:"param,omitempty"`
	DocURL     string `json:"doc_url,omitempty"`
	RequestID  string `json:"request_id,omitempty"`
	Details    any    `json:"details,omitempty"`
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
	ValidationErrors map[string][]string `json:"validation_errors"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("commet: validation error: %s", e.Message)
}
