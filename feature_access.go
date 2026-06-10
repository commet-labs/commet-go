package commet

import (
	"context"
	"fmt"
)

type ListFeatureAccessParams struct {
	CustomerID string `json:"customer_id"`
}

type GetFeatureAccessParams struct {
	CustomerID string  `json:"customer_id"`
	Action     *string `json:"action,omitempty"`
}

type CanUseFeatureParams struct {
	CustomerID string `json:"customer_id"`
}

type FeatureAccessResource struct {
	http *httpClient
}

// List all features for a customer's active subscription, scoped by the customerId query parameter.
func (r *FeatureAccessResource) List(ctx context.Context, params *ListFeatureAccessParams) (*ApiResponse[[]FeatureAccess], error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseResponse[[]FeatureAccess](r.http.get(ctx, "/feature-access", query))
}

// Get feature access details for a customer. Use action=canUse to check if the customer can consume one more unit.
func (r *FeatureAccessResource) Get(ctx context.Context, code string, params *GetFeatureAccessParams) (*ApiResponse[FeatureLookup], error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	if params.Action != nil {
		query["action"] = *params.Action
	}
	return parseResponse[FeatureLookup](r.http.get(ctx, fmt.Sprintf("/feature-access/%s", code), query))
}

// Get feature access details for a customer. Use action=canUse to check if the customer can consume one more unit.
func (r *FeatureAccessResource) CanUse(ctx context.Context, code string, params *CanUseFeatureParams) (*ApiResponse[FeatureLookup], error) {
	query := map[string]string{}
	query["action"] = "canUse"
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseResponse[FeatureLookup](r.http.get(ctx, fmt.Sprintf("/feature-access/%s", code), query))
}
