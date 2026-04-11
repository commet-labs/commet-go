package commet

import (
	"context"
	"fmt"
)

type FeaturesResource struct {
	http *httpClient
}

func (r *FeaturesResource) Get(ctx context.Context, code string, customerID string) (*ApiResponse[FeatureAccess], error) {
	return parseResponse[FeatureAccess](r.http.get(ctx, fmt.Sprintf("/features/%s", code), map[string]string{
		"customer_id": customerID,
	}))
}

func (r *FeaturesResource) Check(ctx context.Context, code string, customerID string) (*ApiResponse[CheckResult], error) {
	raw, err := r.http.get(ctx, fmt.Sprintf("/features/%s", code), map[string]string{
		"customer_id": customerID,
	})
	if err != nil {
		return nil, err
	}

	featureResp, err := parseResponse[FeatureAccess](raw, nil)
	if err != nil {
		return nil, err
	}

	if !featureResp.Success {
		return &ApiResponse[CheckResult]{
			Success: false,
			Code:    featureResp.Code,
			Message: featureResp.Message,
		}, nil
	}

	return &ApiResponse[CheckResult]{
		Success: true,
		Data:    CheckResult{Allowed: featureResp.Data.Allowed},
		Message: featureResp.Message,
	}, nil
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
