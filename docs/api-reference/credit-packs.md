# Credit Packs

API version: `2026-08-27`

## Update

`client.CreditPacks.Update(ctx, ...)`

`PATCH /credit-packs/{id}` · operation `update-credit-pack`

Update a credit pack's name, description, credits, price, active status, or the plans that sell it.

### Parameters

- `ID` (`string`, required)
- `Name` (`string`, optional)
- `Description` (`string`, optional)
- `Credits` (`int`, optional)
- `Price` (`int`, optional)
- `IsActive` (`bool`, optional)
- `PlanIds` (`[]string`, optional) — Plans that sell the credit pack, replacing the current ones. Empty means every plan; omitted keeps them.

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
- `PlanIds` (`[]string`, optional) — Plans that sell the credit pack. Omitted or empty, every plan sells it. Plans created later sell it too.

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CreditPack`
