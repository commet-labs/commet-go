package commet

import "context"

type QuotaParams struct {
	FeatureCode    string `json:"feature_code"`
	Count          int    `json:"count"`
	CustomerID     string `json:"customer_id"`
	IdempotencyKey string `json:"-"`
}

type GetQuotaParams struct {
	FeatureCode string `json:"feature_code"`
	CustomerID  string `json:"customer_id"`
}

type GetAllQuotaParams struct {
	CustomerID string `json:"customer_id"`
}

type QuotaResource struct {
	http *httpClient
}

func (r *QuotaResource) Add(ctx context.Context, params *QuotaParams) (*ApiResponse[QuotaEvent], error) {
	count := params.Count
	if count == 0 {
		count = 1
	}
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"count":        count,
	})
	return parseResponse[QuotaEvent](r.http.post(ctx, "/usage/quota", body, params.IdempotencyKey))
}

func (r *QuotaResource) Set(ctx context.Context, params *QuotaParams) (*ApiResponse[QuotaEvent], error) {
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"count":        params.Count,
	})
	return parseResponse[QuotaEvent](r.http.put(ctx, "/usage/quota", body, params.IdempotencyKey))
}

func (r *QuotaResource) Remove(ctx context.Context, params *QuotaParams) (*ApiResponse[QuotaEvent], error) {
	count := params.Count
	if count == 0 {
		count = 1
	}
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"count":        count,
	})
	return parseResponse[QuotaEvent](r.http.delete(ctx, "/usage/quota", body, params.IdempotencyKey))
}

func (r *QuotaResource) Get(ctx context.Context, params *GetQuotaParams) (*ApiResponse[QuotaAllowance], error) {
	queryParams := map[string]string{
		"feature_code": params.FeatureCode,
	}
	if params.CustomerID != "" {
		queryParams["customer_id"] = params.CustomerID
	}
	return parseResponse[QuotaAllowance](r.http.get(ctx, "/usage/quota", queryParams))
}

func (r *QuotaResource) GetAll(ctx context.Context, params *GetAllQuotaParams) (*ApiResponse[[]QuotaAllowance], error) {
	queryParams := map[string]string{}
	if params != nil && params.CustomerID != "" {
		queryParams["customer_id"] = params.CustomerID
	}
	return parseResponse[[]QuotaAllowance](r.http.get(ctx, "/usage/quota/all", queryParams))
}
