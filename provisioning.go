package commet

import "context"

type ProvisioningResource struct {
	http *httpClient
}

// Issue a fresh claim link for an organization that was provisioned headlessly and has not been claimed yet. Any previously issued link stops working.
func (r *ProvisioningResource) CreateClaimLink(ctx context.Context) (*ClaimLink, error) {
	return parseDirectResponse[ClaimLink](r.http.post(ctx, "/claim-link", map[string]any{}, ""))
}
