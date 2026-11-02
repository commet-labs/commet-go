# Changelog

## 9.8.0

- `Subscriptions.PurchaseCredits` buys a chosen `Quantity` of credits at the plan's price, besides a `CreditPackID`.
- `Subscriptions.GetAutoTopUp` and `Subscriptions.UpdateAutoTopUp` read and configure automatic credit purchases.
- `Subscriptions.ChangePlan` and `Subscriptions.PreviewChange` accept `PromoCode`, as an alternative to `OfferID`.
- `Subscriptions.RevokeScheduledPlanChange` keeps a subscription on its current plan by revoking the change scheduled for the end of the period.
- `Customers.Create` and `Customers.CreateBatch` accept `PlanGroupID`, so a customer starts on that plan group's starting plan.
- `Plans.Create` and `Plans.Update` accept `UnusedPlanCredits`, `PlanCreditLimit` and `IsDefault`, and plan prices accept `TrialRequiresPaymentMethod`.
- `PromoCodes.Create` and `PromoCodes.Update` accept `FirstPurchaseOnly`.
- Creating a subscription accepts a custom trial length or skipping the trial together with an Offer.
- Responses add whether a subscription has a payment method, which integration manages an Offer or promo code, whether a promo code is limited to a first purchase, the unused plan credit and plan credit limit settings of plans, and the plan credit limit window of usage checks and events.
- Creating or updating an Offer no longer accepts a `free_trial` phase, which the API already rejects; trials belong to plan prices. The generated offer phase variants are renumbered: `CreateOfferParamsPhasesItemVariant1`, `Variant2` and `Variant3` are now the percentage, amount-off and fixed-price phases, and the same applies to `UpdateOfferParamsPhasesItem` and `OfferPhasesItem`.

## 9.7.0

- Add paying billing periods in advance: create a prepaid-periods checkout for a subscription, and set `prepaidPeriods` and `cancelAtPeriodEnd` when creating one.
- Support the `prepayment` invoice type and payment reason, the `prepaid_period` invoice line type, and the `google_pay` and `smartpix` payment methods.

## 9.6.0

- Add the `semiannual` billing interval.

## 9.5.0

- Install the webhooks contract (`docs/webhooks.openapi.json`) with the SDK documentation in place of per-event pages, and refresh the documentation from commet.co/docs.

## 9.4.0

- Add subscription pause, pause updates, pause revocation, resume, and their lifecycle webhooks, including indefinite pauses.
- Support restricted API keys and permission-scoped key creation.
- Expose payment context, payment methods, and subscription payment methods from API `2026-08-27`.
- Refresh the installed API documentation to match the released contract.

## 9.3.0

- Align currency inputs and active-subscription Offer behavior with the current API contract.
- Install the complete version-matched Commet documentation with the SDK module.

## 9.2.0

- Add customer credits and Plan Grant management, including effective versus base feature access and Plan Grant lifecycle webhooks.
- Add pending-checkout Offer application, card-promotion selection, exact payment-connection routing, and provider details on payment recovery webhooks.
- Expose durable test-clock runs and latest-run state.
- Align deprecated payout verification and manual test-clock billing methods with their current no-payload contracts.
- Preserve exact server request IDs on API errors and never fabricate correlation IDs locally.

## 9.0.0

- Pin requests to API `2026-07-31`.
- Replace pricing Market Groups with the top-level Markets resource.
- Make Offers independent from purposes and plan-price associations.
- Expose target-aware Offer Applications on subscriptions.

## 8.0.1

- Document Offers, reusable pricing Market Groups, and selectable price variants.
- Clarify plan-change preview behavior for plan sort order and billing intervals.

## 8.0.0

- Add pricing markets and selectable price variants, including market groups and explicit price selection.
- Add Promotional Offers and the current usage check, track, and set operations.
- Add webhook endpoint management while preserving signature verification and event dispatch helpers.
- Return typed resources and list envelopes directly from API methods.
- Align subscription, feature access, invoice, payout, seat, and transaction types with API v8.
- Rename legacy request fields to their v8 replacements, including `featureCode` and `offerId`.
