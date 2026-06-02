package commet

import "testing"

func TestParseQuotaEvent(t *testing.T) {
	raw := &rawApiResponse{
		Success: true,
		Data:    []byte(`{"id":"qe_1","customerId":"cus_1","featureCode":"tasks","previousBalance":4,"newBalance":5,"ts":"2024-01-01","createdAt":"2024-01-01"}`),
	}

	resp, err := parseResponse[QuotaEvent](raw, nil)
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

func TestParseQuotaAllowance(t *testing.T) {
	t.Run("bounded allowance keeps remaining", func(t *testing.T) {
		raw := &rawApiResponse{
			Success: true,
			Data:    []byte(`{"featureCode":"tasks","current":5,"included":10,"remaining":5,"billedQuantity":2,"unlimited":false,"overageEnabled":true,"asOf":"2024-01-01"}`),
		}

		resp, err := parseResponse[QuotaAllowance](raw, nil)
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
			Data:    []byte(`{"featureCode":"tasks","current":5,"included":0,"remaining":null,"unlimited":true,"overageEnabled":false,"asOf":null}`),
		}

		resp, err := parseResponse[QuotaAllowance](raw, nil)
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
