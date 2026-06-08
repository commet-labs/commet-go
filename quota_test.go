package commet

import (
	"encoding/json"
	"testing"
)

// TestQuotaWireRoundTrip exercises the REAL deserialization path that http.go
// uses: the API sends camelCase on the wire, http.go runs convertKeys(toSnake)
// before populating rawApiResponse.Data, and only then does parseResponse run.
// The mock-based tests skip convertKeys, so this guards against json tags drifting
// out of snake_case (which silently empties every field).
func TestQuotaWireRoundTrip(t *testing.T) {
	t.Run("UsageQuota survives camel wire", func(t *testing.T) {
		wire := `{"featureCode":"tasks","current":5,"included":10,"remaining":5,"billedQuantity":2,"unlimited":false,"overageEnabled":true,"asOf":"2024-01-01"}`

		data := simulateWireResponse(t, wire)
		resp, err := parseResponse[UsageQuota](&rawApiResponse{Success: true, Data: data}, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.Data.FeatureCode != "tasks" {
			t.Errorf("FeatureCode = %q, want tasks (snake tags broken)", resp.Data.FeatureCode)
		}
		if resp.Data.Current != 5 {
			t.Errorf("Current = %v, want 5", resp.Data.Current)
		}
		if resp.Data.Included != 10 {
			t.Errorf("Included = %v, want 10", resp.Data.Included)
		}
		if resp.Data.Remaining == nil || *resp.Data.Remaining != 5 {
			t.Errorf("Remaining = %v, want 5", resp.Data.Remaining)
		}
		if resp.Data.BilledQuantity != 2 {
			t.Errorf("BilledQuantity = %v, want 2", resp.Data.BilledQuantity)
		}
		if !resp.Data.OverageEnabled {
			t.Error("expected OverageEnabled=true")
		}
		if resp.Data.AsOf == nil || *resp.Data.AsOf != "2024-01-01" {
			t.Errorf("AsOf = %v, want 2024-01-01", resp.Data.AsOf)
		}
	})

	t.Run("UsageQuotaEvent survives camel wire", func(t *testing.T) {
		wire := `{"id":"qe_1","customerId":"cus_1","featureCode":"tasks","previousBalance":4,"newBalance":5,"ts":"2024-01-01","createdAt":"2024-01-01"}`

		data := simulateWireResponse(t, wire)
		resp, err := parseResponse[UsageQuotaEvent](&rawApiResponse{Success: true, Data: data}, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.Data.CustomerID != "cus_1" {
			t.Errorf("CustomerID = %q, want cus_1 (snake tags broken)", resp.Data.CustomerID)
		}
		if resp.Data.FeatureCode != "tasks" {
			t.Errorf("FeatureCode = %q, want tasks", resp.Data.FeatureCode)
		}
		if resp.Data.PreviousBalance != 4 {
			t.Errorf("PreviousBalance = %d, want 4", resp.Data.PreviousBalance)
		}
		if resp.Data.NewBalance != 5 {
			t.Errorf("NewBalance = %d, want 5", resp.Data.NewBalance)
		}
		if resp.Data.CreatedAt != "2024-01-01" {
			t.Errorf("CreatedAt = %q, want 2024-01-01", resp.Data.CreatedAt)
		}
	})
}

// simulateWireResponse mirrors http.go: parse the camelCase wire body, snake-ize
// the keys, then re-marshal into the bytes that land in rawApiResponse.Data.
func simulateWireResponse(t *testing.T, wireBody string) []byte {
	t.Helper()
	var parsed map[string]any
	if err := json.Unmarshal([]byte(wireBody), &parsed); err != nil {
		t.Fatalf("failed to parse wire body: %v", err)
	}
	snaked := convertKeys(parsed, toSnake)
	data, err := json.Marshal(snaked)
	if err != nil {
		t.Fatalf("failed to re-marshal snaked data: %v", err)
	}
	return data
}

func TestParseUsageQuotaEvent(t *testing.T) {
	raw := &rawApiResponse{
		Success: true,
		Data:    []byte(`{"id":"qe_1","customer_id":"cus_1","feature_code":"tasks","previous_balance":4,"new_balance":5,"ts":"2024-01-01","created_at":"2024-01-01"}`),
	}

	resp, err := parseResponse[UsageQuotaEvent](raw, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data.FeatureCode != "tasks" {
		t.Errorf("FeatureCode = %v, want tasks", resp.Data.FeatureCode)
	}
	if resp.Data.PreviousBalance != 4 {
		t.Errorf("PreviousBalance = %v, want 4", resp.Data.PreviousBalance)
	}
	if resp.Data.NewBalance != 5 {
		t.Errorf("NewBalance = %v, want 5", resp.Data.NewBalance)
	}
}

func TestParseUsageQuota(t *testing.T) {
	t.Run("bounded allowance keeps remaining", func(t *testing.T) {
		raw := &rawApiResponse{
			Success: true,
			Data:    []byte(`{"feature_code":"tasks","current":5,"included":10,"remaining":5,"billed_quantity":2,"unlimited":false,"overage_enabled":true,"as_of":"2024-01-01"}`),
		}

		resp, err := parseResponse[UsageQuota](raw, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.Data.Remaining == nil || *resp.Data.Remaining != 5 {
			t.Errorf("Remaining = %v, want 5", resp.Data.Remaining)
		}
		if !resp.Data.OverageEnabled {
			t.Error("expected OverageEnabled=true")
		}
	})

	t.Run("unlimited allowance has null remaining", func(t *testing.T) {
		raw := &rawApiResponse{
			Success: true,
			Data:    []byte(`{"feature_code":"tasks","current":5,"included":0,"remaining":null,"unlimited":true,"overage_enabled":false,"as_of":null}`),
		}

		resp, err := parseResponse[UsageQuota](raw, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.Data.Remaining != nil {
			t.Errorf("Remaining = %v, want nil", resp.Data.Remaining)
		}
		if resp.Data.AsOf != nil {
			t.Errorf("AsOf = %v, want nil", resp.Data.AsOf)
		}
		if !resp.Data.Unlimited {
			t.Error("expected Unlimited=true")
		}
	})
}
