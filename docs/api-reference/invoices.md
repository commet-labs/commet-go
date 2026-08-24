# Invoices

API version: `2026-07-31`

## GetDownloadURL

`client.Invoices.GetDownloadURL(ctx, ...)`

`POST /invoices/{id}/download-links` · operation `download-invoice`

Generate a signed URL to download the invoice as a PDF. The URL expires after 7 days.

### Parameters

- `ID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`InvoiceDownload`

## Get

`client.Invoices.Get(ctx, ...)`

`GET /invoices/{id}` · operation `get-invoice`

Retrieve a single invoice by its public ID, including line items.

### Parameters

- `ID` (`string`, required)

### Returns

`Invoice`

## Send

`client.Invoices.Send(ctx, ...)`

`POST /invoices/{id}/send` · operation `send-invoice`

Send the invoice to the customer via email.

### Parameters

- `ID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`SentInvoice`

## UpdateStatus

`client.Invoices.UpdateStatus(ctx, ...)`

`PATCH /invoices/{id}/status` · operation `update-invoice-status`

Mark an outstanding invoice as "paid" or "void" and return the updated invoice. Cannot change the status of already paid or voided invoices.

### Parameters

- `ID` (`string`, required)
- `Status` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Invoice`

## List

`client.Invoices.List(ctx, ...)`

`GET /invoices` · operation `list-invoices`

List invoices with cursor-based pagination. Filter by customer, status, or subscription.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)
- `CustomerID` (`string`, optional)
- `Status` (`string`, optional)
- `SubscriptionID` (`string`, optional)

### Returns

`InvoicesListResult`

## CreateAdjustment

`client.Invoices.CreateAdjustment(ctx, ...)`

`POST /invoices` · operation `create-adjustment-invoice`

Create a one-off adjustment invoice and return the created invoice. Use a negative amount for a credit.

### Parameters

- `CustomerID` (`string`, required)
- `Amount` (`int`, required)
- `Description` (`string`, required)
- `Metadata` (`map[string]any`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Invoice`
