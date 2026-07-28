package commet

import (
	"context"
	"fmt"
)

type ListActiveAddonsParams struct {
	CustomerID string `json:"customer_id"`
}

type UpdateAddonParams struct {
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	BasePrice      *int    `json:"base_price,omitempty"`
	IncludedUnits  *int    `json:"included_units,omitempty"`
	OverageRate    *int    `json:"overage_rate,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type ListAddonsParams struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type CreateAddonParams struct {
	Name             string  `json:"name"`
	Description      *string `json:"description,omitempty"`
	BasePrice        int     `json:"base_price"`
	FeatureID        string  `json:"feature_id"`
	ConsumptionModel string  `json:"consumption_model"`
	IncludedUnits    *int    `json:"included_units,omitempty"`
	OverageRate      *int    `json:"overage_rate,omitempty"`
	CreditCost       *int    `json:"credit_cost,omitempty"`
	IdempotencyKey   string  `json:"-"`
}

type AddonsResource struct {
	http *httpClient
}

// List all active add-ons for a customer's subscription.
func (r *AddonsResource) ListActive(ctx context.Context, params *ListActiveAddonsParams) (*AddonsListActiveResult, error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseDirectResponse[AddonsListActiveResult](r.http.get(ctx, "/active-addons", query))
}

// Retrieve an add-on by its public ID or slug.
func (r *AddonsResource) Get(ctx context.Context, id string) (*Addon, error) {
	return parseDirectResponse[Addon](r.http.get(ctx, fmt.Sprintf("/addons/%s", id), nil))
}

// Update an add-on's name, description, or pricing.
func (r *AddonsResource) Update(ctx context.Context, id string, params *UpdateAddonParams) (*Addon, error) {
	body := buildBody(map[string]any{
		"name":           params.Name,
		"description":    params.Description,
		"base_price":     params.BasePrice,
		"included_units": params.IncludedUnits,
		"overage_rate":   params.OverageRate,
	})
	return parseDirectResponse[Addon](r.http.patch(ctx, fmt.Sprintf("/addons/%s", id), body, params.IdempotencyKey))
}

// Soft-delete an add-on. Fails if the add-on has active subscriptions.
func (r *AddonsResource) Delete(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/addons/%s", id), nil, ""))
}

// List all add-ons with cursor-based pagination.
func (r *AddonsResource) List(ctx context.Context, params *ListAddonsParams) (*AddonsListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	return parseDirectResponse[AddonsListResult](r.http.get(ctx, "/addons", query))
}

// Create a new add-on linked to a feature. Each feature can only be assigned to one add-on.
func (r *AddonsResource) Create(ctx context.Context, params *CreateAddonParams) (*Addon, error) {
	body := buildBody(map[string]any{
		"name":              params.Name,
		"description":       params.Description,
		"base_price":        params.BasePrice,
		"feature_id":        params.FeatureID,
		"consumption_model": params.ConsumptionModel,
		"included_units":    params.IncludedUnits,
		"overage_rate":      params.OverageRate,
		"credit_cost":       params.CreditCost,
	})
	return parseDirectResponse[Addon](r.http.post(ctx, "/addons", body, params.IdempotencyKey))
}
