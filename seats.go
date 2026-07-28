package commet

import "context"

type GetSeatBalanceParams struct {
	CustomerID  string `json:"customer_id"`
	FeatureCode string `json:"feature_code"`
}

type GetAllSeatBalancesParams struct {
	CustomerID string `json:"customer_id"`
}

type BulkSetSeatsParams struct {
	CustomerID     string         `json:"customer_id"`
	Seats          map[string]int `json:"seats"`
	IdempotencyKey string         `json:"-"`
}

type RemoveSeatsParams struct {
	CustomerID     string `json:"customer_id"`
	FeatureCode    string `json:"feature_code"`
	Count          int    `json:"count"`
	IdempotencyKey string `json:"-"`
}

type AddSeatsParams struct {
	CustomerID     string `json:"customer_id"`
	FeatureCode    string `json:"feature_code"`
	Count          int    `json:"count"`
	IdempotencyKey string `json:"-"`
}

type SetSeatsParams struct {
	CustomerID     string `json:"customer_id"`
	FeatureCode    string `json:"feature_code"`
	Count          int    `json:"count"`
	IdempotencyKey string `json:"-"`
}

type SeatsResource struct {
	http *httpClient
}

// Get current balance for a specific seat type.
func (r *SeatsResource) GetBalance(ctx context.Context, params *GetSeatBalanceParams) (*SeatBalance, error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	if params.FeatureCode != "" {
		query["feature_code"] = params.FeatureCode
	}
	return parseDirectResponse[SeatBalance](r.http.get(ctx, "/seats/balance", query))
}

// Get the current balance for all seat types in a customer's subscription.
func (r *SeatsResource) GetAllBalances(ctx context.Context, params *GetAllSeatBalancesParams) (*SeatBalanceCollection, error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseDirectResponse[SeatBalanceCollection](r.http.get(ctx, "/seats/balances", query))
}

// Set all seat types at once.
func (r *SeatsResource) SetAll(ctx context.Context, params *BulkSetSeatsParams) (*SeatsSetAllResult, error) {
	body := buildBody(map[string]any{
		"customer_id": params.CustomerID,
		"seats":       params.Seats,
	})
	return parseDirectResponse[SeatsSetAllResult](r.http.put(ctx, "/seats/bulk", body, params.IdempotencyKey))
}

// Remove seats from a customer's subscription. Takes effect at the end of the billing period.
func (r *SeatsResource) Remove(ctx context.Context, params *RemoveSeatsParams) (*SeatEvent, error) {
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"count":        params.Count,
	})
	return parseDirectResponse[SeatEvent](r.http.post(ctx, "/seats/remove", body, params.IdempotencyKey))
}

// Add seats to a customer's subscription. Prorates charges for the current billing period.
func (r *SeatsResource) Add(ctx context.Context, params *AddSeatsParams) (*SeatEvent, error) {
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"count":        params.Count,
	})
	return parseDirectResponse[SeatEvent](r.http.post(ctx, "/seats", body, params.IdempotencyKey))
}

// Set seats to an exact count.
func (r *SeatsResource) Set(ctx context.Context, params *SetSeatsParams) (*SeatEvent, error) {
	body := buildBody(map[string]any{
		"customer_id":  params.CustomerID,
		"feature_code": params.FeatureCode,
		"count":        params.Count,
	})
	return parseDirectResponse[SeatEvent](r.http.put(ctx, "/seats", body, params.IdempotencyKey))
}
