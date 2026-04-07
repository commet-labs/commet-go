package commet

import (
	"context"
	"fmt"
)

// CreateSubscriptionParams holds parameters for creating a subscription.
type CreateSubscriptionParams struct {
	CustomerID      string         `json:"customer_id,omitempty"`
	PlanCode        string         `json:"plan_code,omitempty"`
	PlanID          string         `json:"plan_id,omitempty"`
	BillingInterval string         `json:"billing_interval,omitempty"`
	InitialSeats    map[string]int `json:"initial_seats,omitempty"`
	SkipTrial       *bool          `json:"skip_trial,omitempty"`
	Name            string         `json:"name,omitempty"`
	StartDate       string         `json:"start_date,omitempty"`
	SuccessURL      string         `json:"success_url,omitempty"`
	IdempotencyKey  string         `json:"-"`
}

// CancelSubscriptionParams holds parameters for cancelling a subscription.
type CancelSubscriptionParams struct {
	Reason         string `json:"reason,omitempty"`
	Immediate      *bool  `json:"immediate,omitempty"`
	IdempotencyKey string `json:"-"`
}

// SubscriptionsResource provides access to subscription operations.
type SubscriptionsResource struct {
	http *httpClient
}

// Create creates a new subscription.
func (r *SubscriptionsResource) Create(ctx context.Context, params *CreateSubscriptionParams) (*ApiResponse, error) {
	body := buildBody(map[string]any{
		"customer_id":      params.CustomerID,
		"plan_code":        params.PlanCode,
		"plan_id":          params.PlanID,
		"billing_interval": params.BillingInterval,
		"initial_seats":    params.InitialSeats,
		"skip_trial":       params.SkipTrial,
		"name":             params.Name,
		"start_date":       params.StartDate,
		"success_url":      params.SuccessURL,
	})
	return r.http.post(ctx, "/subscriptions", body, params.IdempotencyKey)
}

// Get retrieves the active subscription for a customer.
func (r *SubscriptionsResource) Get(ctx context.Context, customerID string) (*ApiResponse, error) {
	return r.http.get(ctx, "/subscriptions/active", map[string]string{"customer_id": customerID})
}

// Cancel cancels a subscription.
func (r *SubscriptionsResource) Cancel(ctx context.Context, subscriptionID string, params *CancelSubscriptionParams) (*ApiResponse, error) {
	body := buildBody(map[string]any{
		"reason":    params.Reason,
		"immediate": params.Immediate,
	})
	idempotencyKey := ""
	if params != nil {
		idempotencyKey = params.IdempotencyKey
	}
	return r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/cancel", subscriptionID), body, idempotencyKey)
}
