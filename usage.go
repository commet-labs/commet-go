package commet

import "context"

type CheckUsageAvailabilityParams struct {
	CustomerID     string `json:"customer_id"`
	FeatureCode    string `json:"feature_code"`
	Quantity       *int   `json:"quantity,omitempty"`
	IdempotencyKey string `json:"-"`
}

type TrackUsageParams struct {
	FeatureCode      string                           `json:"feature_code"`
	CustomerID       string                           `json:"customer_id"`
	EventID          *string                          `json:"event_id,omitempty"`
	Timestamp        *string                          `json:"timestamp,omitempty"`
	Properties       []TrackUsageParamsPropertiesItem `json:"properties,omitempty"`
	Model            *string                          `json:"model,omitempty"`
	InputTokens      *int                             `json:"input_tokens,omitempty"`
	OutputTokens     *int                             `json:"output_tokens,omitempty"`
	Value            *float64                         `json:"value,omitempty"`
	CacheReadTokens  *int                             `json:"cache_read_tokens,omitempty"`
	CacheWriteTokens *int                             `json:"cache_write_tokens,omitempty"`
	IdempotencyKey   string                           `json:"-"`
}

type SetUsageParams struct {
	CustomerID     string  `json:"customer_id"`
	FeatureCode    string  `json:"feature_code"`
	Value          int     `json:"value"`
	Reason         *string `json:"reason,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type UsageResource struct {
	http *httpClient
}

// Check if a customer can consume a feature before actual consumption. Returns availability and cost estimates based on the plan's consumption model.
func (r *UsageResource) Check(ctx context.Context, params *CheckUsageAvailabilityParams) (*UsageCheck, error) {
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"quantity":     params.Quantity,
	})
	return parseDirectResponse[UsageCheck](r.http.post(ctx, "/usage/check", body, params.IdempotencyKey))
}

// Track a usage event for a metered feature. Deducts from balance/credits if applicable.
func (r *UsageResource) Track(ctx context.Context, params *TrackUsageParams) (*UsageEvent, error) {
	body := buildBody(map[string]any{
		"feature_code":       params.FeatureCode,
		"customer_id":        params.CustomerID,
		"event_id":           params.EventID,
		"timestamp":          params.Timestamp,
		"properties":         params.Properties,
		"model":              params.Model,
		"input_tokens":       params.InputTokens,
		"output_tokens":      params.OutputTokens,
		"value":              params.Value,
		"cache_read_tokens":  params.CacheReadTokens,
		"cache_write_tokens": params.CacheWriteTokens,
	})
	return parseDirectResponse[UsageEvent](r.http.post(ctx, "/usage/events", body, params.IdempotencyKey))
}

// Set a metered feature's usage to an exact value for the current period. Use the Idempotency-Key header to make retries safe.
func (r *UsageResource) Set(ctx context.Context, params *SetUsageParams) (*UsageAdjustment, error) {
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"value":        params.Value,
		"reason":       params.Reason,
	})
	return parseDirectResponse[UsageAdjustment](r.http.put(ctx, "/usage", body, params.IdempotencyKey))
}
