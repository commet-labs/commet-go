package commet

import (
	"context"
	"fmt"
)

type ActivateAddonParams struct {
	AddonID        string `json:"addon_id"`
	IdempotencyKey string `json:"-"`
}

type AdjustBalanceParams struct {
	Amount         int     `json:"amount"`
	Reason         *string `json:"reason,omitempty"`
	Type           *string `json:"type,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type TopupBalanceParams struct {
	Amount         int    `json:"amount"`
	IdempotencyKey string `json:"-"`
}

type CancelSubscriptionParams struct {
	Reason         *string `json:"reason,omitempty"`
	Immediate      *bool   `json:"immediate,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type ChangePlanParams struct {
	NewPlanID          *string `json:"new_plan_id,omitempty"`
	NewBillingInterval *string `json:"new_billing_interval,omitempty"`
	SuccessURL         *string `json:"success_url,omitempty"`
	OfferID            *string `json:"offer_id,omitempty"`
	IdempotencyKey     string  `json:"-"`
}

type PurchaseCreditsParams struct {
	CreditPackID   string `json:"credit_pack_id"`
	IdempotencyKey string `json:"-"`
}

type ApplySubscriptionOfferParams struct {
	OfferID        string  `json:"offer_id"`
	ExpiresAt      *string `json:"expires_at,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type UpdatePaymentMethodParams struct {
	SuccessURL     *string `json:"success_url,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type PreviewChangePlanParams struct {
	PlanID          string  `json:"plan_id"`
	BillingInterval *string `json:"billing_interval,omitempty"`
	OfferID         *string `json:"offer_id,omitempty"`
	IdempotencyKey  string  `json:"-"`
}

