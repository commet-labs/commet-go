package commet

import (
	"context"
	"fmt"
)

type ActivateAddonParams struct {
	AddonID        string `json:"addon_id"`
	IdempotencyKey string `json:"-"`
}

type UpdateSubscriptionAutoTopUpParams struct {
	Enabled        *bool  `json:"enabled,omitempty"`
	Threshold      *int   `json:"threshold,omitempty"`
	Quantity       *int   `json:"quantity,omitempty"`
	Amount         *int   `json:"amount,omitempty"`
	MaxPerPeriod   *int   `json:"max_per_period,omitempty"`
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
	PromoCode          *string `json:"promo_code,omitempty"`
	IdempotencyKey     string  `json:"-"`
}

type ChangeSubscriptionVersionParams struct {
	PlanVersionID  string `json:"plan_version_id"`
	Effective      string `json:"effective"`
	IdempotencyKey string `json:"-"`
}

type PurchaseCreditsParams struct {
	CreditPackID   *string `json:"credit_pack_id,omitempty"`
	Quantity       *int    `json:"quantity,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type ApplySubscriptionOfferParams struct {
	OfferID        string  `json:"offer_id"`
	ExpiresAt      *string `json:"expires_at,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type PauseSubscriptionParams struct {
	Mode           string `json:"mode"`
	DurationDays   *int   `json:"duration_days"`
	IdempotencyKey string `json:"-"`
}

type UpdateSubscriptionPauseParams struct {
	DurationDays   *int   `json:"duration_days"`
	IdempotencyKey string `json:"-"`
}

type UpdatePaymentMethodParams struct {
	SuccessURL     *string `json:"success_url,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type CreatePrepaidPeriodsCheckoutParams struct {
	Periods        int     `json:"periods"`
	OfferID        *string `json:"offer_id,omitempty"`
	SuccessURL     *string `json:"success_url,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type PreviewChangePlanParams struct {
	PlanID          string  `json:"plan_id"`
	BillingInterval *string `json:"billing_interval,omitempty"`
	OfferID         *string `json:"offer_id,omitempty"`
	PromoCode       *string `json:"promo_code,omitempty"`
	IdempotencyKey  string  `json:"-"`
}

