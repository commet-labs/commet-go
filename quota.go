package commet

import "context"

type GetAllQuotaAllowancesParams struct {
	CustomerID string `json:"customer_id"`
}

type RemoveQuotaParams struct {
	FeatureCode    string  `json:"feature_code"`
	Count          *int    `json:"count,omitempty"`
	CustomerID     *string `json:"customer_id,omitempty"`
	ExternalID     *string `json:"external_id,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type GetQuotaAllowanceParams struct {
	CustomerID  string `json:"customer_id"`
	FeatureCode string `json:"feature_code"`
}

type AddQuotaParams struct {
	FeatureCode    string  `json:"feature_code"`
	Count          *int    `json:"count,omitempty"`
	CustomerID     *string `json:"customer_id,omitempty"`
	ExternalID     *string `json:"external_id,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type SetQuotaParams struct {
	FeatureCode    string  `json:"feature_code"`
	Count          int     `json:"count"`
	CustomerID     *string `json:"customer_id,omitempty"`
	ExternalID     *string `json:"external_id,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type QuotaResource struct {
	http *httpClient
}

// Get all quota allowances for a customer across every quota feature in their plan.
func (r *QuotaResource) GetAll(ctx context.Context, params *GetAllQuotaAllowancesParams) (*QuotaGetAllResult, error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseDirectResponse[QuotaGetAllResult](r.http.get(ctx, "/usage/quota/all", query))
}

// Remove from a customer's quota allowance for a feature. Defaults to 1 if count is omitted. Returns 400 insufficient_balance if the balance would go negative.
func (r *QuotaResource) Remove(ctx context.Context, params *RemoveQuotaParams) (*UsageQuotaEvent, error) {
	body := buildBody(map[string]any{
		"feature_code": params.FeatureCode,
		"count":        params.Count,
		"customer_id":  params.CustomerID,
		"external_id":  params.ExternalID,
	})
	return parseDirectResponse[UsageQuotaEvent](r.http.post(ctx, "/usage/quota/remove", body, params.IdempotencyKey))
}

// Get the current quota allowance (used vs included) for a specific feature.
func (r *QuotaResource) Get(ctx context.Context, params *GetQuotaAllowanceParams) (*UsageQuota, error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	if params.FeatureCode != "" {
		query["feature_code"] = params.FeatureCode
	}
	return parseDirectResponse[UsageQuota](r.http.get(ctx, "/usage/quota", query))
}

// Add to a customer's quota allowance for a feature. Defaults to 1 if count is omitted.
func (r *QuotaResource) Add(ctx context.Context, params *AddQuotaParams) (*UsageQuotaEvent, error) {
	body := buildBody(map[string]any{
		"feature_code": params.FeatureCode,
		"count":        params.Count,
		"customer_id":  params.CustomerID,
		"external_id":  params.ExternalID,
	})
	return parseDirectResponse[UsageQuotaEvent](r.http.post(ctx, "/usage/quota", body, params.IdempotencyKey))
}

// Set a customer's quota allowance for a feature to an exact value.
func (r *QuotaResource) Set(ctx context.Context, params *SetQuotaParams) (*UsageQuotaEvent, error) {
	body := buildBody(map[string]any{
		"feature_code": params.FeatureCode,
		"count":        params.Count,
		"customer_id":  params.CustomerID,
		"external_id":  params.ExternalID,
	})
	return parseDirectResponse[UsageQuotaEvent](r.http.put(ctx, "/usage/quota", body, params.IdempotencyKey))
}
