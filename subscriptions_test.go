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

func TestGetActiveDeserializesRawSubscription(t *testing.T) {
	client, captured := newWireServer(t, http.StatusOK,
		`{"id":"sub_123","customerId":"cus_456","name":"Pro","status":"active","cancelAtPeriodEnd":false,"startDate":"2026-01-01"}`)

	resp, err := client.Subscriptions.GetActive(context.Background(), &GetActiveSubscriptionParams{
		CustomerID: "cus_456",
	})
	if err != nil {
		t.Fatalf("GetActive: %v", err)
	}

	if captured.Path != "/api/v1/subscriptions/active" {
		t.Errorf("path = %q, want /api/v1/subscriptions/active", captured.Path)
	}
	if resp == nil {
		t.Fatalf("resp = nil, want the deserialized subscription")
	}
	if resp.ID != "sub_123" {
		t.Errorf("resp.ID = %q, want sub_123", resp.ID)
	}
	if resp.CustomerID != "cus_456" {
		t.Errorf("resp.CustomerID = %q, want cus_456", resp.CustomerID)
	}
	if resp.Status != SubscriptionStatusActive {
		t.Errorf("resp.Status = %q, want active", resp.Status)
	}
	if resp.Name != "Pro" {
		t.Errorf("resp.Name = %q, want Pro", resp.Name)
	}
}

func TestGetActiveDeserializesNullData(t *testing.T) {
	client, _ := newWireServer(t, http.StatusOK, `null`)

	resp, err := client.Subscriptions.GetActive(context.Background(), &GetActiveSubscriptionParams{
		CustomerID: "cus_456",
	})
	if err != nil {
		t.Fatalf("GetActive: %v", err)
	}

	if resp != nil {
		t.Errorf("resp = %+v, want nil for a null response", resp)
	}
}
