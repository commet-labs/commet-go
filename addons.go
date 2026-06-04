package commet

import (
	"context"
	"fmt"
)

type ActiveAddon struct {
	Object           string                `json:"object"`
	Livemode         bool                  `json:"livemode"`
	Slug             string                `json:"slug"`
	Name             string                `json:"name"`
	BasePrice        int                   `json:"base_price"`
	FeatureCode      string                `json:"feature_code"`
	FeatureName      string                `json:"feature_name"`
	FeatureType      FeatureType           `json:"feature_type"`
	ConsumptionModel AddonConsumptionModel `json:"consumption_model,omitempty"`
	ActivatedAt      string                `json:"activated_at"`
}

type ListAddonsParams struct {
	Limit  *int   `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type CreateAddonParams struct {
	Name             string   `json:"name"`
	Description      string   `json:"description,omitempty"`
	BasePrice        int      `json:"base_price"`
	FeatureID        string   `json:"feature_id"`
	ConsumptionModel string   `json:"consumption_model"`
	IncludedUnits    *int     `json:"included_units,omitempty"`
	OverageRate      *float64 `json:"overage_rate,omitempty"`
	CreditCost       *int     `json:"credit_cost,omitempty"`
}

type UpdateAddonParams struct {
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	BasePrice     *int     `json:"base_price,omitempty"`
	IncludedUnits *int     `json:"included_units,omitempty"`
	OverageRate   *float64 `json:"overage_rate,omitempty"`
}

type AddonsResource struct {
	http *httpClient
}

func (r *AddonsResource) ListActive(ctx context.Context, customerID string) (*ApiResponse[[]ActiveAddon], error) {
	return parseResponse[[]ActiveAddon](r.http.get(ctx, "/addons/active", map[string]string{
		"customer_id": customerID,
	}))
}

func (r *AddonsResource) List(ctx context.Context, params *ListAddonsParams) (*ApiResponse[[]Addon], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]Addon](r.http.get(ctx, "/addons", queryParams))
}

func (r *AddonsResource) Get(ctx context.Context, addonID string) (*ApiResponse[Addon], error) {
	return parseResponse[Addon](r.http.get(ctx, fmt.Sprintf("/addons/%s", addonID), nil))
}

func (r *AddonsResource) Create(ctx context.Context, params *CreateAddonParams) (*ApiResponse[Addon], error) {
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
	return parseResponse[Addon](r.http.post(ctx, "/addons", body, ""))
}

func (r *AddonsResource) Update(ctx context.Context, addonID string, params *UpdateAddonParams) (*ApiResponse[Addon], error) {
	body := buildBody(map[string]any{
		"name":           params.Name,
		"description":    params.Description,
		"base_price":     params.BasePrice,
		"included_units": params.IncludedUnits,
		"overage_rate":   params.OverageRate,
	})
	return parseResponse[Addon](r.http.put(ctx, fmt.Sprintf("/addons/%s", addonID), body, ""))
}

func (r *AddonsResource) Delete(ctx context.Context, addonID string) (*ApiResponse[DeleteResult], error) {
	return parseResponse[DeleteResult](r.http.delete(ctx, fmt.Sprintf("/addons/%s", addonID), nil, ""))
}
