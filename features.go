package commet

import (
	"context"
	"fmt"
)

type CreateFeatureParams struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
	UnitName    string `json:"unit_name,omitempty"`
}

type UpdateFeatureParams struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UnitName    string `json:"unit_name,omitempty"`
}

type FeaturesResource struct {
	http *httpClient
}

func (r *FeaturesResource) Get(ctx context.Context, code string, customerID string) (*ApiResponse[FeatureAccess], error) {
	return parseResponse[FeatureAccess](r.http.get(ctx, fmt.Sprintf("/features/%s", code), map[string]string{
		"customer_id": customerID,
	}))
}

func (r *FeaturesResource) CanUse(ctx context.Context, code string, customerID string) (*ApiResponse[CanUseResult], error) {
	return parseResponse[CanUseResult](r.http.get(ctx, fmt.Sprintf("/features/%s", code), map[string]string{
		"customer_id": customerID,
		"action":      "canUse",
	}))
}

func (r *FeaturesResource) List(ctx context.Context, customerID string) (*ApiResponse[[]FeatureAccess], error) {
	return parseResponse[[]FeatureAccess](r.http.get(ctx, "/features", map[string]string{
		"customer_id": customerID,
	}))
}

func (r *FeaturesResource) Create(ctx context.Context, params *CreateFeatureParams) (*ApiResponse[Feature], error) {
	body := buildBody(map[string]any{
		"code":        params.Code,
		"name":        params.Name,
		"type":        params.Type,
		"description": params.Description,
		"unit_name":   params.UnitName,
	})
	return parseResponse[Feature](r.http.post(ctx, "/features/manage", body, ""))
}

func (r *FeaturesResource) Update(ctx context.Context, code string, params *UpdateFeatureParams) (*ApiResponse[Feature], error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"unit_name":   params.UnitName,
	})
	return parseResponse[Feature](r.http.put(ctx, fmt.Sprintf("/features/%s/manage", code), body, ""))
}

func (r *FeaturesResource) Delete(ctx context.Context, code string) (*ApiResponse[DeleteResult], error) {
	return parseResponse[DeleteResult](r.http.delete(ctx, fmt.Sprintf("/features/%s/manage", code), nil, ""))
}
