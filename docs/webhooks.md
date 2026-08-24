# Webhooks

Generated from Commet API version `2026-07-31`.

## subscription.created

Fired when a subscription record is created with status pending_payment. The first charge has not been confirmed yet — do NOT grant access here. Wait for subscription.activated.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `planId` (`string`)
- `planName` (`string`)
- `status` (`string`)
- `startDate` (`string`)
- `name` (`string`)

## subscription.activated

Fired once, when the subscription's first charge succeeds and it becomes active — this is where you grant access. Never re-fired on renewals; use payment.received for per-charge notifications.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `currentPeriodStart` (`string`)
- `currentPeriodEnd` (`string`)
- `name` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `invoiceTotal` (`float64`)
- `invoiceCurrency` (`string`)
- `provider` (`string`)

## subscription.reactivated

Fired when a canceled subscription is reactivated and its reactivation charge succeeds. The subscription returns to active with a fresh invoice and a billing period anchored to the reactivation date. Distinct from subscription.activated (first activation) and payment.recovered (past_due recovery, which keeps the original anchor).

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `currentPeriodStart` (`string`)
- `currentPeriodEnd` (`string`)
- `name` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `invoiceTotal` (`float64`)
- `invoiceCurrency` (`string`)
- `provider` (`string`)

## subscription.canceled

Fired when a subscription is actually terminated. A scheduled cancellation fires it at the end of the billing period; immediate cancellations, full refunds (cancelReason refund), and exhausted dunning retries (cancelReason dunning_exhausted) fire it right away. The status is now canceled and access should be revoked. This event is NOT fired when cancellation is scheduled — that triggers subscription.updated instead. See the cancellation lifecycle below.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `canceledAt` (`string`)
- `cancelReason` (`string`)
- `endDate` (`string`)

## subscription.updated

Fired when subscription details change. The most common trigger is scheduling a cancellation — when a customer cancels, the status stays "active" until the billing period ends, but canceledAt and endDate are set immediately. Use this event to show "your subscription will end on {endDate}" in your UI. Access should NOT be revoked here — wait for subscription.canceled.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `canceledAt` (`string`)
- `cancelReason` (`string`)
- `endDate` (`string`)

## subscription.plan_changed

Fired when a subscription changes from one plan to another, including upgrades, downgrades, and billing interval changes. Access does not change on this event — the subscription stays active.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `previousPlan` (`WebhookPlanRef`)
- `currentPlan` (`WebhookPlanRef`)
- `billingInterval` (`string`)
- `credit` (`float64`)
- `charge` (`float64`)
- `totalCharged` (`float64`)

## subscription.cancellation_scheduled

Fired when a cancellation is scheduled for the end of the billing period. The subscription stays active until effectiveAt — do NOT revoke access here. subscription.updated also fires for backward compatibility.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `canceledAt` (`string`)
- `cancelReason` (`string`)
- `effectiveAt` (`string`)

## subscription.cancellation_revoked

Fired when a scheduled cancellation is reverted before it executes. The subscription continues on its current plan and billing period as if it had never been canceled.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `currentPeriodEnd` (`string`)

## subscription.plan_change_scheduled

Fired when a plan change (downgrade or shorter interval) is scheduled for the end of the billing period. The subscription stays on the current plan until effectiveAt, when subscription.plan_changed fires.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `currentPlan` (`WebhookPlanRef`)
- `scheduledPlan` (`WebhookPlanRef`)
- `billingInterval` (`string`)
- `scheduledBillingInterval` (`string`)
- `effectiveAt` (`string`)

## subscription.plan_change_revoked

Fired when a scheduled plan change is replaced by a different one before it executes. The replacement also fires subscription.plan_change_scheduled with the new target plan.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `currentPlan` (`WebhookPlanRef`)
- `revokedPlan` (`WebhookPlanRef`)
- `billingInterval` (`string`)
- `revokedBillingInterval` (`string`)

## subscription.past_due

Fired when a recurring payment fails on a previously paid subscription and its status becomes past_due. past_due is a grace window, not a cutoff: usage and seats keep working while new purchases are blocked, and dunning retries the charge — use this to notify the customer and recover the payment.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)

## trial.started

Fired when a subscription enters its trial period after checkout. Grant access here — trialing subscriptions have full access until trialEndsAt.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `planId` (`string`)
- `planName` (`string`)
- `trialEndsAt` (`string`)

