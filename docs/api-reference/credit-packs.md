# Credit Packs

API version: `2026-07-31`

## Update

`client.CreditPacks.Update(ctx, ...)`

`PATCH /credit-packs/{id}` · operation `update-credit-pack`

Update a credit pack's name, description, credits, price, or active status.

### Parameters

- `ID` (`string`, required)
- `Name` (`string`, optional)
- `Description` (`string`, optional)
- `Credits` (`int`, optional)
- `Price` (`int`, optional)
- `IsActive` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CreditPack`

## Delete

`client.CreditPacks.Delete(ctx, ...)`

`DELETE /credit-packs/{id}` · operation `delete-credit-pack`

Soft-delete a credit pack.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedObject`

## List

`client.CreditPacks.List(ctx, ...)`

`GET /credit-packs` · operation `list-credit-packs`

List all active credit packs.

### Returns

`CreditPacksListResult`

## Create

`client.CreditPacks.Create(ctx, ...)`

`POST /credit-packs` · operation `create-credit-pack`

Create a new credit pack.

### Parameters

- `Name` (`string`, required)
- `Description` (`string`, optional)
- `Credits` (`int`, required)
- `Price` (`int`, required)
- `IsActive` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CreditPack`
