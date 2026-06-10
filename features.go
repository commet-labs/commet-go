package commet

import (
	"context"
	"fmt"
)

type CreateFeatureParams struct {
	Name           string      `json:"name"`
	Code           string      `json:"code"`
	Type           FeatureType `json:"type"`
	Description    *string     `json:"description,omitempty"`
	UnitName       *string     `json:"unit_name,omitempty"`
	IdempotencyKey string      `json:"-"`
}

type UpdateFeatureParams struct {
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	UnitName       *string `json:"unit_name,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type FeaturesResource struct {
	http *httpClient
}

// List every feature defined in the organization. This is the organization's feature catalog (definitions), not a customer's feature access.
func (r *FeaturesResource) List(ctx context.Context) (*ApiResponse[[]Feature], error) {
	return parseResponse[[]Feature](r.http.get(ctx, "/features", nil))
}

// Get a single feature definition by code from the organization's feature catalog.
func (r *FeaturesResource) Get(ctx context.Context, code string) (*ApiResponse[Feature], error) {
	return parseResponse[Feature](r.http.get(ctx, fmt.Sprintf("/features/%s", code), nil))
}

// Create a new feature. Code must be lowercase alphanumeric with underscores.
func (r *FeaturesResource) Create(ctx context.Context, params *CreateFeatureParams) (*ApiResponse[Feature], error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"code":        params.Code,
		"type":        params.Type,
		"description": params.Description,
		"unit_name":   params.UnitName,
	})
	return parseResponse[Feature](r.http.post(ctx, "/features/manage", body, params.IdempotencyKey))
}

// Update a feature's name, description, or unit name. At least one field must be provided.
func (r *FeaturesResource) Update(ctx context.Context, code string, params *UpdateFeatureParams) (*ApiResponse[Feature], error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"unit_name":   params.UnitName,
	})
	return parseResponse[Feature](r.http.put(ctx, fmt.Sprintf("/features/%s/manage", code), body, params.IdempotencyKey))
}

// Delete a feature. Fails if the feature is attached to active plans or has an active add-on.
func (r *FeaturesResource) Delete(ctx context.Context, code string) (*ApiResponse[DeletedObject], error) {
	return parseResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/features/%s/manage", code), nil, ""))
}
