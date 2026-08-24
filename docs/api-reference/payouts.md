# Payouts

API version: `2026-07-31`

## AddBankAccount

`client.Payouts.AddBankAccount(ctx, ...)`

`POST /payouts/bank-accounts` · operation `add-payout-bank-account`

Add an additional destination bank account to the organization's existing payout account. Country and currency are resolved from the organization. The full account number is never returned — only `last4`.

### Parameters

- `AccountNumber` (`string`, required)
- `AccountHolderName` (`string`, required)
- `RoutingNumber` (`string`, optional)
- `AccountType` (`string`, optional)
- `SetDefault` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PayoutBankAccount`

## Request

`client.Payouts.Request(ctx, ...)`

`POST /payouts` · operation `request-payout`

Withdraw available balance to the organization's verified payout account. `amount` is in cents (USD, minimum 1000 = $10). The payout is created in `pending` and settles to `paid` asynchronously as provider webhooks arrive.

### Parameters

- `Amount` (`int`, required)
- `Description` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Payout`

## CompleteVerification

`client.Payouts.CompleteVerification(ctx, ...)`

`POST /payouts/verification` · operation `complete-payout-verification`

Deprecated. Complete business and identity verification in the Commet dashboard. This endpoint no longer accepts or processes KYC data.

Deprecated.

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`void`
