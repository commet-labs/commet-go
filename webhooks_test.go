package commet

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func computeSignature(payload, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

func TestWebhooksVerify(t *testing.T) {
	w := &WebhooksResource{}
	secret := "whsec_test_secret_123"
	payload := `{"event":"subscription.created","data":{"id":"sub_123"}}`
	validSignature := computeSignature(payload, secret)

	tests := []struct {
		name      string
		payload   string
		signature string
		secret    string
		want      bool
	}{
		{
			name:      "valid signature",
			payload:   payload,
			signature: validSignature,
			secret:    secret,
			want:      true,
		},
		{
			name:      "invalid signature",
			payload:   payload,
			signature: "deadbeef1234567890abcdef",
			secret:    secret,
			want:      false,
		},
		{
			name:      "tampered payload",
			payload:   `{"event":"subscription.created","data":{"id":"sub_HACKED"}}`,
			signature: validSignature,
			secret:    secret,
			want:      false,
		},
		{
			name:      "wrong secret",
			payload:   payload,
			signature: validSignature,
			secret:    "whsec_wrong_secret",
			want:      false,
		},
		{
			name:      "empty payload",
			payload:   "",
			signature: validSignature,
			secret:    secret,
			want:      false,
		},
		{
			name:      "empty signature",
			payload:   payload,
			signature: "",
			secret:    secret,
			want:      false,
		},
		{
			name:      "empty secret",
			payload:   payload,
			signature: validSignature,
			secret:    "",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := w.Verify(tt.payload, tt.signature, tt.secret)
			if got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebhooksVerifyAndParse(t *testing.T) {
	w := &WebhooksResource{}
	secret := "whsec_test_secret_456"

	t.Run("valid signature parses a typed event", func(t *testing.T) {
		payload := `{"event":"subscription.created","organizationId":"org_1","data":{"subscriptionId":"sub_123","customerId":"cust_1","planId":"plan_1","planName":"Pro","status":"pending_payment","startDate":null,"name":null}}`
		signature := computeSignature(payload, secret)

		result, err := w.VerifyAndParse(payload, signature, secret)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Event != EventSubscriptionCreated {
			t.Errorf("event = %v, want subscription.created", result.Event)
		}
		data, err := result.AsSubscriptionCreated()
		if err != nil {
			t.Fatalf("AsSubscriptionCreated: %v", err)
		}
		if data.SubscriptionID != "sub_123" {
			t.Errorf("subscriptionId = %v, want sub_123", data.SubscriptionID)
		}
		if data.CustomerID != "cust_1" {
			t.Errorf("customerId = %v, want cust_1", data.CustomerID)
		}
	})

	t.Run("invalid signature returns error", func(t *testing.T) {
		payload := `{"event":"customer.created"}`
		_, err := w.VerifyAndParse(payload, "invalid_sig", secret)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		commetErr, ok := err.(*CommetError)
		if !ok {
			t.Fatalf("expected *CommetError, got %T", err)
		}
		if commetErr.Code != "INVALID_SIGNATURE" {
			t.Errorf("code = %v, want INVALID_SIGNATURE", commetErr.Code)
		}
	})

	t.Run("valid signature with invalid JSON returns error", func(t *testing.T) {
		payload := `not valid json`
		signature := computeSignature(payload, secret)

		_, err := w.VerifyAndParse(payload, signature, secret)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		commetErr, ok := err.(*CommetError)
		if !ok {
			t.Fatalf("expected *CommetError, got %T", err)
		}
		if commetErr.Code != "INVALID_JSON" {
			t.Errorf("code = %v, want INVALID_JSON", commetErr.Code)
		}
	})
}
