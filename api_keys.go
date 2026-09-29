package commet

import (
	"context"
	"encoding/json"
	"fmt"
)

type ListApiKeysParams struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type CreateApiKeyParams struct {
	Name           string                         `json:"name"`
	ExpiresInDays  *int                           `json:"expires_in_days,omitempty"`
	Permissions    *CreateApiKeyParamsPermissions `json:"permissions,omitempty"`
	IdempotencyKey string                         `json:"-"`
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

// Create a full-access or restricted API key. Provide permissions to restrict access; the full key is returned only once. A restricted key with api_key: write may only create restricted keys with the same or fewer permissions, and they expire no later than the key that creates them.
func (r *ApiKeysResource) Create(ctx context.Context, params *CreateApiKeyParams) (*CreatedApiKey, error) {
	body := buildBody(map[string]any{
		"name":            params.Name,
		"expires_in_days": params.ExpiresInDays,
	})
	if params.Permissions != nil {
		encoded, err := json.Marshal(params.Permissions)
		if err != nil {
			return nil, err
		}
		body["permissions"] = json.RawMessage(encoded)
	}
	return parseDirectResponse[CreatedApiKey](r.http.post(ctx, "/api-keys", body, params.IdempotencyKey))
}
