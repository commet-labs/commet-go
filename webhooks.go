package commet

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// WebhookEndpoint, WebhookEndpointCreated, and WebhookTestResult are preserved
// here: webhook endpoint management keeps its hand-written HMAC verification and
// signing logic, so the generator does not emit these response models.
type WebhookEndpoint struct {
	ID          string   `json:"id"`
	Object      string   `json:"object"`
	Livemode    bool     `json:"livemode"`
	URL         string   `json:"url"`
	Events      []string `json:"events"`
	Description string   `json:"description,omitempty"`
	IsActive    bool     `json:"is_active"`
	ApiVersion  string   `json:"api_version,omitempty"`
	CreatedAt   string   `json:"created_at"`
}

type WebhookEndpointCreated struct {
	WebhookEndpoint
	SecretKey string `json:"secret_key"`
}

type WebhookTestResult struct {
	Success     bool   `json:"success"`
	DeliveryId  string `json:"delivery_id"`
	DeliveredAt string `json:"delivered_at"`
}

type ListWebhooksParams struct {
	Limit  *int   `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type CreateWebhookParams struct {
	URL         string   `json:"url"`
	Events      []string `json:"events"`
	Description string   `json:"description,omitempty"`
	ApiVersion  string   `json:"api_version,omitempty"`
}

type UpdateWebhookParams struct {
	URL         string   `json:"url,omitempty"`
	Events      []string `json:"events,omitempty"`
	Description string   `json:"description,omitempty"`
	IsActive    *bool    `json:"is_active,omitempty"`
	ApiVersion  string   `json:"api_version,omitempty"`
}

type WebhooksResource struct {
	http *httpClient
}

func (w *WebhooksResource) Verify(payload string, signature string, secret string) bool {
	if payload == "" || signature == "" || secret == "" {
		return false
	}

	expected := sign(payload, secret)
	return hmac.Equal([]byte(signature), []byte(expected))
}

func (w *WebhooksResource) VerifyAndParse(rawBody string, signature string, secret string) (*WebhookEvent, error) {
	if !w.Verify(rawBody, signature, secret) {
		return nil, &CommetError{
			Message: "Invalid webhook signature",
			Code:    "INVALID_SIGNATURE",
		}
	}

	var event WebhookEvent
	if err := json.Unmarshal([]byte(rawBody), &event); err != nil {
		return nil, &CommetError{
			Message: "Failed to parse webhook payload",
			Code:    "INVALID_JSON",
		}
	}

	return &event, nil
}

func (w *WebhooksResource) List(ctx context.Context, params *ListWebhooksParams) (*ApiResponse[[]WebhookEndpoint], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]WebhookEndpoint](w.http.get(ctx, "/webhooks", queryParams))
}

func (w *WebhooksResource) Create(ctx context.Context, params *CreateWebhookParams) (*ApiResponse[WebhookEndpointCreated], error) {
	body := buildBody(map[string]any{
		"url":         params.URL,
		"events":      params.Events,
		"description": params.Description,
		"api_version": params.ApiVersion,
	})
	return parseResponse[WebhookEndpointCreated](w.http.post(ctx, "/webhooks", body, ""))
}

func (w *WebhooksResource) Get(ctx context.Context, webhookID string) (*ApiResponse[WebhookEndpoint], error) {
	return parseResponse[WebhookEndpoint](w.http.get(ctx, fmt.Sprintf("/webhooks/%s", webhookID), nil))
}

func (w *WebhooksResource) Update(ctx context.Context, webhookID string, params *UpdateWebhookParams) (*ApiResponse[WebhookEndpoint], error) {
	body := buildBody(map[string]any{
		"url":         params.URL,
		"events":      params.Events,
		"description": params.Description,
		"api_version": params.ApiVersion,
	})
	if params.IsActive != nil {
		body["is_active"] = *params.IsActive
	}
	return parseResponse[WebhookEndpoint](w.http.put(ctx, fmt.Sprintf("/webhooks/%s", webhookID), body, ""))
}

func (w *WebhooksResource) Delete(ctx context.Context, webhookID string) (*ApiResponse[DeletedObject], error) {
	return parseResponse[DeletedObject](w.http.delete(ctx, fmt.Sprintf("/webhooks/%s", webhookID), nil, ""))
}

func (w *WebhooksResource) Test(ctx context.Context, webhookID string) (*ApiResponse[WebhookTestResult], error) {
	return parseResponse[WebhookTestResult](w.http.post(ctx, fmt.Sprintf("/webhooks/%s/test", webhookID), map[string]any{}, ""))
}

func sign(payload string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}
