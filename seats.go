package commet

import "context"

type SeatParams struct {
	FeatureCode    string `json:"feature_code"`
	Count          int    `json:"count"`
	CustomerID     string `json:"customer_id"`
	IdempotencyKey string `json:"-"`
}

type SetAllSeatsParams struct {
	Seats          map[string]int `json:"seats"`
	CustomerID     string         `json:"customer_id"`
	IdempotencyKey string         `json:"-"`
}

type GetSeatBalanceParams struct {
	FeatureCode string `json:"feature_code"`
	CustomerID  string `json:"customer_id"`
}

type GetAllSeatBalancesParams struct {
	CustomerID string `json:"customer_id"`
}

type SeatsResource struct {
	http *httpClient
}

func (r *SeatsResource) Add(ctx context.Context, params *SeatParams) (*ApiResponse[SeatEvent], error) {
	count := params.Count
	if count == 0 {
		count = 1
	}
	body := buildBody(map[string]any{
		"feature_code": params.FeatureCode,
		"count":        count,
		"customer_id":  params.CustomerID,
	})
	return parseResponse[SeatEvent](r.http.post(ctx, "/seats", body, params.IdempotencyKey))
}

func (r *SeatsResource) Remove(ctx context.Context, params *SeatParams) (*ApiResponse[SeatEvent], error) {
	count := params.Count
	if count == 0 {
		count = 1
	}
	body := buildBody(map[string]any{
		"feature_code": params.FeatureCode,
		"count":        count,
		"customer_id":  params.CustomerID,
	})
	return parseResponse[SeatEvent](r.http.delete(ctx, "/seats", body, params.IdempotencyKey))
}

func (r *SeatsResource) Set(ctx context.Context, params *SeatParams) (*ApiResponse[SeatEvent], error) {
	body := buildBody(map[string]any{
		"feature_code": params.FeatureCode,
		"count":        params.Count,
		"customer_id":  params.CustomerID,
	})
	return parseResponse[SeatEvent](r.http.put(ctx, "/seats", body, params.IdempotencyKey))
}

func (r *SeatsResource) SetAll(ctx context.Context, params *SetAllSeatsParams) (*ApiResponse[[]SeatEvent], error) {
	body := buildBody(map[string]any{
		"seats":       params.Seats,
		"customer_id": params.CustomerID,
	})
	return parseResponse[[]SeatEvent](r.http.put(ctx, "/seats/bulk", body, params.IdempotencyKey))
}

func (r *SeatsResource) GetBalance(ctx context.Context, params *GetSeatBalanceParams) (*ApiResponse[SeatBalance], error) {
	queryParams := map[string]string{
		"feature_code": params.FeatureCode,
	}
	if params.CustomerID != "" {
		queryParams["customer_id"] = params.CustomerID
	}
	return parseResponse[SeatBalance](r.http.get(ctx, "/seats/balance", queryParams))
}

func (r *SeatsResource) GetAllBalances(ctx context.Context, params *GetAllSeatBalancesParams) (*ApiResponse[map[string]SeatBalance], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.CustomerID != "" {
			queryParams["customer_id"] = params.CustomerID
		}
	}
	return parseResponse[map[string]SeatBalance](r.http.get(ctx, "/seats/balances", queryParams))
}
