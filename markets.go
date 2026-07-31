package commet

import (
	"context"
	"fmt"
)

type UpdateMarketParams struct {
	Name           string         `json:"name"`
	CountryCodes   []string       `json:"country_codes"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	IdempotencyKey string         `json:"-"`
}

type CreateMarketParams struct {
	Name           string         `json:"name"`
	CountryCodes   []string       `json:"country_codes"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	IdempotencyKey string         `json:"-"`
}

type MarketsResource struct {
	http *httpClient
}

// Get one reusable market.
func (r *MarketsResource) Get(ctx context.Context, id string) (*Market, error) {
	return parseDirectResponse[Market](r.http.get(ctx, fmt.Sprintf("/markets/%s", id), nil))
}

// Replace the name, countries, and metadata of a market.
func (r *MarketsResource) Update(ctx context.Context, id string, params *UpdateMarketParams) (*Market, error) {
	body := buildBody(map[string]any{
		"name":          params.Name,
		"country_codes": params.CountryCodes,
		"metadata":      params.Metadata,
	})
	return parseDirectResponse[Market](r.http.patch(ctx, fmt.Sprintf("/markets/%s", id), body, params.IdempotencyKey))
}

// Delete an unused market. Markets referenced by prices or subscriptions cannot be deleted.
func (r *MarketsResource) Delete(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/markets/%s", id), nil, ""))
}

// List reusable country groups that resolve market-specific prices independently from currency.
func (r *MarketsResource) List(ctx context.Context) (*MarketsListResult, error) {
	return parseDirectResponse[MarketsListResult](r.http.get(ctx, "/markets", nil))
}

// Create a reusable market without attaching it to a plan or price. Countries can belong to only one active market.
func (r *MarketsResource) Create(ctx context.Context, params *CreateMarketParams) (*Market, error) {
	body := buildBody(map[string]any{
		"name":          params.Name,
		"country_codes": params.CountryCodes,
		"metadata":      params.Metadata,
	})
	return parseDirectResponse[Market](r.http.post(ctx, "/markets", body, params.IdempotencyKey))
}
