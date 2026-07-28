package commet

import (
	"context"
	"fmt"
)

type UpdateFeatureParams struct {
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	UnitName       *string `json:"unit_name,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type CreateFeatureParams struct {
	Name           string  `json:"name"`
	Code           string  `json:"code"`
	Type           string  `json:"type"`
	Description    *string `json:"description,omitempty"`
	UnitName       *string `json:"unit_name,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type FeaturesResource struct {
	http *httpClient
}

// Get a single feature definition by code from the organization's feature catalog.
func (r *FeaturesResource) Get(ctx context.Context, code string) (*Feature, error) {
	return parseDirectResponse[Feature](r.http.get(ctx, fmt.Sprintf("/features/%s", code), nil))
}

// Update a feature's name, description, or unit name. At least one field must be provided.
func (r *FeaturesResource) Update(ctx context.Context, code string, params *UpdateFeatureParams) (*Feature, error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"unit_name":   params.UnitName,
	})
	return parseDirectResponse[Feature](r.http.patch(ctx, fmt.Sprintf("/features/%s", code), body, params.IdempotencyKey))
}

// Delete a feature. Fails if the feature is attached to active plans or has an active add-on.
func (r *FeaturesResource) Delete(ctx context.Context, code string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/features/%s", code), nil, ""))
}

// List every feature defined in the organization. This is the organization's feature catalog (definitions), not a customer's feature access.
func (r *FeaturesResource) List(ctx context.Context) (*FeaturesListResult, error) {
	return parseDirectResponse[FeaturesListResult](r.http.get(ctx, "/features", nil))
}

// Create a new feature. Code must be lowercase alphanumeric with underscores.
func (r *FeaturesResource) Create(ctx context.Context, params *CreateFeatureParams) (*Feature, error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"code":        params.Code,
		"type":        params.Type,
		"description": params.Description,
		"unit_name":   params.UnitName,
	})
	return parseDirectResponse[Feature](r.http.post(ctx, "/features", body, params.IdempotencyKey))
}
