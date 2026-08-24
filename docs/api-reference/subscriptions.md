# Subscriptions

API version: `2026-07-31`

## DeactivateAddon

`client.Subscriptions.DeactivateAddon(ctx, ...)`

`DELETE /subscriptions/{id}/addons/{addonId}` · operation `deactivate-addon`

Deactivate an add-on from a subscription.

### Parameters

- `ID` (`string`, required)
- `AddonID` (`string`, required)

### Returns

`DeletedSubscriptionAddon`

## ActivateAddon

`client.Subscriptions.ActivateAddon(ctx, ...)`

`POST /subscriptions/{id}/addons` · operation `activate-addon`

Activate an add-on on a subscription. Charges a prorated amount for the current billing period.

### Parameters

- `ID` (`string`, required)
- `AddonID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`SubscriptionAddon`

## AdjustBalance

`client.Subscriptions.AdjustBalance(ctx, ...)`

`POST /subscriptions/{id}/balance/adjust` · operation `adjust-balance`

Adjust a subscription's balance or credits by a signed amount. Positive adds, negative subtracts.

### Parameters

- `ID` (`string`, required)
- `Amount` (`int`, required)
- `Reason` (`string`, optional)
- `Type` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`BalanceAdjustment`

## TopupBalance

`client.Subscriptions.TopupBalance(ctx, ...)`

`POST /subscriptions/{id}/balance/topup` · operation `topup-balance`

Top up a subscription's balance. Charges the customer's payment method for the specified amount.

### Parameters

- `ID` (`string`, required)
- `Amount` (`int`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`BalanceTopup`

## Cancel

`client.Subscriptions.Cancel(ctx, ...)`

`POST /subscriptions/{id}/cancel` · operation `cancel-subscription`

Cancel immediately or at period end and return the updated subscription.

### Parameters

- `ID` (`string`, required)
- `Reason` (`string`, optional)
- `Immediate` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Subscription`

## ChangePlan

`client.Subscriptions.ChangePlan(ctx, ...)`

`POST /subscriptions/{id}/change-plan` · operation `change-plan`

Upgrade or change billing interval immediately, optionally applying an Offer. Scheduled changes do not accept offers.

### Parameters

- `ID` (`string`, required)
- `NewPlanID` (`string`, optional)
- `NewBillingInterval` (`string`, optional)
- `SuccessURL` (`string`, optional)
- `OfferID` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanChange`

## PurchaseCredits

`client.Subscriptions.PurchaseCredits(ctx, ...)`

`POST /subscriptions/{id}/credits` · operation `purchase-credits`

Purchase a credit pack for a subscription. Charges the customer and adds credits to their balance.

### Parameters

- `ID` (`string`, required)
- `CreditPackID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CreditGrant`

## ApplyOffer

`client.Subscriptions.ApplyOffer(ctx, ...)`

`PUT /subscriptions/{id}/offer` · operation `apply-subscription-offer`

Apply a direct Offer to a subscription. On a pending payment checkout it quotes or replaces the checkout discount and the existing checkout URL remains unchanged. On an active subscription it applies the Offer immediately with its discount phases starting at the next billing cycle; the call is rejected while another applied Offer still has active or upcoming discount phases. Offers with a free trial phase cannot be applied after checkout creation.

### Parameters

- `ID` (`string`, required)
- `OfferID` (`string`, required)
- `ExpiresAt` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Subscription`

## RemoveOffer

`client.Subscriptions.RemoveOffer(ctx, ...)`

`DELETE /subscriptions/{id}/offer` · operation `remove-subscription-offer`

Remove the quoted direct Offer from a subscription's pending payment checkout. The existing checkout URL remains unchanged and returns to its undiscounted price.

### Parameters

- `ID` (`string`, required)

### Returns

`Subscription`

## UpdatePaymentMethod

`client.Subscriptions.UpdatePaymentMethod(ctx, ...)`

`POST /subscriptions/{id}/payment-method/update` · operation `update-payment-method`

Creates a hosted checkout session for the customer to update the subscription's default payment method.

### Parameters

- `ID` (`string`, required)
- `SuccessURL` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PaymentMethodUpdateCheckout`

## PreviewChange

`client.Subscriptions.PreviewChange(ctx, ...)`

`POST /subscriptions/{id}/preview-change` · operation `preview-change-plan`

Preview proration details for an immediate plan change without applying it. Free-to-paid changes are never scheduled and the change-plan endpoint always returns hosted checkout for them. For paid plans, interval direction takes precedence: a longer interval is immediate and a shorter interval is scheduled. When the interval is unchanged, a higher-sort-order plan is immediate and a lower-sort-order plan is scheduled. A paid-to-free change is always scheduled. Returns credit, charge, and net amount. The target plan must belong to the same plan group as the current plan, otherwise a 400 with code `plans_not_in_same_group` is returned. A change between two free plans has nothing to prorate and returns a zero-amount estimate. Scheduled changes return a 400 with code `plan_change_scheduled`; apply those via the change-plan endpoint. Pass offerId to quote the destination plan with an Offer.

### Parameters

- `ID` (`string`, required)
- `PlanID` (`string`, required)
- `BillingInterval` (`string`, optional)
- `OfferID` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PreviewChange`

