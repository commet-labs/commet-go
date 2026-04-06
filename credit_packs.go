package commet

import "context"

// CreditPacksResource provides access to credit pack operations.
type CreditPacksResource struct {
	http *httpClient
}

// List retrieves all available credit packs.
func (r *CreditPacksResource) List(ctx context.Context) (*ApiResponse, error) {
	return r.http.get(ctx, "/credit-packs", nil)
}
