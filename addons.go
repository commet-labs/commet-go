package commet

import "context"

type ActiveAddon struct {
	Slug             string `json:"slug"`
	Name             string `json:"name"`
	BasePrice        int    `json:"base_price"`
	FeatureCode      string `json:"feature_code"`
	FeatureName      string `json:"feature_name"`
	FeatureType      string `json:"feature_type"`
	ConsumptionModel string `json:"consumption_model,omitempty"`
	ActivatedAt      string `json:"activated_at"`
}

type AddonsResource struct {
	http *httpClient
}

func (r *AddonsResource) GetActive(ctx context.Context, customerID string) (*ApiResponse[[]ActiveAddon], error) {
	return parseResponse[[]ActiveAddon](r.http.get(ctx, "/addons/active", map[string]string{
		"customer_id": customerID,
	}))
}