## Reactivate

`client.Subscriptions.Reactivate(ctx, ...)`

`POST /subscriptions/{id}/reactivate` · operation `reactivate-subscription`

Reactivates a subscription. A past_due subscription retries its outstanding renewal charge (recovering to active on success). A canceled subscription generates a fresh invoice, charges the saved card, and resets the billing period. On a successful charge the subscription becomes active; a declined charge returns an error with a recoveryUrl in the error details that can be sent to the customer to update their card. A canceled subscription may apply an Offer by offerId; past-due recovery cannot.

### Parameters

- `ID` (`string`, required)
- `OfferID` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`ReactivatedSubscription`

## CreateRecoveryLink

`client.Subscriptions.CreateRecoveryLink(ctx, ...)`

`POST /subscriptions/{id}/recovery-links` · operation `create-subscription-recovery-link`

Generates a hosted, signed recovery link that lets the customer pay the outstanding renewal charge for a past_due subscription. Unlike reactivate, which charges server-to-server, this returns a link the merchant can deliver through their own email, SMS, or dashboard. The link carries a self-contained signed token and stays valid until the charge is paid or the subscription is no longer past due.

### Parameters

- `ID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`RecoveryLink`

## Get

`client.Subscriptions.Get(ctx, ...)`

`GET /subscriptions/{id}` · operation `get-subscription`

Get a subscription by its public ID, regardless of status (including pending_payment and past_due).

### Parameters

- `ID` (`string`, required)

### Returns

`Subscription`

## Uncancel

`client.Subscriptions.Uncancel(ctx, ...)`

`POST /subscriptions/{id}/uncancel` · operation `uncancel-subscription`

Revert a scheduled cancellation and return the updated subscription. Only works before cancellation takes effect.

### Parameters

- `ID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Subscription`

## GetActive

`client.Subscriptions.GetActive(ctx, ...)`

`GET /subscriptions/active` · operation `get-active-subscription`

Get the active subscription for a customer. Returns null if none.

### Parameters

- `CustomerID` (`string`, required)

### Returns

`*Subscription`

## List

`client.Subscriptions.List(ctx, ...)`

`GET /subscriptions` · operation `list-subscriptions`

List all subscriptions. Filter by customer ID or status.

### Parameters

- `CustomerID` (`string`, optional)
- `Status` (`SubscriptionStatus`, optional)

### Returns

`SubscriptionsListResult`

## Create

`client.Subscriptions.Create(ctx, ...)`

`POST /subscriptions` · operation `create-subscription`

Create a subscription for a customer. Commet selects the default price when priceId is omitted and resolves its market from the customer's billing country. Without an offer override, Commet applies the price's automatic introductory Offer. Pass offerId to apply an active compatible Offer directly, or cardPromotionId to preselect a card-eligible Promotional Offer for the initial checkout when card promotions are enabled for the organization. For the initial checkout, provider accepts either a processor name or an exact payment connection ID.

### Parameters

- `CustomerID` (`string`, required)
- `BillingInterval` (`string | null`, optional)
- `PriceID` (`string`, optional) — Public price ID. When omitted, Commet selects the default price for the billing interval and still applies its market pricing.
- `InitialSeats` (`map[string]int`, optional)
- `Provider` (`string`, optional) — Payment provider name or exact public payment connection ID for the initial checkout. Overrides country routing when present.
- `Name` (`string`, optional)
- `StartDate` (`string`, optional)
- `SuccessURL` (`string`, optional)
- `OfferID` (`string`, optional)
- `PromoCode` (`string`, optional)
- `CustomTrialDays` (`int`, optional)
- `SkipTrial` (`bool`, optional)
- `PlanID` (`string`, optional)
- `PlanCode` (`string`, optional)
- `CardPromotionID` (`string`, optional) — Public card promotion ID. The offer is shown immediately and remains conditional on card eligibility until checkout confirmation.

### Valid parameter combinations

- `CustomerID` + `PlanID`
- `CustomerID` + `PlanCode`
- `CustomerID` + `OfferID` + `PlanID`
- `CustomerID` + `OfferID` + `PlanCode`
- `CustomerID` + `CardPromotionID` + `PlanID`
- `CustomerID` + `CardPromotionID` + `PlanCode`

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CreatedSubscription`
