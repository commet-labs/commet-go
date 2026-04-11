package commet

import (
	"encoding/json"
	"testing"
)

func TestCommetErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		err  CommetError
		want string
	}{
		{
			name: "with code",
			err: CommetError{
				Message:    "Customer not found",
				Code:       "not_found",
				StatusCode: 404,
			},
			want: "commet: Customer not found (code=not_found, status=404)",
		},
		{
			name: "without code",
			err: CommetError{
				Message:    "Internal server error",
				StatusCode: 500,
			},
			want: "commet: Internal server error (status=500)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.want {
				t.Errorf("Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestValidationErrorMessage(t *testing.T) {
	err := &ValidationError{
		CommetError: CommetError{
			Message:    "Validation failed",
			Code:       "validation_error",
			StatusCode: 422,
		},
		ValidationErrors: map[string][]string{
			"email": {"is required", "must be valid"},
			"name":  {"is too short"},
		},
	}

	got := err.Error()
	want := "commet: validation error: Validation failed"
	if got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}

	if len(err.ValidationErrors["email"]) != 2 {
		t.Errorf("email errors count = %d, want 2", len(err.ValidationErrors["email"]))
	}
	if len(err.ValidationErrors["name"]) != 1 {
		t.Errorf("name errors count = %d, want 1", len(err.ValidationErrors["name"]))
	}
}

func TestHandleErrorParsesApiError(t *testing.T) {
	h := &httpClient{}

	t.Run("standard API error", func(t *testing.T) {
		data := map[string]any{
			"code":    "not_found",
			"message": "Customer not found",
		}

		err := h.handleError(404, data)
		commetErr, ok := err.(*CommetError)
		if !ok {
			t.Fatalf("expected *CommetError, got %T", err)
		}
		if commetErr.StatusCode != 404 {
			t.Errorf("StatusCode = %d, want 404", commetErr.StatusCode)
		}
		if commetErr.Code != "not_found" {
			t.Errorf("Code = %q, want not_found", commetErr.Code)
		}
		if commetErr.Message != "Customer not found" {
			t.Errorf("Message = %q, want Customer not found", commetErr.Message)
		}
	})

	t.Run("validation error with field details", func(t *testing.T) {
		data := map[string]any{
			"code":    "validation_error",
			"message": "Validation failed",
			"details": []any{
				map[string]any{"field": "email", "message": "is required"},
				map[string]any{"field": "email", "message": "must be valid email"},
				map[string]any{"field": "plan_code", "message": "plan not found"},
			},
		}

		err := h.handleError(422, data)
		valErr, ok := err.(*ValidationError)
		if !ok {
			t.Fatalf("expected *ValidationError, got %T", err)
		}
		if valErr.StatusCode != 422 {
			t.Errorf("StatusCode = %d, want 422", valErr.StatusCode)
		}
		if len(valErr.ValidationErrors["email"]) != 2 {
			t.Errorf("email errors = %d, want 2", len(valErr.ValidationErrors["email"]))
		}
		if len(valErr.ValidationErrors["plan_code"]) != 1 {
			t.Errorf("plan_code errors = %d, want 1", len(valErr.ValidationErrors["plan_code"]))
		}
	})

	t.Run("validation error with missing field defaults to unknown", func(t *testing.T) {
		data := map[string]any{
			"code":    "validation_error",
			"message": "Validation failed",
			"details": []any{
				map[string]any{"message": "something went wrong"},
			},
		}

		err := h.handleError(422, data)
		valErr, ok := err.(*ValidationError)
		if !ok {
			t.Fatalf("expected *ValidationError, got %T", err)
		}
		if len(valErr.ValidationErrors["unknown"]) != 1 {
			t.Errorf("unknown errors = %d, want 1", len(valErr.ValidationErrors["unknown"]))
		}
	})

	t.Run("error with no message uses fallback", func(t *testing.T) {
		data := map[string]any{
			"code": "server_error",
		}

		err := h.handleError(500, data)
		commetErr, ok := err.(*CommetError)
		if !ok {
			t.Fatalf("expected *CommetError, got %T", err)
		}
		if commetErr.Message != "Request failed with status 500" {
			t.Errorf("Message = %q, want fallback message", commetErr.Message)
		}
	})
}

func TestErrorFromJSON(t *testing.T) {
	t.Run("parses standard error JSON", func(t *testing.T) {
		jsonData := `{"code":"rate_limited","message":"Too many requests"}`
		var data map[string]any
		if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		h := &httpClient{}
		err := h.handleError(429, data)
		commetErr, ok := err.(*CommetError)
		if !ok {
			t.Fatalf("expected *CommetError, got %T", err)
		}
		if commetErr.Code != "rate_limited" {
			t.Errorf("Code = %q, want rate_limited", commetErr.Code)
		}
		if commetErr.StatusCode != 429 {
			t.Errorf("StatusCode = %d, want 429", commetErr.StatusCode)
		}
	})

	t.Run("parses validation error JSON with details", func(t *testing.T) {
		jsonData := `{"code":"validation_error","message":"Invalid input","details":[{"field":"customer_id","message":"is required"},{"field":"plan_code","message":"must be a valid plan"}]}`
		var data map[string]any
		if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		h := &httpClient{}
		err := h.handleError(422, data)
		valErr, ok := err.(*ValidationError)
		if !ok {
			t.Fatalf("expected *ValidationError, got %T", err)
		}
		if valErr.ValidationErrors["customer_id"][0] != "is required" {
			t.Errorf("customer_id error = %v, want 'is required'", valErr.ValidationErrors["customer_id"])
		}
		if valErr.ValidationErrors["plan_code"][0] != "must be a valid plan" {
			t.Errorf("plan_code error = %v, want 'must be a valid plan'", valErr.ValidationErrors["plan_code"])
		}
	})
}
