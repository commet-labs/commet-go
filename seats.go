package commet

import "context"

// SeatParams holds parameters for seat add/remove/set operations.
type SeatParams struct {
	SeatType       string `json:"seat_type"`
	Count          int    `json:"count"`
	CustomerID     string `json:"customer_id,omitempty"`
	ExternalID     string `json:"external_id,omitempty"`
	IdempotencyKey string `json:"-"`
}

// SetAllSeatsParams holds parameters for setting all seat types at once.
type SetAllSeatsParams struct {
	Seats          map[string]int `json:"seats"`
	CustomerID     string         `json:"customer_id,omitempty"`
	ExternalID     string         `json:"external_id,omitempty"`
	IdempotencyKey string         `json:"-"`
}

// GetSeatBalanceParams holds parameters for retrieving a seat balance.
type GetSeatBalanceParams struct {
	SeatType   string `json:"seat_type"`
	CustomerID string `json:"customer_id,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

// GetAllSeatBalancesParams holds parameters for retrieving all seat balances.
type GetAllSeatBalancesParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

// SeatsResource provides access to seat operations.
type SeatsResource struct {
	http *httpClient
}

// Add adds seats of a given type.
func (r *SeatsResource) Add(ctx context.Context, params *SeatParams) (*ApiResponse, error) {
	body := buildBody(map[string]any{
		"seat_type":   params.SeatType,
		"count":       params.Count,
		"customer_id": params.CustomerID,
		"external_id": params.ExternalID,
	})
	return r.http.post(ctx, "/seats", body, params.IdempotencyKey)
}

// Remove removes seats of a given type.
func (r *SeatsResource) Remove(ctx context.Context, params *SeatParams) (*ApiResponse, error) {
	body := buildBody(map[string]any{
		"seat_type":   params.SeatType,
		"count":       params.Count,
		"customer_id": params.CustomerID,
		"external_id": params.ExternalID,
	})
	return r.http.delete(ctx, "/seats", body, params.IdempotencyKey)
}

// Set sets the seat count for a given type.
func (r *SeatsResource) Set(ctx context.Context, params *SeatParams) (*ApiResponse, error) {
	body := buildBody(map[string]any{
		"seat_type":   params.SeatType,
		"count":       params.Count,
		"customer_id": params.CustomerID,
		"external_id": params.ExternalID,
	})
	return r.http.put(ctx, "/seats", body, params.IdempotencyKey)
}

// SetAll sets multiple seat types at once.
func (r *SeatsResource) SetAll(ctx context.Context, params *SetAllSeatsParams) (*ApiResponse, error) {
	body := buildBody(map[string]any{
		"seats":       params.Seats,
		"customer_id": params.CustomerID,
		"external_id": params.ExternalID,
	})
	return r.http.put(ctx, "/seats/bulk", body, params.IdempotencyKey)
}

// GetBalance retrieves the balance for a specific seat type.
func (r *SeatsResource) GetBalance(ctx context.Context, params *GetSeatBalanceParams) (*ApiResponse, error) {
	queryParams := map[string]string{
		"seat_type": params.SeatType,
	}
	if params.CustomerID != "" {
		queryParams["customer_id"] = params.CustomerID
	}
	if params.ExternalID != "" {
		queryParams["external_id"] = params.ExternalID
	}
	return r.http.get(ctx, "/seats/balance", queryParams)
}

// GetAllBalances retrieves balances for all seat types.
func (r *SeatsResource) GetAllBalances(ctx context.Context, params *GetAllSeatBalancesParams) (*ApiResponse, error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.CustomerID != "" {
			queryParams["customer_id"] = params.CustomerID
		}
		if params.ExternalID != "" {
			queryParams["external_id"] = params.ExternalID
		}
	}
	return r.http.get(ctx, "/seats/balances", queryParams)
}
