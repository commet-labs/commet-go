package commet

import "context"

type ProcessTestClockBillingParams struct {
	IdempotencyKey string `json:"-"`
}

type AdvanceTestClockParams struct {
	AdvanceDays    *int    `json:"advance_days,omitempty"`
	FrozenTime     *string `json:"frozen_time,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type TestClockResource struct {
	http *httpClient
}

// Deprecated. POST /test-clock now advances time and processes every due billing deadline in one durable run.
// Deprecated.
func (r *TestClockResource) ProcessBilling(ctx context.Context, params *ProcessTestClockBillingParams) (*struct{}, error) {
	return parseDirectResponse[struct{}](r.http.post(ctx, "/test-clock/process-billing", nil, params.IdempotencyKey))
}

// Returns the organization's current test clock state and latest durable run. Sandbox only.
func (r *TestClockResource) Get(ctx context.Context) (*TestClock, error) {
	return parseDirectResponse[TestClock](r.http.get(ctx, "/test-clock", nil))
}

// Starts a durable run that moves the test clock forward and processes every billing deadline due before the target time. Poll GET /test-clock for progress and terminal results. Sandbox only.
func (r *TestClockResource) Advance(ctx context.Context, params *AdvanceTestClockParams) (*TestClockRun, error) {
	body := buildBody(map[string]any{
		"advance_days": params.AdvanceDays,
		"frozen_time":  params.FrozenTime,
	})
	return parseDirectResponse[TestClockRun](r.http.post(ctx, "/test-clock", body, params.IdempotencyKey))
}
