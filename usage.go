package commet

import (
	"context"
	"time"
)

type TrackUsageParams struct {
	Feature          string            `json:"feature"`
	CustomerID       string            `json:"customer_id"`
	Value            *int              `json:"value,omitempty"`
	Model            string            `json:"model,omitempty"`
	InputTokens      *int              `json:"input_tokens,omitempty"`
	OutputTokens     *int              `json:"output_tokens,omitempty"`
	CacheReadTokens  *int              `json:"cache_read_tokens,omitempty"`
	CacheWriteTokens *int              `json:"cache_write_tokens,omitempty"`
	IdempotencyKey   string            `json:"-"`
	Timestamp        string            `json:"timestamp,omitempty"`
	Properties       map[string]string `json:"properties,omitempty"`
}

type CheckUsageParams struct {
	CustomerID  string `json:"customer_id"`
	FeatureCode string `json:"feature_code"`
	Quantity    int    `json:"quantity"`
}

type UsageCheckResult struct {
	Allowed            bool    `json:"allowed"`
	ConsumptionModel   string  `json:"consumption_model"`
	Feature            string  `json:"feature"`
	Quantity           int     `json:"quantity"`
	Current            *int    `json:"current,omitempty"`
	Remaining          *int    `json:"remaining,omitempty"`
	Unlimited          *bool   `json:"unlimited,omitempty"`
	Included           *int    `json:"included,omitempty"`
	OverageEnabled     *bool   `json:"overage_enabled,omitempty"`
	OverageUnitPrice   *float64 `json:"overage_unit_price,omitempty"`
	CreditsPerUnit     *int    `json:"credits_per_unit,omitempty"`
	EstimatedCredits   *int    `json:"estimated_credits,omitempty"`
	PlanCredits        *int    `json:"plan_credits,omitempty"`
	PurchasedCredits   *int    `json:"purchased_credits,omitempty"`
	TotalCredits       *int    `json:"total_credits,omitempty"`
	UnitPrice          *float64 `json:"unit_price,omitempty"`
	EstimatedAmount    *float64 `json:"estimated_amount,omitempty"`
	CurrentBalance     *float64 `json:"current_balance,omitempty"`
	BlockOnExhaustion  *bool   `json:"block_on_exhaustion,omitempty"`
	Currency           string  `json:"currency,omitempty"`
	Reason             string  `json:"reason,omitempty"`
	Message            string  `json:"message,omitempty"`
}

type UsageResource struct {
	http *httpClient
}

func (r *UsageResource) Track(ctx context.Context, params *TrackUsageParams) (*ApiResponse[UsageEvent], error) {
	body := buildUsageBody(params)
	return parseResponse[UsageEvent](r.http.post(ctx, "/usage/events", body, params.IdempotencyKey))
}

func (r *UsageResource) Check(ctx context.Context, params *CheckUsageParams) (*ApiResponse[UsageCheckResult], error) {
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"quantity":     params.Quantity,
	})
	return parseResponse[UsageCheckResult](r.http.post(ctx, "/usage/check", body, ""))
}

type TrackModelTokensParams struct {
	Feature          string            `json:"feature"`
	CustomerID       string            `json:"customer_id"`
	Model            string            `json:"model"`
	InputTokens      int               `json:"input_tokens"`
	OutputTokens     int               `json:"output_tokens"`
	CacheReadTokens  *int              `json:"cache_read_tokens,omitempty"`
	CacheWriteTokens *int              `json:"cache_write_tokens,omitempty"`
	IdempotencyKey   string            `json:"-"`
	Timestamp        string            `json:"timestamp,omitempty"`
	Properties       map[string]string `json:"properties,omitempty"`
}

func (r *UsageResource) TrackModelTokens(ctx context.Context, params *TrackModelTokensParams) (*ApiResponse[UsageEvent], error) {
	var props []any
	if params.Properties != nil {
		props = make([]any, 0, len(params.Properties))
		for k, v := range params.Properties {
			props = append(props, map[string]any{"property": k, "value": v})
		}
	}

	timestamp := params.Timestamp
	if timestamp == "" {
		timestamp = time.Now().UTC().Format(time.RFC3339)
	}

	body := buildBody(map[string]any{
		"feature":            params.Feature,
		"customer_id":        params.CustomerID,
		"idempotency_key":    params.IdempotencyKey,
		"timestamp":          timestamp,
		"properties":         props,
		"model":              params.Model,
		"input_tokens":       params.InputTokens,
		"output_tokens":      params.OutputTokens,
		"cache_read_tokens":  params.CacheReadTokens,
		"cache_write_tokens": params.CacheWriteTokens,
	})

	return parseResponse[UsageEvent](r.http.post(ctx, "/usage/events", body, params.IdempotencyKey))
}

func buildUsageBody(params *TrackUsageParams) map[string]any {
	var props []any
	if params.Properties != nil {
		props = make([]any, 0, len(params.Properties))
		for k, v := range params.Properties {
			props = append(props, map[string]any{"property": k, "value": v})
		}
	}

	timestamp := params.Timestamp
	if timestamp == "" {
		timestamp = time.Now().UTC().Format(time.RFC3339)
	}

	body := buildBody(map[string]any{
		"feature":         params.Feature,
		"customer_id":     params.CustomerID,
		"idempotency_key": params.IdempotencyKey,
		"timestamp":       timestamp,
		"properties":      props,
	})

	if params.Model != "" {
		merged := buildBody(map[string]any{
			"model":              params.Model,
			"input_tokens":       params.InputTokens,
			"output_tokens":      params.OutputTokens,
			"cache_read_tokens":  params.CacheReadTokens,
			"cache_write_tokens": params.CacheWriteTokens,
		})
		for k, v := range merged {
			body[k] = v
		}
	} else if params.Value != nil {
		body["value"] = *params.Value
	}

	return body
}
