package commet

import (
	"context"
	"fmt"
)

type UpdateWebhookEndpointParams struct {
	URL            *string  `json:"url,omitempty"`
	Events         []string `json:"events,omitempty"`
	Description    *string  `json:"description,omitempty"`
	IsActive       *bool    `json:"is_active,omitempty"`
	APIVersion     *string  `json:"api_version,omitempty"`
	IdempotencyKey string   `json:"-"`
}

type TestWebhookEndpointParams struct {
	IdempotencyKey string `json:"-"`
}

type ListWebhookEndpointsParams struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type CreateWebhookEndpointParams struct {
	URL            string   `json:"url"`
	Events         []string `json:"events"`
	Description    *string  `json:"description,omitempty"`
	APIVersion     *string  `json:"api_version,omitempty"`
	IdempotencyKey string   `json:"-"`
}

type GeneratedWebhooksResource struct {
	http *httpClient
}

// Retrieve a webhook endpoint by its public ID.
func (r *GeneratedWebhooksResource) Get(ctx context.Context, id string) (*Webhook, error) {
	return parseDirectResponse[Webhook](r.http.get(ctx, fmt.Sprintf("/webhooks/%s", id), nil))
}

// Update a webhook endpoint. Only the provided fields change.
func (r *GeneratedWebhooksResource) Update(ctx context.Context, id string, params *UpdateWebhookEndpointParams) (*Webhook, error) {
	body := buildBody(map[string]any{
		"url":         params.URL,
		"events":      params.Events,
		"description": params.Description,
		"is_active":   params.IsActive,
		"api_version": params.APIVersion,
	})
	return parseDirectResponse[Webhook](r.http.patch(ctx, fmt.Sprintf("/webhooks/%s", id), body, params.IdempotencyKey))
}

// Permanently delete a webhook endpoint.
func (r *GeneratedWebhooksResource) Delete(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/webhooks/%s", id), nil, ""))
}

// Send a test event to a webhook endpoint to verify connectivity.
func (r *GeneratedWebhooksResource) Test(ctx context.Context, id string, params *TestWebhookEndpointParams) (*WebhookTest, error) {
	return parseDirectResponse[WebhookTest](r.http.post(ctx, fmt.Sprintf("/webhooks/%s/test", id), map[string]any{}, params.IdempotencyKey))
}

// List webhook endpoints with cursor-based pagination.
func (r *GeneratedWebhooksResource) List(ctx context.Context, params *ListWebhookEndpointsParams) (*WebhooksListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	return parseDirectResponse[WebhooksListResult](r.http.get(ctx, "/webhooks", query))
}

// Create a new webhook endpoint. The response includes the signing secret which is only returned once.
func (r *GeneratedWebhooksResource) Create(ctx context.Context, params *CreateWebhookEndpointParams) (*CreatedWebhook, error) {
	body := buildBody(map[string]any{
		"url":         params.URL,
		"events":      params.Events,
		"description": params.Description,
		"api_version": params.APIVersion,
	})
	return parseDirectResponse[CreatedWebhook](r.http.post(ctx, "/webhooks", body, params.IdempotencyKey))
}
