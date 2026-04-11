package commet

import "context"

type GetPortalURLParams struct {
	CustomerID     string `json:"customer_id,omitempty"`
	Email          string `json:"email,omitempty"`
	IdempotencyKey string `json:"-"`
}

type PortalResource struct {
	http *httpClient
}

func (r *PortalResource) GetURL(ctx context.Context, params *GetPortalURLParams) (*ApiResponse[PortalSession], error) {
	body := buildBody(map[string]any{
		"customer_id": params.CustomerID,
		"email":       params.Email,
	})
	return parseResponse[PortalSession](r.http.post(ctx, "/portal/request-access", body, params.IdempotencyKey))
}
