# Markets

API version: `2026-07-31`

## Get

`client.Markets.Get(ctx, ...)`

`GET /markets/{id}` · operation `get-market`

Get one reusable market.

### Parameters

- `ID` (`string`, required)

### Returns

`Market`

## Update

`client.Markets.Update(ctx, ...)`

`PATCH /markets/{id}` · operation `update-market`

Replace the name, countries, and metadata of a market.

### Parameters

- `ID` (`string`, required)
- `Name` (`string`, required)
- `CountryCodes` (`[]string`, required)
- `Metadata` (`map[string]any`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Market`

## Delete

`client.Markets.Delete(ctx, ...)`

`DELETE /markets/{id}` · operation `delete-market`

Delete an unused market. Markets referenced by prices or subscriptions cannot be deleted.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedObject`

## List

`client.Markets.List(ctx, ...)`

`GET /markets` · operation `list-markets`

List reusable country groups that resolve market-specific prices independently from currency.

### Returns

`MarketsListResult`

## Create

`client.Markets.Create(ctx, ...)`

`POST /markets` · operation `create-market`

Create a reusable market without attaching it to a plan or price. Countries can belong to only one active market.

### Parameters

- `Name` (`string`, required)
- `CountryCodes` (`[]string`, required)
- `Metadata` (`map[string]any`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Market`
