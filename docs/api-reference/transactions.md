# Transactions

API version: `2026-07-31`

## Refund

`client.Transactions.Refund(ctx, ...)`

`POST /transactions/{id}/refund` · operation `refund-transaction`

Issue a full refund and return the provider-neutral refund resource with its actual status.

### Parameters

- `ID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Refund`

## Retry

`client.Transactions.Retry(ctx, ...)`

`POST /transactions/{id}/retry` · operation `retry-transaction`

Retry a failed subscription renewal and return an honest retry result. The original failed transaction remains immutable.

### Parameters

- `ID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`TransactionRetry`

## Get

`client.Transactions.Get(ctx, ...)`

`GET /transactions/{id}` · operation `get-transaction`

Retrieve a single payment transaction by its public ID, including provider details.

### Parameters

- `ID` (`string`, required)

### Returns

`Transaction`

## List

`client.Transactions.List(ctx, ...)`

`GET /transactions` · operation `list-transactions`

List payment transactions with cursor-based pagination. Filter by status or customer email.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)
- `Status` (`TransactionStatus`, optional)
- `CustomerEmail` (`string`, optional)

### Returns

`TransactionsListResult`
