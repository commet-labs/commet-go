package commet

import (
	"context"
	"fmt"
)

type GetFeatureAccessParams struct {
	CustomerID string `json:"customer_id"`
}

type ListFeatureAccessParams struct {
	CustomerID string `json:"customer_id"`
}

type FeatureAccessResource struct {
	http *httpClient
}

// Get one feature's access and current usage for a customer. To evaluate a prospective consumption, use POST /usage/check.
func (r *FeatureAccessResource) Get(ctx context.Context, code string, params *GetFeatureAccessParams) (*FeatureAccess, error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseDirectResponse[FeatureAccess](r.http.get(ctx, fmt.Sprintf("/feature-access/%s", code), query))
}

// List a customer's feature access and current usage.
func (r *FeatureAccessResource) List(ctx context.Context, params *ListFeatureAccessParams) (*FeatureAccessListResult, error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseDirectResponse[FeatureAccessListResult](r.http.get(ctx, "/feature-access", query))
}