type ReactivateSubscriptionParams struct {
	OfferID        *string `json:"offer_id,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type CreateSubscriptionRecoveryLinkParams struct {
	IdempotencyKey string `json:"-"`
}

type ResumeSubscriptionParams struct {
	IdempotencyKey string `json:"-"`
}

type UncancelSubscriptionParams struct {
	IdempotencyKey string `json:"-"`
}

type GetActiveSubscriptionParams struct {
	CustomerID string `json:"customer_id"`
}

type ListSubscriptionsParams struct {
	CustomerID    *string             `json:"customer_id,omitempty"`
	Status        *SubscriptionStatus `json:"status,omitempty"`
	PlanVersionID *string             `json:"plan_version_id,omitempty"`
}

type CreateSubscriptionParams struct {
	CustomerID         string         `json:"customer_id"`
	BillingInterval    *string        `json:"billing_interval,omitempty"`
	PriceID            *string        `json:"price_id,omitempty"`
	InitialSeats       map[string]int `json:"initial_seats,omitempty"`
	Provider           *string        `json:"provider,omitempty"`
	PaymentMethodTypes []string       `json:"payment_method_types,omitempty"`
	Name               *string        `json:"name,omitempty"`
	StartDate          *string        `json:"start_date,omitempty"`
	SuccessURL         *string        `json:"success_url,omitempty"`
	PrepaidPeriods     *int           `json:"prepaid_periods,omitempty"`
	CancelAtPeriodEnd  *bool          `json:"cancel_at_period_end,omitempty"`
	PlanVersionID      *string        `json:"plan_version_id,omitempty"`
	OfferID            *string        `json:"offer_id,omitempty"`
	PromoCode          *string        `json:"promo_code,omitempty"`
	CustomTrialDays    *int           `json:"custom_trial_days,omitempty"`
	SkipTrial          *bool          `json:"skip_trial,omitempty"`
	PlanID             *string        `json:"plan_id,omitempty"`
	PlanCode           *string        `json:"plan_code,omitempty"`
	CardPromotionID    *string        `json:"card_promotion_id,omitempty"`
	IdempotencyKey     string         `json:"-"`
}

type SubscriptionsResource struct {
	http *httpClient
}

// Deactivate an add-on from a subscription.
func (r *SubscriptionsResource) DeactivateAddon(ctx context.Context, id string, addonID string) (*DeletedSubscriptionAddon, error) {
	return parseDirectResponse[DeletedSubscriptionAddon](r.http.delete(ctx, fmt.Sprintf("/subscriptions/%s/addons/%s", id, addonID), nil, ""))
}

// Activate an add-on on a subscription. Charges a prorated amount for the current billing period. An add-on that the subscription's plan version does not sell is rejected with addon_unavailable.
func (r *SubscriptionsResource) ActivateAddon(ctx context.Context, id string, params *ActivateAddonParams) (*SubscriptionAddon, error) {
	body := buildBody(map[string]any{
		"addon_id": params.AddonID,
	})
	return parseDirectResponse[SubscriptionAddon](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/addons", id), body, params.IdempotencyKey))
}

// Read whether the subscription tops up automatically with its saved payment method when its available credits or balance reach a threshold, with the top-ups used this period and the last failure.
func (r *SubscriptionsResource) GetAutoTopUp(ctx context.Context, id string) (*SubscriptionAutoTopUp, error) {
	return parseDirectResponse[SubscriptionAutoTopUp](r.http.get(ctx, fmt.Sprintf("/subscriptions/%s/auto-top-up", id), nil))
}

// Turn automatic top-ups on or off and set the threshold, the size of each top-up and the maximum top-ups per period. Send only the fields that change. On a credits plan each top-up buys `quantity` credits; on a balance plan it adds `amount` in minor units of the subscription currency, plus tax, and `threshold` is in the same minor units. Enabling requires a plan that sells credits or has balance top-ups on. The size must be at least the plan's minimum when auto top-up is enabled or when the size changes, so turning it off keeps working after the minimum rises. Saving it enabled resumes a paused auto top-up.
func (r *SubscriptionsResource) UpdateAutoTopUp(ctx context.Context, id string, params *UpdateSubscriptionAutoTopUpParams) (*SubscriptionAutoTopUp, error) {
	body := buildBody(map[string]any{
		"enabled":        params.Enabled,
		"threshold":      params.Threshold,
		"quantity":       params.Quantity,
		"amount":         params.Amount,
		"max_per_period": params.MaxPerPeriod,
	})
	return parseDirectResponse[SubscriptionAutoTopUp](r.http.patch(ctx, fmt.Sprintf("/subscriptions/%s/auto-top-up", id), body, params.IdempotencyKey))
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

// Top up a balance subscription. Charges the customer's payment method for the amount plus tax, adds it to the subscription's purchased balance (after paying back any overage owed on the plan balance), creates its invoice and emits `balance.topped_up`. Purchased balance is spent after the plan balance and is kept until the subscription ends. Rejected with 400 when the subscription is not on a balance plan, its plan has top-ups turned off, or the amount is below the effective minimum: the plan's minimum top-up in the subscription currency, raised to the currency's minimum charge.
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

// Upgrade or change billing interval immediately, optionally applying an Offer by offerId or the Offer of a promoCode. Scheduled changes accept neither offerId nor promoCode. From a free plan, a promoCode opens the checkout with the code applied and counts its use when the checkout is paid; without offerId or promoCode, a destination price whose free trial requires no payment method starts that trial right away without checkout. During a trial without a payment method, an immediate change returns 409 no_payment_method; scheduled changes still apply.
func (r *SubscriptionsResource) ChangePlan(ctx context.Context, id string, params *ChangePlanParams) (*PlanChange, error) {
	body := buildBody(map[string]any{
		"new_plan_id":          params.NewPlanID,
		"new_billing_interval": params.NewBillingInterval,
		"success_url":          params.SuccessURL,
		"offer_id":             params.OfferID,
		"promo_code":           params.PromoCode,
	})
	return parseDirectResponse[PlanChange](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/change-plan", id), body, params.IdempotencyKey))
}

// Move a subscription to a newer published version of its plan, now or at the end of the period, without proration; an older version is rejected with version_not_newer, and a subscription on a running experiment moves only to a later version of that experiment. The subscription keeps its currency and its price: a destination that does not sell its currency is rejected with missing_regional_price, and one that does not sell the price it pays with price_not_found. A free version moves only to a free version and a paid one only to a paid one (validation_error). A subscription whose usage does not fit the destination, or with an active add-on the destination does not sell, is not moved either, and the error says why. Each change is recorded as a move of this one subscription in the plan's moves. Repeating the same request changes nothing; a new one replaces the pending move.
func (r *SubscriptionsResource) ChangeVersion(ctx context.Context, id string, params *ChangeSubscriptionVersionParams) (*Subscription, error) {
	body := buildBody(map[string]any{
		"plan_version_id": params.PlanVersionID,
		"effective":       params.Effective,
	})
	return parseDirectResponse[Subscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/change-version", id), body, params.IdempotencyKey))
}

// Purchase a fixed credit pack or choose a whole credit quantity at the subscription plan price and minimum. Charges the saved payment method and adds purchased credits. A credit pack that the subscription's plan version does not sell is rejected with credit_pack_unavailable.
func (r *SubscriptionsResource) PurchaseCredits(ctx context.Context, id string, params *PurchaseCreditsParams) (*CreditGrant, error) {
	body := buildBody(map[string]any{
		"credit_pack_id": params.CreditPackID,
		"quantity":       params.Quantity,
	})
	return parseDirectResponse[CreditGrant](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/credits", id), body, params.IdempotencyKey))
}

// Apply a direct Offer to a subscription. On a pending checkout it quotes or replaces the checkout discount, which starts after the plan price trial when there is one, and the existing checkout URL remains unchanged. On an active subscription it applies the Offer immediately with its discount phases starting at the next billing cycle; the call is rejected while another applied Offer still has active or upcoming discount phases.
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

// Pause immediately or schedule a pause for the end of the current billing or trial period. Set durationDays to null for an indefinite pause.
func (r *SubscriptionsResource) Pause(ctx context.Context, id string, params *PauseSubscriptionParams) (*Subscription, error) {
	body := buildBody(map[string]any{
		"mode": params.Mode,
	})
	body["duration_days"] = params.DurationDays
	return parseDirectResponse[Subscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/pause", id), body, params.IdempotencyKey))
}

// Change the duration of a scheduled or active pause. Set durationDays to null to make it indefinite.
func (r *SubscriptionsResource) UpdatePause(ctx context.Context, id string, params *UpdateSubscriptionPauseParams) (*Subscription, error) {
	body := buildBody(map[string]any{})
	body["duration_days"] = params.DurationDays
	return parseDirectResponse[Subscription](r.http.patch(ctx, fmt.Sprintf("/subscriptions/%s/pause", id), body, params.IdempotencyKey))
}

// Revoke a pause before it becomes effective. Active pauses must be resumed instead.
func (r *SubscriptionsResource) RevokePause(ctx context.Context, id string) (*Subscription, error) {
	return parseDirectResponse[Subscription](r.http.delete(ctx, fmt.Sprintf("/subscriptions/%s/pause", id), nil, ""))
}

// Creates a hosted checkout to update the payment method of an active, trialing or past_due subscription. Pending-payment subscriptions must use their payment checkout.
func (r *SubscriptionsResource) UpdatePaymentMethod(ctx context.Context, id string, params *UpdatePaymentMethodParams) (*PaymentMethodUpdateCheckout, error) {
	body := buildBody(map[string]any{
		"success_url": params.SuccessURL,
	})
	return parseDirectResponse[PaymentMethodUpdateCheckout](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/payment-method/update", id), body, params.IdempotencyKey))
}

// Creates a hosted checkout to pay billing periods of a recurring subscription in advance with any payment method. On an active subscription the paid periods start after the last paid period and replace any open prepaid-periods checkout. On a pending-payment subscription the first payment covers the requested periods and replaces its pending checkout. The billing interval does not change.
func (r *SubscriptionsResource) CreatePrepaidPeriodsCheckout(ctx context.Context, id string, params *CreatePrepaidPeriodsCheckoutParams) (*PrepaidPeriodsCheckout, error) {
	body := buildBody(map[string]any{
		"periods":     params.Periods,
		"offer_id":    params.OfferID,
		"success_url": params.SuccessURL,
	})
	return parseDirectResponse[PrepaidPeriodsCheckout](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/prepaid-periods", id), body, params.IdempotencyKey))
}

// Preview proration details for an immediate plan change without applying it. Free-to-paid changes are never scheduled and the change-plan endpoint returns hosted checkout for them, except a change without offerId or promoCode to a price whose free trial needs no payment method, for a customer who never had that plan: it starts the trial at once, so the preview reports nothing due today. For paid plans, interval direction takes precedence: a longer interval is immediate and a shorter interval is scheduled. When the interval is unchanged, a higher-sort-order plan is immediate and a lower-sort-order plan is scheduled. A paid-to-free change is always scheduled. Returns credit, charge, and net amount. The target plan must belong to the same plan group as the current plan, otherwise a 400 with code `plans_not_in_same_group` is returned. A change between two free plans has nothing to prorate and returns a zero-amount estimate. Scheduled changes return a 400 with code `plan_change_scheduled`; apply those via the change-plan endpoint. Pass offerId to quote the destination plan with an Offer, or promoCode to quote it with the Offer of a promo code validated with the checkout's rules.
func (r *SubscriptionsResource) PreviewChange(ctx context.Context, id string, params *PreviewChangePlanParams) (*PreviewChange, error) {
	body := buildBody(map[string]any{
		"plan_id":          params.PlanID,
		"billing_interval": params.BillingInterval,
		"offer_id":         params.OfferID,
		"promo_code":       params.PromoCode,
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

// Resume a paused subscription. Immediate pauses continue the preserved period without a charge. Period-end pauses charge a new period before access is restored.
func (r *SubscriptionsResource) Resume(ctx context.Context, id string, params *ResumeSubscriptionParams) (*SubscriptionResume, error) {
	return parseDirectResponse[SubscriptionResume](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/resume", id), map[string]any{}, params.IdempotencyKey))
}

// Get a subscription by its public ID, regardless of status (including pending_payment and past_due).
func (r *SubscriptionsResource) Get(ctx context.Context, id string) (*Subscription, error) {
	return parseDirectResponse[Subscription](r.http.get(ctx, fmt.Sprintf("/subscriptions/%s", id), nil))
}

// Revoke the plan or interval change scheduled for the end of the period, so the subscription stays on its current plan. Returns the updated subscription.
func (r *SubscriptionsResource) RevokeScheduledPlanChange(ctx context.Context, id string) (*Subscription, error) {
	return parseDirectResponse[Subscription](r.http.delete(ctx, fmt.Sprintf("/subscriptions/%s/scheduled-plan-change", id), nil, ""))
}

// Cancel the move to another plan version pending for the next renewal, so the subscription stays on its version, and send subscription.version_change_canceled with reason canceled_by_merchant. A move made now has nothing pending to revoke. With nothing pending it changes nothing. Returns the subscription.
func (r *SubscriptionsResource) RevokeScheduledVersionChange(ctx context.Context, id string) (*Subscription, error) {
	return parseDirectResponse[Subscription](r.http.delete(ctx, fmt.Sprintf("/subscriptions/%s/scheduled-version-change", id), nil, ""))
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
	if params.PlanVersionID != nil {
		query["plan_version_id"] = *params.PlanVersionID
	}
	return parseDirectResponse[SubscriptionsListResult](r.http.get(ctx, "/subscriptions", query))
}

// Create a subscription for a customer. Commet selects the default price when priceId is omitted. The subscription starts at the base USD price; market and currency are resolved from the buyer's request country when the checkout opens. Without an offer override, Commet applies the price's automatic introductory Offer. Pass offerId to apply an active compatible Offer directly, or cardPromotionId to preselect a card-eligible Promotional Offer for the initial checkout when card promotions are enabled for the organization. For the initial checkout, provider accepts either a processor name or an exact payment connection ID. When the price's free trial requires no payment method, the subscription starts trialing right away with no checkout (checkoutUrl null, hasPaymentMethod false); it cannot be combined with offerId, promoCode, cardPromotionId or provider, and an organization that cannot charge yet receives 403.
func (r *SubscriptionsResource) Create(ctx context.Context, params *CreateSubscriptionParams) (*CreatedSubscription, error) {
	body := buildBody(map[string]any{
		"customer_id":          params.CustomerID,
		"billing_interval":     params.BillingInterval,
		"price_id":             params.PriceID,
		"initial_seats":        params.InitialSeats,
		"provider":             params.Provider,
		"payment_method_types": params.PaymentMethodTypes,
		"name":                 params.Name,
		"start_date":           params.StartDate,
		"success_url":          params.SuccessURL,
		"prepaid_periods":      params.PrepaidPeriods,
		"cancel_at_period_end": params.CancelAtPeriodEnd,
		"plan_version_id":      params.PlanVersionID,
		"offer_id":             params.OfferID,
		"promo_code":           params.PromoCode,
		"custom_trial_days":    params.CustomTrialDays,
		"skip_trial":           params.SkipTrial,
		"plan_id":              params.PlanID,
		"plan_code":            params.PlanCode,
		"card_promotion_id":    params.CardPromotionID,
	})
	return parseDirectResponse[CreatedSubscription](r.http.post(ctx, "/subscriptions", body, params.IdempotencyKey))
}
