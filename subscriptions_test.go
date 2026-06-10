package commet

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestChangePlanSendsSuccessURLAsCamelCase verifies that ChangePlan, when given a
// SuccessURL, emits it on the wire as the camelCase key "successUrl". The struct
// declares the snake_case json tag "success_url", and http.go runs
// convertKeys(body, toCamel) before sending, so the wire body must use camelCase.
// This guards both that the SuccessURL param is actually wired into the request and
// that the snake->camel conversion happened (no leftover snake_case keys).
func TestChangePlanSendsSuccessURLAsCamelCase(t *testing.T) {
	var capturedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":{"id":"sub_123"}}`))
	}))
	defer server.Close()

	client, err := New("ck_test_123", WithTelemetry(false), WithRetries(0))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	// Point the client at the fake server. The SDK builds requests against
	// http.baseURL, so override it to the test server (mirrors the real layout:
	// baseURL + "/api/v1" + endpoint).
	client.http.baseURL = server.URL + "/api/v1"

	const wantSuccessURL = "https://app.example.com/billing/success"
	_, err = client.Subscriptions.ChangePlan(context.Background(), "sub_123", &ChangePlanParams{
		NewPlanID:  strPtr("plan_pro"),
		SuccessURL: strPtr(wantSuccessURL),
	})
	if err != nil {
		t.Fatalf("ChangePlan: %v", err)
	}

	var body map[string]any
	if err := json.Unmarshal(capturedBody, &body); err != nil {
		t.Fatalf("unmarshal request body: %v (raw: %s)", err, capturedBody)
	}

	if body["successUrl"] != wantSuccessURL {
		t.Errorf("body[\"successUrl\"] = %v, want %q (raw: %s)", body["successUrl"], wantSuccessURL, capturedBody)
	}

	if _, exists := body["success_url"]; exists {
		t.Errorf("expected no snake_case \"success_url\" key on the wire (conversion to camelCase did not happen), raw: %s", capturedBody)
	}

	if body["newPlanId"] != "plan_pro" {
		t.Errorf("body[\"newPlanId\"] = %v, want plan_pro (raw: %s)", body["newPlanId"], capturedBody)
	}
}

