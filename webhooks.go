package commet

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type ListWebhooksParams struct {
	Limit  *int   `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type CreateWebhookParams struct {
	URL         string   `json:"url"`
	Events      []string `json:"events"`
	Description string   `json:"description,omitempty"`
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

func (w *WebhooksResource) VerifyAndParse(rawBody string, signature string, secret string) (map[string]any, error) {
	if !w.Verify(rawBody, signature, secret) {
		return nil, &CommetError{
			Message: "Invalid webhook signature",
			Code:    "INVALID_SIGNATURE",
		}
	}

	var result map[string]any
	if err := json.Unmarshal([]byte(rawBody), &result); err != nil {
		return nil, &CommetError{
			Message: "Failed to parse webhook payload",
			Code:    "INVALID_JSON",
		}
	}

	return result, nil
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
	})
	return parseResponse[WebhookEndpointCreated](w.http.post(ctx, "/webhooks", body, ""))
}

func (w *WebhooksResource) Delete(ctx context.Context, webhookID string) (*ApiResponse[DeleteResult], error) {
	return parseResponse[DeleteResult](w.http.delete(ctx, fmt.Sprintf("/webhooks/%s", webhookID), nil, ""))
}

func (w *WebhooksResource) Test(ctx context.Context, webhookID string) (*ApiResponse[WebhookTestResult], error) {
	return parseResponse[WebhookTestResult](w.http.post(ctx, fmt.Sprintf("/webhooks/%s/test", webhookID), map[string]any{}, ""))
}

func sign(payload string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}
