package commet

import "context"

// GetPortalURLParams holds parameters for requesting a portal URL.
type GetPortalURLParams struct {
	CustomerID     string `json:"customer_id,omitempty"`
	ExternalID     string `json:"external_id,omitempty"`
	Email          string `json:"email,omitempty"`
	IdempotencyKey string `json:"-"`
}

// PortalResource provides access to customer portal operations.
type PortalResource struct {
	http *httpClient
}

// GetURL requests a portal access URL for a customer.
func (r *PortalResource) GetURL(ctx context.Context, params *GetPortalURLParams) (*ApiResponse, error) {
	body := buildBody(map[string]any{
		"customer_id": params.CustomerID,
		"external_id": params.ExternalID,
		"email":       params.Email,
	})
	return r.http.post(ctx, "/portal/request-access", body, params.IdempotencyKey)
}
