# Promo Codes

API version: `2026-07-31`

## Get

`client.PromoCodes.Get(ctx, ...)`

`GET /promo-codes/{id}` · operation `get-promo-code`

Retrieve a promo code by its public ID.

### Parameters

- `ID` (`string`, required)

### Returns

`PromoCode`

## Update

`client.PromoCodes.Update(ctx, ...)`

`PATCH /promo-codes/{id}` · operation `update-promo-code`

Update a promo code's billing interval, redemption limits, expiration, active status, or plan restrictions.

### Parameters

- `ID` (`string`, required)
- `BillingInterval` (`string | null`, optional)
- `MaxRedemptions` (`int | null`, optional)
- `ExpiresAt` (`string | null`, optional)
- `Active` (`bool`, optional)
- `PlanIds` (`[]string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PromoCode`

## List

`client.PromoCodes.List(ctx, ...)`

`GET /promo-codes` · operation `list-promo-codes`

List promo codes with cursor-based pagination.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)

### Returns

`PromoCodesListResult`

## Create

`client.PromoCodes.Create(ctx, ...)`

`POST /promo-codes` · operation `create-promo-code`

Create a distribution code for an existing Offer. The referenced Offer owns the benefit and duration; the promo code owns redemption restrictions.

### Parameters

- `Code` (`string`, required)
- `OfferID` (`string`, required)
- `BillingInterval` (`string | null`, optional)
- `MaxRedemptions` (`int`, optional)
- `ExpiresAt` (`string`, optional)
- `PlanIds` (`[]string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PromoCode`
