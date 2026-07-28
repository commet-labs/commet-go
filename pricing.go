package commet

import (
	"context"
	"fmt"
)

type UpdateMarketGroupParams struct {
	Name           string         `json:"name"`
	CountryCodes   []string       `json:"country_codes"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	IdempotencyKey string         `json:"-"`
}

type CreateMarketGroupParams struct {
	Name           string         `json:"name"`
	CountryCodes   []string       `json:"country_codes"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	IdempotencyKey string         `json:"-"`
}

type PricingResource struct {
	http *httpClient
}

// Get one reusable pricing market group.
func (r *PricingResource) GetMarketGroup(ctx context.Context, id string) (*MarketGroup, error) {
	return parseDirectResponse[MarketGroup](r.http.get(ctx, fmt.Sprintf("/pricing/market-groups/%s", id), nil))
}

// Replace the name, countries, and metadata of a pricing market group.
func (r *PricingResource) UpdateMarketGroup(ctx context.Context, id string, params *UpdateMarketGroupParams) (*MarketGroup, error) {
	body := buildBody(map[string]any{
		"name":          params.Name,
		"country_codes": params.CountryCodes,
		"metadata":      params.Metadata,
	})
	return parseDirectResponse[MarketGroup](r.http.patch(ctx, fmt.Sprintf("/pricing/market-groups/%s", id), body, params.IdempotencyKey))
}

// Delete an unused pricing market group. Groups referenced by prices or subscriptions cannot be deleted.
func (r *PricingResource) DeleteMarketGroup(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/pricing/market-groups/%s", id), nil, ""))
}

// List reusable country groups used to resolve market-specific prices independently from currency.
func (r *PricingResource) ListMarketGroups(ctx context.Context) (*PricingListMarketGroupsResult, error) {
	return parseDirectResponse[PricingListMarketGroupsResult](r.http.get(ctx, "/pricing/market-groups", nil))
}

// Create a reusable country group. Countries can belong to only one active group; each price chooses its currency independently.
func (r *PricingResource) CreateMarketGroup(ctx context.Context, params *CreateMarketGroupParams) (*MarketGroup, error) {
	body := buildBody(map[string]any{
		"name":          params.Name,
		"country_codes": params.CountryCodes,
		"metadata":      params.Metadata,
	})
	return parseDirectResponse[MarketGroup](r.http.post(ctx, "/pricing/market-groups", body, params.IdempotencyKey))
}
