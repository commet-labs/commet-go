package commet

import (
	"context"
	"fmt"
)

type ListApiKeysParams struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type CreateApiKeyParams struct {
	Name           string `json:"name"`
	ExpiresInDays  *int   `json:"expires_in_days,omitempty"`
	IdempotencyKey string `json:"-"`
}

type ApiKeysResource struct {
	http *httpClient
}

// Permanently revoke and delete an API key.
func (r *ApiKeysResource) Delete(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/api-keys/%s", id), nil, ""))
}

// List API keys with cursor-based pagination. Keys are returned without the full secret.
func (r *ApiKeysResource) List(ctx context.Context, params *ListApiKeysParams) (*ApiKeysListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	return parseDirectResponse[ApiKeysListResult](r.http.get(ctx, "/api-keys", query))
}

// Create a new API key. The full key is only returned once in the response.
func (r *ApiKeysResource) Create(ctx context.Context, params *CreateApiKeyParams) (*CreatedApiKey, error) {
	body := buildBody(map[string]any{
		"name":            params.Name,
		"expires_in_days": params.ExpiresInDays,
	})
	return parseDirectResponse[CreatedApiKey](r.http.post(ctx, "/api-keys", body, params.IdempotencyKey))
}
