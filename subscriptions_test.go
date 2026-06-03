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
	_, err = client.Subscriptions.ChangePlan(context.Background(), &ChangePlanParams{
		SubscriptionID: "sub_123",
		NewPlanID:      "plan_pro",
		SuccessURL:     wantSuccessURL,
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
