package commet

import "context"

type CreditPacksResource struct {
	http *httpClient
}

func (r *CreditPacksResource) List(ctx context.Context) (*ApiResponse[[]CreditPack], error) {
	return parseResponse[[]CreditPack](r.http.get(ctx, "/credit-packs", nil))
}
