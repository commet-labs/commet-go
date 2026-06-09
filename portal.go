package commet

import "context"

type RequestPortalAccessParams struct {
	Email          *string `json:"email,omitempty"`
	CustomerID     *string `json:"customer_id,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type PortalResource struct {
	http *httpClient
}

// Generate a customer portal URL. Exactly one identifier (email or customerId) is required.
func (r *PortalResource) GetURL(ctx context.Context, params *RequestPortalAccessParams) (*ApiResponse[PortalAccess], error) {
	body := buildBody(map[string]any{
		"email":       params.Email,
		"customer_id": params.CustomerID,
	})
	return parseResponse[PortalAccess](r.http.post(ctx, "/portal/request-access", body, params.IdempotencyKey))
}
