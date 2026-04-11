package commet

import "testing"

func TestToSnake(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"camelCase", "camel_case"},
		{"customerId", "customer_id"},
		{"billingEmail", "billing_email"},
		{"APIKey", "api_key"},
		{"HTMLParser", "html_parser"},
		{"getHTTPResponse", "get_http_response"},
		{"already_snake", "already_snake"},
		{"single", "single"},
		{"A", "a"},
		{"", ""},
		{"userID", "user_id"},
		{"hasMore", "has_more"},
		{"nextCursor", "next_cursor"},
		{"isActive", "is_active"},
		{"billingDayOfMonth", "billing_day_of_month"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toSnake(tt.input)
			if got != tt.want {
				t.Errorf("toSnake(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"snake_case", "snakeCase"},
		{"customer_id", "customerId"},
		{"billing_email", "billingEmail"},
		{"already", "already"},
		{"a", "a"},
		{"", ""},
		{"has_more", "hasMore"},
		{"next_cursor", "nextCursor"},
		{"is_active", "isActive"},
		{"billing_day_of_month", "billingDayOfMonth"},
		{"full_name", "fullName"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toCamel(tt.input)
			if got != tt.want {
				t.Errorf("toCamel(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestConvertKeys(t *testing.T) {
	t.Run("nested map conversion", func(t *testing.T) {
		input := map[string]any{
			"customerId": "cust_123",
			"billingEmail": "test@example.com",
			"address": map[string]any{
				"postalCode": "12345",
				"countryCode": "US",
			},
		}

		result := convertKeys(input, toSnake)
		converted, ok := result.(map[string]any)
		if !ok {
			t.Fatal("expected map[string]any")
		}
		if converted["customer_id"] != "cust_123" {
			t.Errorf("customer_id = %v, want cust_123", converted["customer_id"])
		}
		if converted["billing_email"] != "test@example.com" {
			t.Errorf("billing_email = %v, want test@example.com", converted["billing_email"])
		}
		addr, ok := converted["address"].(map[string]any)
		if !ok {
			t.Fatal("address is not a map")
		}
		if addr["postal_code"] != "12345" {
			t.Errorf("postal_code = %v, want 12345", addr["postal_code"])
		}
	})

	t.Run("array conversion", func(t *testing.T) {
		input := []any{
			map[string]any{"firstName": "Alice"},
			map[string]any{"firstName": "Bob"},
		}

		result := convertKeys(input, toSnake)
		converted, ok := result.([]any)
		if !ok {
			t.Fatal("expected []any")
		}
		if len(converted) != 2 {
			t.Fatalf("len = %d, want 2", len(converted))
		}
		first, ok := converted[0].(map[string]any)
		if !ok {
			t.Fatal("first element is not a map")
		}
		if first["first_name"] != "Alice" {
			t.Errorf("first_name = %v, want Alice", first["first_name"])
		}
	})

	t.Run("scalar passthrough", func(t *testing.T) {
		if convertKeys("hello", toSnake) != "hello" {
			t.Error("string passthrough failed")
		}
		if convertKeys(42, toSnake) != 42 {
			t.Error("int passthrough failed")
		}
		if convertKeys(nil, toSnake) != nil {
			t.Error("nil passthrough failed")
		}
	})
}

func TestBuildBody(t *testing.T) {
	t.Run("filters nil and empty strings", func(t *testing.T) {
		body := buildBody(map[string]any{
			"email":    "test@example.com",
			"name":     "",
			"metadata": nil,
			"active":   true,
		})

		if body["email"] != "test@example.com" {
			t.Errorf("email = %v, want test@example.com", body["email"])
		}
		if body["active"] != true {
			t.Errorf("active = %v, want true", body["active"])
		}
		if _, exists := body["name"]; exists {
			t.Error("name should be filtered out")
		}
		if _, exists := body["metadata"]; exists {
			t.Error("metadata should be filtered out")
		}
	})

	t.Run("keeps zero values that are not nil or empty string", func(t *testing.T) {
		body := buildBody(map[string]any{
			"count": 0,
			"flag":  false,
		})

		if body["count"] != 0 {
			t.Errorf("count = %v, want 0", body["count"])
		}
		if body["flag"] != false {
			t.Errorf("flag = %v, want false", body["flag"])
		}
	})
}

func TestParseResponse(t *testing.T) {
	t.Run("parses typed data from raw response", func(t *testing.T) {
		raw := &rawApiResponse{
			Success: true,
			Data:    []byte(`{"id":"cust_123","billing_email":"test@example.com","is_active":true,"created_at":"2024-01-01","updated_at":"2024-01-01"}`),
			Code:    "ok",
			Message: "Customer created",
		}

		resp, err := parseResponse[Customer](raw, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !resp.Success {
			t.Error("expected success=true")
		}
		if resp.Data.ID != "cust_123" {
			t.Errorf("ID = %v, want cust_123", resp.Data.ID)
		}
		if resp.Data.BillingEmail != "test@example.com" {
			t.Errorf("BillingEmail = %v, want test@example.com", resp.Data.BillingEmail)
		}
		if !resp.Data.IsActive {
			t.Error("expected IsActive=true")
		}
	})

	t.Run("parses list data", func(t *testing.T) {
		raw := &rawApiResponse{
			Success:    true,
			Data:       []byte(`[{"id":"cust_1","billing_email":"a@b.com","is_active":true,"created_at":"2024-01-01","updated_at":"2024-01-01"},{"id":"cust_2","billing_email":"c@d.com","is_active":false,"created_at":"2024-01-01","updated_at":"2024-01-01"}]`),
			HasMore:    true,
			NextCursor: "cursor_abc",
		}

		resp, err := parseResponse[[]Customer](raw, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2", len(resp.Data))
		}
		if resp.Data[0].ID != "cust_1" {
			t.Errorf("Data[0].ID = %v, want cust_1", resp.Data[0].ID)
		}
		if !resp.HasMore {
			t.Error("expected HasMore=true")
		}
		if resp.NextCursor != "cursor_abc" {
			t.Errorf("NextCursor = %v, want cursor_abc", resp.NextCursor)
		}
	})

	t.Run("propagates error from http layer", func(t *testing.T) {
		_, err := parseResponse[Customer](nil, &CommetError{
			Message:    "Not found",
			StatusCode: 404,
			Code:       "not_found",
		})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("handles empty data", func(t *testing.T) {
		raw := &rawApiResponse{
			Success: true,
			Message: "No content",
		}

		resp, err := parseResponse[Customer](raw, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !resp.Success {
			t.Error("expected success=true")
		}
	})
}
