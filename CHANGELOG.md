# Changelog

## 9.12.0

- `Subscriptions.PurchaseCredits`, `Subscriptions.TopupBalance` and `Subscriptions.ActivateAddon` accept `SuccessURL`. When the subscription has no saved payment method, they return the status `requires_checkout` with a `CheckoutURL` where the customer pays and saves the method in one step.
- `Subscriptions.ChangePlan` can return the outcome `processing` while its charge is being confirmed.
- Payment methods include `apple_pay` and `link`.

## 9.11.0

- Plans can give unlimited plan credits or plan balance. `Plans.Create` and `Plans.Update` accept `UnlimitedPlanCredits` and `UnlimitedPlanBalance`, draft rules accept them too, and plans return them. Subscriptions, usage checks, usage events and feature access return `Unlimited`, and omit the plan and total amounts when it is true.
- Credit purchases, balance top-ups and add-on activations return a `Status`: `processing` when the payment method confirms later, and the credits, balance or add-on arrive once it is confirmed.
- Transactions return `AmountRefunded`, and `Transactions.Refund` accepts `Amount` for a partial refund.
- A version move can charge a subscriber in USD: `Subscriptions.ChangeVersion`, `Plans.CreateMove`, `Plans.PublishDraft` and `Plans.ArchiveExperiment` accept `Currency`, and the scheduled version change returns it. New errors `currency_held_by_payment_method`, `currency_held_by_customer_credit`, `currency_held_by_balance_plan`, `currency_held_by_fixed_amount_discount`, `currency_held_by_prepaid_period` and `currency_held_by_scheduled_pause`.
- `Customers.CreateBatch` accepts `BusinessTaxID` for each customer.
- `Invoices.UpdateStatus` is deprecated and no longer takes a status: the API no longer changes the invoice, which is paid when its charge settles.

## 9.10.0

- Plans have versions. Editing a plan publishes a new version, and existing subscriptions stay on theirs until moved. `Plans.ListVersions`, `Plans.GetVersion`, `Plans.CompareVersions` and `Plans.CreateVersion` read and publish versions; `Plans.CreateDraft`, `Plans.UpdateDraft`, `Plans.PublishDraft` and the other draft methods prepare one; `Plans.CreateExperiment`, `Plans.MakeExperimentCurrent`, `Plans.ArchiveExperiment` and the other experiment methods sell a version through its own code; `Plans.CreateMove`, `Plans.CancelMove` and the other move methods move subscribers to a newer version; `Plans.ListSubscribersNeedingAChange` and `Plans.KeepSubscriberOnVersion` handle the subscribers that do not fit it.
- `Subscriptions.ChangeVersion` moves one subscription to a newer version of its plan, now or at the end of the period, and `Subscriptions.RevokeScheduledVersionChange` cancels a pending move. Subscriptions return `PlanVersion` and `ScheduledVersionChange`.
- `Subscriptions.Create` accepts `PlanVersionID`, and `PlanCode` as `<plan code>:<experiment code>` to sell an experiment. `Subscriptions.List` filters by `PlanVersionID`.
- Plans return `Countries`. `Addons.Create`, `Addons.Update`, `CreditPacks.Create` and `CreditPacks.Update` accept `PlanIds`, and add-ons and credit packs return them. `Subscriptions.PreviewChange` returns `DeactivatedAddons`.
- New webhook events `subscription.version_changed` and `subscription.version_change_canceled`.
- New errors `plan_version_not_found`, `plan_experiment_not_found`, `subscribers_do_not_fit`, `version_not_newer`, `addon_unavailable` and `credit_pack_unavailable`.

## 9.9.0

- Subscriptions on a balance plan split their balance in two parts: the plan balance (spent first, replaced or rolled over at each reset) and the purchased balance (top-ups and manual adjustments, kept until the subscription ends), plus the total. Credits subscriptions add their plan credits and total. The existing remaining amount keeps its meaning.
- `Plans.Create` and `Plans.Update` accept `UnusedPlanBalance` (expire or roll over up to a maximum plan balance) and `PlanBalanceLimit` (the most plan balance spent per day or week), and `Plans.Update` accepts `BlockOnExhaustion` to turn overage on or off. Plans return both settings.
- `Subscriptions.UpdateAutoTopUp` accepts `Amount` for balance plans and `Subscriptions.GetAutoTopUp` returns it, so balance subscriptions top up automatically; `amount_below_minimum` is a new failure reason.
- Usage checks, usage events and the `customer.state_changed` balance add the plan balance limit window and the plan, purchased and total balance parts.
- New webhook events `balance.granted`, `balance.expired` and `payment_method.detached`.
- `Customers.Create` and `Customers.Update` accept `BusinessTaxID`, and customers return it.
- `Subscriptions.Create` accepts `PaymentMethodTypes` to limit the payment methods its initial checkout offers.

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