type ReactivateSubscriptionParams struct {
	OfferID        *string `json:"offer_id,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type CreateSubscriptionRecoveryLinkParams struct {
	IdempotencyKey string `json:"-"`
}

type UncancelSubscriptionParams struct {
	IdempotencyKey string `json:"-"`
}

type GetActiveSubscriptionParams struct {
	CustomerID string `json:"customer_id"`
}

type ListSubscriptionsParams struct {
	CustomerID *string             `json:"customer_id,omitempty"`
	Status     *SubscriptionStatus `json:"status,omitempty"`
}

type CreateSubscriptionParams struct {
	CustomerID      string         `json:"customer_id"`
	BillingInterval *string        `json:"billing_interval,omitempty"`
	PriceID         *string        `json:"price_id,omitempty"`
	InitialSeats    map[string]int `json:"initial_seats,omitempty"`
	Provider        *string        `json:"provider,omitempty"`
	Name            *string        `json:"name,omitempty"`
	StartDate       *string        `json:"start_date,omitempty"`
	SuccessURL      *string        `json:"success_url,omitempty"`
	OfferID         *string        `json:"offer_id,omitempty"`
	PromoCode       *string        `json:"promo_code,omitempty"`
	CustomTrialDays *int           `json:"custom_trial_days,omitempty"`
	SkipTrial       *bool          `json:"skip_trial,omitempty"`
	PlanID          *string        `json:"plan_id,omitempty"`
	PlanCode        *string        `json:"plan_code,omitempty"`
	CardPromotionID *string        `json:"card_promotion_id,omitempty"`
	IdempotencyKey  string         `json:"-"`
}

type SubscriptionsResource struct {
	http *httpClient
}

// Deactivate an add-on from a subscription.
func (r *SubscriptionsResource) DeactivateAddon(ctx context.Context, id string, addonID string) (*DeletedSubscriptionAddon, error) {
	return parseDirectResponse[DeletedSubscriptionAddon](r.http.delete(ctx, fmt.Sprintf("/subscriptions/%s/addons/%s", id, addonID), nil, ""))
}

// Activate an add-on on a subscription. Charges a prorated amount for the current billing period.
func (r *SubscriptionsResource) ActivateAddon(ctx context.Context, id string, params *ActivateAddonParams) (*SubscriptionAddon, error) {
	body := buildBody(map[string]any{
		"addon_id": params.AddonID,
	})
	return parseDirectResponse[SubscriptionAddon](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/addons", id), body, params.IdempotencyKey))
}

// Adjust a subscription's balance or credits by a signed amount. Positive adds, negative subtracts.
func (r *SubscriptionsResource) AdjustBalance(ctx context.Context, id string, params *AdjustBalanceParams) (*BalanceAdjustment, error) {
	body := buildBody(map[string]any{
		"amount": params.Amount,
		"reason": params.Reason,
		"type":   params.Type,
	})
	return parseDirectResponse[BalanceAdjustment](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/balance/adjust", id), body, params.IdempotencyKey))
}

// Top up a subscription's balance. Charges the customer's payment method for the specified amount.
func (r *SubscriptionsResource) TopupBalance(ctx context.Context, id string, params *TopupBalanceParams) (*BalanceTopup, error) {
	body := buildBody(map[string]any{
		"amount": params.Amount,
	})
	return parseDirectResponse[BalanceTopup](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/balance/topup", id), body, params.IdempotencyKey))
}

// Cancel immediately or at period end and return the updated subscription.
func (r *SubscriptionsResource) Cancel(ctx context.Context, id string, params *CancelSubscriptionParams) (*Subscription, error) {
	body := buildBody(map[string]any{
		"reason":    params.Reason,
		"immediate": params.Immediate,
	})
	return parseDirectResponse[Subscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/cancel", id), body, params.IdempotencyKey))
}

// Upgrade or change billing interval immediately, optionally applying an Offer. Scheduled changes do not accept offers.
func (r *SubscriptionsResource) ChangePlan(ctx context.Context, id string, params *ChangePlanParams) (*PlanChange, error) {
	body := buildBody(map[string]any{
		"new_plan_id":          params.NewPlanID,
		"new_billing_interval": params.NewBillingInterval,
		"success_url":          params.SuccessURL,
		"offer_id":             params.OfferID,
	})
	return parseDirectResponse[PlanChange](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/change-plan", id), body, params.IdempotencyKey))
}

// Purchase a credit pack for a subscription. Charges the customer and adds credits to their balance.
func (r *SubscriptionsResource) PurchaseCredits(ctx context.Context, id string, params *PurchaseCreditsParams) (*CreditGrant, error) {
	body := buildBody(map[string]any{
		"credit_pack_id": params.CreditPackID,
	})
	return parseDirectResponse[CreditGrant](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/credits", id), body, params.IdempotencyKey))
}

// Apply a direct Offer to a subscription. On a pending payment checkout it quotes or replaces the checkout discount and the existing checkout URL remains unchanged. On an active subscription it applies the Offer immediately with its discount phases starting at the next billing cycle; the call is rejected while another applied Offer still has active or upcoming discount phases. Offers with a free trial phase cannot be applied after checkout creation.
func (r *SubscriptionsResource) ApplyOffer(ctx context.Context, id string, params *ApplySubscriptionOfferParams) (*Subscription, error) {
	body := buildBody(map[string]any{
		"offer_id":   params.OfferID,
		"expires_at": params.ExpiresAt,
	})
	return parseDirectResponse[Subscription](r.http.put(ctx, fmt.Sprintf("/subscriptions/%s/offer", id), body, params.IdempotencyKey))
}

// Remove the quoted direct Offer from a subscription's pending payment checkout. The existing checkout URL remains unchanged and returns to its undiscounted price.
func (r *SubscriptionsResource) RemoveOffer(ctx context.Context, id string) (*Subscription, error) {
	return parseDirectResponse[Subscription](r.http.delete(ctx, fmt.Sprintf("/subscriptions/%s/offer", id), nil, ""))
}

// Creates a hosted checkout session for the customer to update the subscription's default payment method.
func (r *SubscriptionsResource) UpdatePaymentMethod(ctx context.Context, id string, params *UpdatePaymentMethodParams) (*PaymentMethodUpdateCheckout, error) {
	body := buildBody(map[string]any{
		"success_url": params.SuccessURL,
	})
	return parseDirectResponse[PaymentMethodUpdateCheckout](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/payment-method/update", id), body, params.IdempotencyKey))
}

// Preview proration details for an immediate plan change without applying it. Free-to-paid changes are never scheduled and the change-plan endpoint always returns hosted checkout for them. For paid plans, interval direction takes precedence: a longer interval is immediate and a shorter interval is scheduled. When the interval is unchanged, a higher-sort-order plan is immediate and a lower-sort-order plan is scheduled. A paid-to-free change is always scheduled. Returns credit, charge, and net amount. The target plan must belong to the same plan group as the current plan, otherwise a 400 with code `plans_not_in_same_group` is returned. A change between two free plans has nothing to prorate and returns a zero-amount estimate. Scheduled changes return a 400 with code `plan_change_scheduled`; apply those via the change-plan endpoint. Pass offerId to quote the destination plan with an Offer.
func (r *SubscriptionsResource) PreviewChange(ctx context.Context, id string, params *PreviewChangePlanParams) (*PreviewChange, error) {
	body := buildBody(map[string]any{
		"plan_id":          params.PlanID,
		"billing_interval": params.BillingInterval,
		"offer_id":         params.OfferID,
	})
	return parseDirectResponse[PreviewChange](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/preview-change", id), body, params.IdempotencyKey))
}

// Reactivates a subscription. A past_due subscription retries its outstanding renewal charge (recovering to active on success). A canceled subscription generates a fresh invoice, charges the saved card, and resets the billing period. On a successful charge the subscription becomes active; a declined charge returns an error with a recoveryUrl in the error details that can be sent to the customer to update their card. A canceled subscription may apply an Offer by offerId; past-due recovery cannot.
func (r *SubscriptionsResource) Reactivate(ctx context.Context, id string, params *ReactivateSubscriptionParams) (*ReactivatedSubscription, error) {
	body := buildBody(map[string]any{
		"offer_id": params.OfferID,
	})
	return parseDirectResponse[ReactivatedSubscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/reactivate", id), body, params.IdempotencyKey))
}

// Generates a hosted, signed recovery link that lets the customer pay the outstanding renewal charge for a past_due subscription. Unlike reactivate, which charges server-to-server, this returns a link the merchant can deliver through their own email, SMS, or dashboard. The link carries a self-contained signed token and stays valid until the charge is paid or the subscription is no longer past due.
func (r *SubscriptionsResource) CreateRecoveryLink(ctx context.Context, id string, params *CreateSubscriptionRecoveryLinkParams) (*RecoveryLink, error) {
	return parseDirectResponse[RecoveryLink](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/recovery-links", id), map[string]any{}, params.IdempotencyKey))
}

// Get a subscription by its public ID, regardless of status (including pending_payment and past_due).
func (r *SubscriptionsResource) Get(ctx context.Context, id string) (*Subscription, error) {
	return parseDirectResponse[Subscription](r.http.get(ctx, fmt.Sprintf("/subscriptions/%s", id), nil))
}

// Revert a scheduled cancellation and return the updated subscription. Only works before cancellation takes effect.
func (r *SubscriptionsResource) Uncancel(ctx context.Context, id string, params *UncancelSubscriptionParams) (*Subscription, error) {
	return parseDirectResponse[Subscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/uncancel", id), map[string]any{}, params.IdempotencyKey))
}

// Get the active subscription for a customer. Returns null if none.
func (r *SubscriptionsResource) GetActive(ctx context.Context, params *GetActiveSubscriptionParams) (*Subscription, error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseDirectResponse[Subscription](r.http.get(ctx, "/subscriptions/active", query))
}

// List all subscriptions. Filter by customer ID or status.
func (r *SubscriptionsResource) List(ctx context.Context, params *ListSubscriptionsParams) (*SubscriptionsListResult, error) {
	query := map[string]string{}
	if params.CustomerID != nil {
		query["customer_id"] = *params.CustomerID
	}
	if params.Status != nil {
		query["status"] = string(*params.Status)
	}
	return parseDirectResponse[SubscriptionsListResult](r.http.get(ctx, "/subscriptions", query))
}

// Create a subscription for a customer. Commet selects the default price when priceId is omitted and resolves its market from the customer's billing country. Without an offer override, Commet applies the price's automatic introductory Offer. Pass offerId to apply an active compatible Offer directly, or cardPromotionId to preselect a card-eligible Promotional Offer for the initial checkout when card promotions are enabled for the organization. For the initial checkout, provider accepts either a processor name or an exact payment connection ID.
func (r *SubscriptionsResource) Create(ctx context.Context, params *CreateSubscriptionParams) (*CreatedSubscription, error) {
	body := buildBody(map[string]any{
		"customer_id":       params.CustomerID,
		"billing_interval":  params.BillingInterval,
		"price_id":          params.PriceID,
		"initial_seats":     params.InitialSeats,
		"provider":          params.Provider,
		"name":              params.Name,
		"start_date":        params.StartDate,
		"success_url":       params.SuccessURL,
		"offer_id":          params.OfferID,
		"promo_code":        params.PromoCode,
		"custom_trial_days": params.CustomTrialDays,
		"skip_trial":        params.SkipTrial,
		"plan_id":           params.PlanID,
		"plan_code":         params.PlanCode,
		"card_promotion_id": params.CardPromotionID,
	})
	return parseDirectResponse[CreatedSubscription](r.http.post(ctx, "/subscriptions", body, params.IdempotencyKey))
}