## trial.converted

Fired when a trialing customer converts to a paid subscription before the trial ends — today this happens when they change plan during the trial, which charges the full new plan price immediately. Trials that simply run out fire trial.expired instead.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `planId` (`string`)
- `planName` (`string`)

## trial.expired

Fired when a trial period runs out and the billing cycle activates the subscription. The first regular invoice is generated right after — this is the natural trial-to-paid transition.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `planId` (`string`)
- `planName` (`string`)
- `trialEndsAt` (`string`)

## trial.will_end

Predictive event fired once, 3 days before a trial ends. Use it to remind the customer that billing starts soon. Emitted by a daily scan with a deterministic idempotency key, so it never fires twice for the same trial end date.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `planId` (`string`)
- `planName` (`string`)
- `trialEndsAt` (`string`)

## trial.checkout_ready

Fired when a trial checkout link is ready to share with the customer. Completing this checkout saves a payment method and starts the trial (trial.started) — the customer is not charged until the trial ends.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `planName` (`string`)
- `trialDays` (`float64`)
- `checkoutUrl` (`string`)

## checkout.ready

Fired when a checkout link for a subscription's first invoice is ready to share with the customer. Commet also emails the link — use this event to deliver it through your own channels.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `invoiceTotal` (`float64`)
- `invoiceCurrency` (`string`)
- `checkoutUrl` (`string`)

## payment.received

Fired every time a payment settles successfully — the first payment and every renewal alike. subscription.activated fires alongside it only on the first one.

- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `invoiceTotal` (`float64`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `paymentTransactionId` (`string`)
- `provider` (`string`)
- `grossAmount` (`float64`)
- `currency` (`string`)
- `orgNetAmount` (`float64`)
- `customerEmail` (`string`)
- `paidAt` (`string`)

## payment.failed

Fired when a recurring charge fails. This event is for recurring charge failures only — card declines during initial checkout do not trigger this event.

- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `provider` (`string`)
- `failureCode` (`string`)
- `failureMessage` (`string`)
- `recoveryUrl` (`string`)

## payment.recovered

Fired when an outstanding invoice that previously failed is successfully paid — automatically on retry or by the customer through the portal. The subscription returns to active at the same time; use this event to close the dunning flow you opened on payment.failed.

- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `invoiceTotal` (`float64`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `provider` (`string`)

## payment.retry_failed

Fired when all dunning retries are exhausted and the subscription is canceled. This is the terminal event of the dunning flow — payment.recovered will not follow. Revoke access when you receive this.

- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `provider` (`string`)
- `reason` (`string`)

## payment.refunded

Fired when a payment is refunded, fully or partially. A full refund of a subscription invoice also cancels the subscription immediately (subscription.canceled fires with reason refund); partial refunds leave the subscription untouched.

- `paymentTransactionId` (`string`)
- `provider` (`string`)
- `paymentLinkId` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `refundAmount` (`float64`)
- `currency` (`string`)

## payment.disputed

Fired when a cardholder opens a dispute (chargeback) against a payment. The disputed amount is frozen from your payout balance while the dispute is open; Commet, as the Merchant of Record, handles the resolution process. payment.dispute_resolved fires with the outcome.

- `paymentTransactionId` (`string`)
- `provider` (`string`)
- `paymentLinkId` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `disputeAmount` (`float64`)
- `currency` (`string`)
- `disputeReason` (`string`)

## payment.dispute_resolved

Fired when a dispute is closed. Carries the same identifiers as payment.disputed plus the outcome: won restores the frozen amount to your balance, lost keeps the chargeback deducted.

- `paymentTransactionId` (`string`)
- `provider` (`string`)
- `paymentLinkId` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `disputeAmount` (`float64`)
- `currency` (`string`)
- `disputeReason` (`string`)
- `outcome` (`string`)

## payment_link.created

Fired when a payment link is created. The link is pending — the customer has not paid yet. Do NOT fulfill here; wait for payment_link.completed.

- `paymentId` (`string`)
- `status` (`string`)
- `amount` (`float64`)
- `currency` (`string`)
- `description` (`string`)
- `customerId` (`string`)

## payment_link.completed

Fired when a payment link is paid. The charge settled and a one-time invoice was generated. Fulfill the purchase on this event.

- `paymentId` (`string`)
- `status` (`string`)
- `amount` (`float64`)
- `currency` (`string`)
- `description` (`string`)
- `customerId` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `paymentTransactionId` (`string`)

## payment_link.failed

Fired when a payment link charge attempt is declined. The link stays open and can be paid again — a failed link is retryable.

- `paymentId` (`string`)
- `status` (`string`)
- `amount` (`float64`)
- `currency` (`string`)
- `description` (`string`)
- `customerId` (`string`)
- `failureCode` (`string`)
- `failureMessage` (`string`)

## payment_link.canceled

Fired when a pending payment link is canceled before being paid. A canceled link can no longer be paid.

- `paymentId` (`string`)
- `status` (`string`)
- `amount` (`float64`)
- `currency` (`string`)
- `description` (`string`)
- `customerId` (`string`)

## invoice.created

Fired when a new invoice is generated for a subscription, typically at the start of a billing period.

- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `invoiceStatus` (`string`)
- `periodStart` (`string`)
- `periodEnd` (`string`)
- `issueDate` (`string`)
- `dueDate` (`string`)
- `currency` (`string`)
- `subtotal` (`float64`)
- `total` (`float64`)
- `customerId` (`string`)
- `subscriptionId` (`string`)

## invoice.voided

Fired when an invoice is voided — nullified before collection, either manually or automatically when its subscription is canceled. Voiding is terminal: a void invoice is never retried or collected.

- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `invoiceStatus` (`string`)
- `periodStart` (`string`)
- `periodEnd` (`string`)
- `issueDate` (`string`)
- `dueDate` (`string`)
- `currency` (`string`)
- `subtotal` (`float64`)
- `total` (`float64`)
- `customerId` (`string`)
- `subscriptionId` (`string`)

## invoice.overdue

Fired once when an outstanding invoice passes its due date without payment. The invoice keeps its outstanding status — overdue is a fact about the due date, not a new status. Use it to start your own dunning flow.

- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `invoiceStatus` (`string`)
- `periodStart` (`string`)
- `periodEnd` (`string`)
- `issueDate` (`string`)
- `dueDate` (`string`)
- `currency` (`string`)
- `subtotal` (`float64`)
- `total` (`float64`)
- `customerId` (`string`)
- `subscriptionId` (`string`)

## invoice.upcoming

Predictive event fired once, 3 days before an active subscription renews. Use it to notify the customer before they are charged. Carries no amount — usage-based charges are only final at renewal, when invoice.created delivers the actual invoice.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `status` (`string`)
- `planId` (`string`)
- `planName` (`string`)
- `billingInterval` (`string`)
- `currentPeriodEnd` (`string`)

## payment_method.attached

Fired when Commet records a payment method for a subscription: after a paid checkout, when a trial starts with a card on file, or when a zero-total checkout completes. The card object carries display metadata only — full numbers never leave the payment provider.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `card` (`WebhookCardInfo`)

## payment_method.updated

Fired when a customer replaces their default payment method through the customer portal. The new method applies to all of the customer's subscriptions. A payment method update is also a strong recovery signal for past-due subscriptions.

- `customerId` (`string`)
- `card` (`WebhookCardInfo`)

## customer.created

Fired when a customer is created, via the API (including batch create), SDK, or dashboard. The payload is the customer resource exactly as GET /customers returns it.

- `id` (`string`)
- `externalId` (`string`)
- `fullName` (`string`)
- `email` (`string`)
- `taxDocument` (`string`)
- `documentType` (`string`)
- `timezone` (`string`)
- `metadata` (`map[string]any`)
- `createdAt` (`string`)
- `updatedAt` (`string`)

## customer.updated

Fired when a customer's details change (email, name, timezone, externalId, or metadata). Carries the same customer resource shape as customer.created with the current values.

- `id` (`string`)
- `externalId` (`string`)
- `fullName` (`string`)
- `email` (`string`)
- `taxDocument` (`string`)
- `documentType` (`string`)
- `timezone` (`string`)
- `metadata` (`map[string]any`)
- `createdAt` (`string`)
- `updatedAt` (`string`)

## customer.state_changed

Aggregate entitlement event answering one question: what can this customer access right now? Fired on every entitlement transition (subscription lifecycle, plan changes, trials, past due, scheduled cancellations) with the customer's CURRENT subscription, plan, features, seats, and credits or balance. Handle this single event to keep access in sync instead of wiring every lifecycle event.

- `customerId` (`string`)
- `trigger` (`string`)
- `status` (`string`)
- `subscriptionId` (`string`)
- `plan` (`WebhookPlanRef`)
- `billingInterval` (`string`)
- `consumptionModel` (`string`)
- `features` (`[]any`)
- `seats` (`[]WebhookSeatSummary`)
- `credits` (`WebhookCreditsBalance`)
- `balance` (`WebhookBalance`)

## plan_grant.created

Fired after a plan grant is durably created. The payload is the grant snapshot at creation.

- `id` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `basePlanId` (`string`)
- `targetPlanId` (`string`)
- `targetPlanReleaseId` (`string`)
- `status` (`string`)
- `duration` (`string`)
- `durationCycles` (`int`)
- `startsAt` (`string`)
- `expiresAt` (`string`)
- `reason` (`string`)
- `source` (`string`)
- `revokedAt` (`string`)
- `createdAt` (`string`)
- `updatedAt` (`string`)
- `events` (`[]WebhookPlanGrantTimelineEvent`)

## plan_grant.updated

Fired after a plan grant duration or deadline is durably changed. The payload is the grant snapshot at that update.

- `id` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `basePlanId` (`string`)
- `targetPlanId` (`string`)
- `targetPlanReleaseId` (`string`)
- `status` (`string`)
- `duration` (`string`)
- `durationCycles` (`int`)
- `startsAt` (`string`)
- `expiresAt` (`string`)
- `reason` (`string`)
- `source` (`string`)
- `revokedAt` (`string`)
- `createdAt` (`string`)
- `updatedAt` (`string`)
- `events` (`[]WebhookPlanGrantTimelineEvent`)

## plan_grant.expired

Fired after a plan grant is durably expired, whether discovered automatically or while replacing an expired grant.

- `id` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `basePlanId` (`string`)
- `targetPlanId` (`string`)
- `targetPlanReleaseId` (`string`)
- `status` (`string`)
- `duration` (`string`)
- `durationCycles` (`int`)
- `startsAt` (`string`)
- `expiresAt` (`string`)
- `reason` (`string`)
- `source` (`string`)
- `revokedAt` (`string`)
- `createdAt` (`string`)
- `updatedAt` (`string`)
- `events` (`[]WebhookPlanGrantTimelineEvent`)

## plan_grant.revoked

Fired after an active plan grant is durably revoked. The payload is the grant snapshot at revocation.

- `id` (`string`)
- `customerId` (`string`)
- `subscriptionId` (`string`)
- `basePlanId` (`string`)
- `targetPlanId` (`string`)
- `targetPlanReleaseId` (`string`)
- `status` (`string`)
- `duration` (`string`)
- `durationCycles` (`int`)
- `startsAt` (`string`)
- `expiresAt` (`string`)
- `reason` (`string`)
- `source` (`string`)
- `revokedAt` (`string`)
- `createdAt` (`string`)
- `updatedAt` (`string`)
- `events` (`[]WebhookPlanGrantTimelineEvent`)

## credits.granted

Fired when non-purchase credits are granted to a subscription: plan-included credits at the start of each billing period, or a manual adjustment from the dashboard. Credit pack purchases fire credits.purchased instead.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `credits` (`float64`)
- `reason` (`string`)

## credits.purchased

Fired when a customer buys a credit pack through the customer portal and the payment succeeds. Purchased credits never expire — unlike plan credits, they survive period resets. Plan-included credit grants fire credits.granted instead.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `creditPackName` (`string`)
- `credits` (`float64`)

## credits.low

Fired when a subscription's remaining credits cross below 10% of the credits granted for the current period. Emitted once per billing period, when the crossing happens.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `remainingCredits` (`float64`)
- `thresholdCredits` (`float64`)
- `periodCredits` (`float64`)

## credits.depleted

Fired when a subscription's credits hit zero. Usage requests that need more credits than remain are rejected from this point. Also fires customer.state_changed with trigger credits_depleted.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `remainingCredits` (`float64`)

## credits.expired

Fired at the period reset when unused plan credits from the previous period are discarded. Plan credits expire at period end; purchased credits never expire and are not affected.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `expiredCredits` (`float64`)

## balance.topped_up

Fired when a customer on a balance plan tops up their prepaid balance through the customer portal and the payment succeeds.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `invoiceId` (`string`)
- `invoiceNumber` (`string`)
- `amount` (`float64`)
- `currency` (`string`)

## balance.low

Fired when a subscription's prepaid balance crosses below 10% of its last refill (period reset, top-up, or manual adjustment). Emitted once per crossing.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `currentBalance` (`float64`)
- `thresholdBalance` (`float64`)
- `currency` (`string`)

## balance.depleted

Fired when a subscription's prepaid balance crosses to zero or below. With block-on-exhaustion plans further usage is rejected; otherwise the balance can go negative. Also fires customer.state_changed with trigger balance_depleted.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `currentBalance` (`float64`)
- `currency` (`string`)

## quota.threshold_reached

Fired when a metered feature's usage crosses 80% of its included quantity for the current period. Emitted once per feature per billing period, when the crossing happens.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `featureCode` (`string`)
- `currentUsage` (`float64`)
- `includedAmount` (`float64`)
- `periodStart` (`string`)

## quota.exceeded

Fired when a metered feature passes its included quantity. With overage enabled it means overage billing began; with overage disabled it means the hard limit was hit and further usage is rejected (this case also fires customer.state_changed with trigger quota_exceeded). Emitted once per feature per billing period.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `featureCode` (`string`)
- `currentUsage` (`float64`)
- `includedAmount` (`float64`)
- `overageEnabled` (`bool`)
- `periodStart` (`string`)

## seats.updated

Fired when a customer's seat count changes for a seats-type feature — via the SDK seats endpoints or the dashboard. Also fires customer.state_changed with trigger seats_updated.

- `customerId` (`string`)
- `subscriptionId` (`string`)
- `featureCode` (`string`)
- `previousSeats` (`float64`)
- `currentSeats` (`float64`)

## seats.limit_reached

Fired when a seat change reaches or passes the included seat limit of the customer's plan. Emitted once per crossing — only when the count moves from below the limit to at or above it.

- `customerId` (`string`)
- `subscriptionId` (`string`)
- `featureCode` (`string`)
- `currentSeats` (`float64`)
- `includedSeats` (`float64`)

## addon.activated

Fired when an add-on is activated on a subscription — via the API or a customer portal purchase. The prorated activation charge, if any, has already succeeded. Also fires customer.state_changed with trigger addon_activated.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `addon` (`WebhookAddonRef`)
- `featureCode` (`string`)
- `proratedPrice` (`float64`)
- `currency` (`string`)

## addon.deactivated

Fired when an active add-on is deactivated from a subscription. Also fires customer.state_changed with trigger addon_deactivated.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `addon` (`WebhookAddonRef`)
- `featureCode` (`string`)

## usage.recorded

Fired for every processed usage event. HIGH VOLUME: this fires once per tracked event, so it is excluded from family select-all in the dashboard — subscribe to it explicitly and make sure your endpoint can absorb your own ingest rate.

- `subscriptionId` (`string`)
- `customerId` (`string`)
- `usageEventId` (`string`)
- `featureCode` (`string`)
- `value` (`float64`)
- `ts` (`string`)

## payout.available

Organization-level event about YOUR money as the merchant. Fired when payment funds the provider was holding become available to pay out to your bank.

- `availableAmount` (`float64`)
- `currency` (`string`)

## payout.created

Fired when a payout of your available balance is requested and the transfer toward your bank is initiated. The lifecycle continues with payout.paid or payout.failed.

- `payoutId` (`string`)
- `amount` (`float64`)
- `fee` (`float64`)
- `netAmount` (`float64`)
- `currency` (`string`)
- `status` (`string`)
- `destinationBank` (`WebhookBankRef`)
- `createdAt` (`string`)

## payout.paid

Fired when the bank settlement of a payout completes — the moment the money actually reaches your bank account, confirmed by the payment provider. Fires exactly once per payout.

- `payoutId` (`string`)
- `amount` (`float64`)
- `fee` (`float64`)
- `netAmount` (`float64`)
- `currency` (`string`)
- `status` (`string`)
- `destinationBank` (`WebhookBankRef`)
- `paidAt` (`string`)

## payout.failed

Fired when the provider reports a payout could not be completed — most commonly a bank rejection (closed account, invalid details). The funds return to your available balance.

- `payoutId` (`string`)
- `amount` (`float64`)
- `fee` (`float64`)
- `netAmount` (`float64`)
- `currency` (`string`)
- `status` (`string`)
- `destinationBank` (`WebhookBankRef`)
- `failedAt` (`string`)
- `failureCode` (`string`)
- `failureMessage` (`string`)
