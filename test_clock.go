package commet

import "context"

type AdvanceTestClockParams struct {
	AdvanceDays    *int    `json:"advance_days,omitempty"`
	FrozenTime     *string `json:"frozen_time,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type TestClockResource struct {
	http *httpClient
}

// Returns the organization's current test clock state. Sandbox only.
func (r *TestClockResource) Get(ctx context.Context) (*ApiResponse[TestClock], error) {
	return parseResponse[TestClock](r.http.get(ctx, "/test-clock", nil))
}

// Moves the test clock forward, by a number of days (advanceDays) or to an absolute instant (frozenTime). The clock can only move forward. Sandbox only.
func (r *TestClockResource) Advance(ctx context.Context, params *AdvanceTestClockParams) (*ApiResponse[TestClock], error) {
	body := buildBody(map[string]any{
		"advance_days": params.AdvanceDays,
		"frozen_time":  params.FrozenTime,
	})
	return parseResponse[TestClock](r.http.post(ctx, "/test-clock", body, params.IdempotencyKey))
}

// Discovers customers due for billing at the org's current (simulated) time and enqueues a billing cycle for each — renewals, expired trials, pending cancellations. Enqueueing is asynchronous. Sandbox only.
func (r *TestClockResource) ProcessBilling(ctx context.Context) (*ApiResponse[TestClockBilling], error) {
	return parseResponse[TestClockBilling](r.http.post(ctx, "/test-clock/process-billing", map[string]any{}, ""))
}