// TestCreateSendsIntroOfferAsNestedCamelCase verifies that Create, when given an
// IntroOffer, emits it on the wire as a nested object under the camelCase key
// "introOffer" with camelCase inner keys "discountType", "discountValue" and
// "durationCycles". The nested object is a typed struct with snake_case json tags;
// http.go genericizes the body (JSON round-trip) before convertKeys(body, toCamel)
// so the nested keys re-key correctly. This guards both that the field is wired into
// the request and that the snake->camel conversion reached the nested level (no
// leftover snake_case keys).
func TestCreateSendsIntroOfferAsNestedCamelCase(t *testing.T) {
	var capturedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":{"id":"sub_123"}}`))
	}))
	defer server.Close()

	client, err := New("ck_test_123", WithTelemetry(false), WithRetries(0))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	client.http.baseURL = server.URL + "/api/v1"

	_, err = client.Subscriptions.Create(context.Background(), &CreateSubscriptionParams{
		CustomerID: "cus_123",
		PlanCode:   strPtr("pro"),
		IntroOffer: &CreateSubscriptionParamsIntroOffer{
			DiscountType:   DiscountTypePercentage,
			DiscountValue:  2000,
			DurationCycles: 3,
		},
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	var body map[string]any
	if err := json.Unmarshal(capturedBody, &body); err != nil {
		t.Fatalf("unmarshal request body: %v (raw: %s)", err, capturedBody)
	}

	if _, exists := body["intro_offer"]; exists {
		t.Errorf("expected no snake_case \"intro_offer\" key on the wire (conversion to camelCase did not happen), raw: %s", capturedBody)
	}

	offer, ok := body["introOffer"].(map[string]any)
	if !ok {
		t.Fatalf("body[\"introOffer\"] = %v, want nested object (raw: %s)", body["introOffer"], capturedBody)
	}

	if offer["discountType"] != "percentage" {
		t.Errorf("customIntroOffer[\"discountType\"] = %v, want percentage (raw: %s)", offer["discountType"], capturedBody)
	}
	if offer["discountValue"] != float64(2000) {
		t.Errorf("customIntroOffer[\"discountValue\"] = %v, want 2000 (raw: %s)", offer["discountValue"], capturedBody)
	}
	if offer["durationCycles"] != float64(3) {
		t.Errorf("customIntroOffer[\"durationCycles\"] = %v, want 3 (raw: %s)", offer["durationCycles"], capturedBody)
	}

	if _, exists := offer["discount_type"]; exists {
		t.Errorf("expected no snake_case \"discount_type\" key inside customIntroOffer (nested conversion did not happen), raw: %s", capturedBody)
	}
	if _, exists := offer["discount_value"]; exists {
		t.Errorf("expected no snake_case \"discount_value\" key inside customIntroOffer (nested conversion did not happen), raw: %s", capturedBody)
	}
	if _, exists := offer["duration_cycles"]; exists {
		t.Errorf("expected no snake_case \"duration_cycles\" key inside customIntroOffer (nested conversion did not happen), raw: %s", capturedBody)
	}
}

// TestGetActiveDeserializesRawSubscription verifies that GetActive hydrates the
// real wire shape — a raw camelCase Subscription object under "data", never a
// {"value": ...} wrapper — into a non-nil *Subscription with its fields mapped.
func TestGetActiveDeserializesRawSubscription(t *testing.T) {
	client, captured := newWireServer(t, http.StatusOK,
		`{"success":true,"data":{"id":"sub_123","customerId":"cus_456","name":"Pro","status":"active","cancelAtPeriodEnd":false,"startDate":"2026-01-01"}}`)

	resp, err := client.Subscriptions.GetActive(context.Background(), &GetActiveSubscriptionParams{
		CustomerID: "cus_456",
	})
	if err != nil {
		t.Fatalf("GetActive: %v", err)
	}

	if captured.Path != "/api/v1/subscriptions/active" {
		t.Errorf("path = %q, want /api/v1/subscriptions/active", captured.Path)
	}
	if !resp.Success {
		t.Errorf("resp.Success = false, want true")
	}
	if resp.Data == nil {
		t.Fatalf("resp.Data = nil, want the deserialized subscription")
	}
	if resp.Data.ID != "sub_123" {
		t.Errorf("resp.Data.ID = %q, want sub_123", resp.Data.ID)
	}
	if resp.Data.CustomerID != "cus_456" {
		t.Errorf("resp.Data.CustomerID = %q, want cus_456", resp.Data.CustomerID)
	}
	if resp.Data.Status != SubscriptionStatusActive {
		t.Errorf("resp.Data.Status = %q, want active", resp.Data.Status)
	}
	if resp.Data.Name != "Pro" {
		t.Errorf("resp.Data.Name = %q, want Pro", resp.Data.Name)
	}
}

// TestGetActiveDeserializesNullData verifies that a "data": null response (no
// active subscription) parses into a nil Data pointer instead of failing or
// fabricating a zero-value subscription.
func TestGetActiveDeserializesNullData(t *testing.T) {
	client, _ := newWireServer(t, http.StatusOK, `{"success":true,"data":null}`)

	resp, err := client.Subscriptions.GetActive(context.Background(), &GetActiveSubscriptionParams{
		CustomerID: "cus_456",
	})
	if err != nil {
		t.Fatalf("GetActive: %v", err)
	}

	if !resp.Success {
		t.Errorf("resp.Success = false, want true")
	}
	if resp.Data != nil {
		t.Errorf("resp.Data = %+v, want nil for a null data response", resp.Data)
	}
}
