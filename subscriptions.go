package commet

import (
	"context"
	"fmt"
)

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

type CancelSubscriptionParams struct {
	Reason         string `json:"reason,omitempty"`
	Immediate      *bool  `json:"immediate,omitempty"`
	IdempotencyKey string `json:"-"`
}

type SubscriptionsResource struct {
	http *httpClient
}

func (r *SubscriptionsResource) Create(ctx context.Context, params *CreateSubscriptionParams) (*ApiResponse[Subscription], error) {
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
	return parseResponse[Subscription](r.http.post(ctx, "/subscriptions", body, params.IdempotencyKey))
}

func (r *SubscriptionsResource) Get(ctx context.Context, customerID string) (*ApiResponse[ActiveSubscription], error) {
	return parseResponse[ActiveSubscription](r.http.get(ctx, "/subscriptions/active", map[string]string{"customer_id": customerID}))
}

func (r *SubscriptionsResource) Cancel(ctx context.Context, subscriptionID string, params *CancelSubscriptionParams) (*ApiResponse[Subscription], error) {
	body := buildBody(map[string]any{
		"reason":    params.Reason,
		"immediate": params.Immediate,
	})
	return parseResponse[Subscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/cancel", subscriptionID), body, params.IdempotencyKey))
}
