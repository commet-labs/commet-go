package commet

import (
	"context"
	"fmt"
)

type ListApiKeysParams struct {
	Limit  *int   `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type CreateApiKeyParams struct {
	Name          string `json:"name"`
	ExpiresInDays *int   `json:"expires_in_days,omitempty"`
}

type ApiKeysResource struct {
	http *httpClient
}

func (r *ApiKeysResource) List(ctx context.Context, params *ListApiKeysParams) (*ApiResponse[[]ApiKey], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]ApiKey](r.http.get(ctx, "/api-keys", queryParams))
}

func (r *ApiKeysResource) Create(ctx context.Context, params *CreateApiKeyParams) (*ApiResponse[ApiKeyCreated], error) {
	body := buildBody(map[string]any{
		"name":            params.Name,
		"expires_in_days": params.ExpiresInDays,
	})
	return parseResponse[ApiKeyCreated](r.http.post(ctx, "/api-keys", body, ""))
}

func (r *ApiKeysResource) Delete(ctx context.Context, apiKeyID string) (*ApiResponse[struct{}], error) {
	return parseResponse[struct{}](r.http.delete(ctx, fmt.Sprintf("/api-keys/%s", apiKeyID), nil, ""))
}
