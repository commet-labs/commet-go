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

type UsageResource struct {
	http *httpClient
}

func (r *UsageResource) Track(ctx context.Context, params *TrackUsageParams) (*ApiResponse[UsageEvent], error) {
	body := buildUsageBody(params)
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
