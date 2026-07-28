package commet

import (
	"context"
	"testing"
)

// TestAdvanceOmitsUnsetOptionalsAndCamelizes verifies that AdvanceTestClock with
// only AdvanceDays set sends advanceDays (camelCase) and does NOT leak the unset
// FrozenTime pointer as null — the nil-pointer omission bug class.
func TestAdvanceOmitsUnsetOptionalsAndCamelizes(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"simulatedTime":"2026-07-01T00:00:00Z","isActive":true,"now":"2026-07-01T00:00:00Z","object":"test_clock","livemode":false}`)

	_, err := client.TestClock.Advance(context.Background(), &AdvanceTestClockParams{
		AdvanceDays: intPtr(7),
	})
	if err != nil {
		t.Fatalf("Advance: %v", err)
	}

	body := decodeBody(t, captured.Body)
	if body["advanceDays"] != float64(7) {
		t.Errorf("advanceDays = %v, want 7 (raw: %s)", body["advanceDays"], captured.Body)
	}
	if _, exists := body["advance_days"]; exists {
		t.Errorf("snake_case advance_days leaked on the wire: %s", captured.Body)
	}
	if v, exists := body["frozenTime"]; exists {
		t.Errorf("unset FrozenTime must be omitted, got %v (raw: %s)", v, captured.Body)
	}
}

// TestProcessBillingPostsEmptyBodyToCorrectPath guards the no-param POST: it must
// hit /test-clock/process-billing with POST and unmarshal the typed counters.
// This is behavioral (path + body shape + typed decode), not a bare "did it GET".
func TestProcessBillingPostsEmptyBodyToCorrectPath(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"customersFound":3,"enqueued":2,"failed":1,"object":"test_clock_billing","livemode":false}`)

	resp, err := client.TestClock.ProcessBilling(context.Background())
	if err != nil {
		t.Fatalf("ProcessBilling: %v", err)
	}

	if captured.Method != "POST" {
		t.Errorf("method = %s, want POST", captured.Method)
	}
	if captured.Path != "/api/v1/test-clock/process-billing" {
		t.Errorf("path = %s, want /api/v1/test-clock/process-billing", captured.Path)
	}
	// Empty map body marshals to "{}" — must not be null or carry stray keys.
	body := decodeBody(t, captured.Body)
	if len(body) != 0 {
		t.Errorf("expected empty JSON body, got %s", captured.Body)
	}

	if resp.CustomersFound != 3 || resp.Enqueued != 2 || resp.Failed != 1 {
		t.Errorf("counters = (%d,%d,%d), want (3,2,1)", resp.CustomersFound, resp.Enqueued, resp.Failed)
	}
}

// TestGetTestClockUnmarshalsNullableSimulatedTime checks a camelCase response with
// an explicit null simulatedTime maps to a nil *string on the typed struct.
func TestGetTestClockUnmarshalsNullableSimulatedTime(t *testing.T) {
	client, _ := newWireServer(t, 200, `{"simulatedTime":null,"isActive":false,"now":"2026-06-08T00:00:00Z","object":"test_clock","livemode":false}`)

	resp, err := client.TestClock.Get(context.Background())
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if resp.SimulatedTime != nil {
		t.Errorf("SimulatedTime = %v, want nil for wire null", *resp.SimulatedTime)
	}
	if resp.IsActive {
		t.Error("IsActive = true, want false")
	}
	if resp.Now != "2026-06-08T00:00:00Z" {
		t.Errorf("Now = %q, want 2026-06-08T00:00:00Z", resp.Now)
	}
}
