package commet

import (
	"context"
	"fmt"
)

// FeaturesResource provides access to feature operations.
type FeaturesResource struct {
	http *httpClient
}

// Get retrieves a feature by code for a customer.
func (r *FeaturesResource) Get(ctx context.Context, code string, customerID string) (*ApiResponse, error) {
	return r.http.get(ctx, fmt.Sprintf("/features/%s", code), map[string]string{
		"customer_id": customerID,
	})
}

// Check checks if a customer has access to a feature, returning an ApiResponse with data containing "allowed" bool.
func (r *FeaturesResource) Check(ctx context.Context, code string, customerID string) (*ApiResponse, error) {
	result, err := r.http.get(ctx, fmt.Sprintf("/features/%s", code), map[string]string{
		"customer_id": customerID,
	})
	if err != nil {
		return nil, err
	}

	if !result.Success || result.Data == nil {
		return &ApiResponse{
			Success: false,
			Data:    map[string]any{"allowed": false},
			Message: result.Message,
		}, nil
	}

	dataMap, ok := result.Data.(map[string]any)
	if !ok {
		return &ApiResponse{
			Success: false,
			Data:    map[string]any{"allowed": false},
			Message: result.Message,
		}, nil
	}

	allowed, _ := dataMap["allowed"].(bool)
	return &ApiResponse{
		Success: true,
		Data:    map[string]any{"allowed": allowed},
		Message: result.Message,
	}, nil
}

// CanUse checks if a customer can use a feature.
func (r *FeaturesResource) CanUse(ctx context.Context, code string, customerID string) (*ApiResponse, error) {
	return r.http.get(ctx, fmt.Sprintf("/features/%s", code), map[string]string{
		"customer_id": customerID,
		"action":      "canUse",
	})
}

// List retrieves all features for a customer.
func (r *FeaturesResource) List(ctx context.Context, customerID string) (*ApiResponse, error) {
	return r.http.get(ctx, "/features", map[string]string{
		"customer_id": customerID,
	})
}
