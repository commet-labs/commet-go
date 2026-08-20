# Payments

API version: `2026-07-31`

## Cancel

`client.Payments.Cancel(ctx, ...)`

`POST /payments/{id}/cancel` · operation `cancel-payment`

Cancel a pending payment link so it can no longer be paid. Only a link that has not been paid or started processing can be canceled; canceling an already canceled link is a no-op. Charges cannot be canceled.

### Parameters

- `ID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Payment`

## Get

`client.Payments.Get(ctx, ...)`

`GET /payments/{id}` · operation `get-payment`

Retrieve a payment by its public ID.

### Parameters

- `ID` (`string`, required)

### Returns

`Payment`

## Charge

`client.Payments.Charge(ctx, ...)`

`POST /payments/charge` · operation `charge-payment`

Charge a customer's vaulted payment method off-session. Calculates tax, generates an invoice, and sends a receipt. Requires the customer to have a subscription in active, trialing, or past_due state.

### Parameters

- `CustomerID` (`string`, required)
- `Amount` (`int`, required)
- `Currency` (`string`, required)
- `Description` (`string`, required)
- `Metadata` (`map[string]string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Payment`

## List

`client.Payments.List(ctx, ...)`

`GET /payments` · operation `list-payments`

List payments with cursor-based pagination. Filter by customer.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)
- `CustomerID` (`string`, optional)

### Returns

`PaymentsListResult`

## Create

`client.Payments.Create(ctx, ...)`

`POST /payments` · operation `create-payment`

Create a hosted payment link. Returns a url the customer opens to pay with any card. Calculates tax, generates an invoice, and vaults the payment method on confirmation. No subscription or plan required.

### Parameters

- `Amount` (`int`, required)
- `Currency` (`string`, required)
- `CustomerID` (`string`, optional)
- `Description` (`string`, required)
- `SuccessURL` (`string`, optional)
- `Metadata` (`map[string]string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Payment`
