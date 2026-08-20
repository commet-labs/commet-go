# Portal

API version: `2026-07-31`

## GetURL

`client.Portal.GetURL(ctx, ...)`

`POST /portal/sessions` · operation `request-portal-access`

Generate a customer portal URL. Exactly one identifier (email or customerId) is required.

### Parameters

- `Email` (`string`, optional)
- `ReturnURL` (`string`, optional)
- `CustomerID` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PortalAccess`
