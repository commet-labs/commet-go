# Features

API version: `2026-08-27`

## Get

`client.Features.Get(ctx, ...)`

`GET /features/{code}` · operation `get-feature`

Get a single feature definition by code from the organization's feature catalog.

### Parameters

- `Code` (`string`, required)

### Returns

`Feature`

## Update

`client.Features.Update(ctx, ...)`

`PATCH /features/{code}` · operation `update-feature`

Update a feature's name, description, or unit name. At least one field must be provided.

### Parameters

- `Code` (`string`, required)
- `Name` (`string`, optional)
- `Description` (`string | null`, optional)
- `UnitName` (`string | null`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Feature`

## Delete

`client.Features.Delete(ctx, ...)`

`DELETE /features/{code}` · operation `delete-feature`

Delete a feature. Fails while a plan's current version, a running experiment or a draft includes it, or while it has an active add-on. Subscribers whose plan version includes it keep it, and its code can be used by a new feature right away.

### Parameters

- `Code` (`string`, required)

### Returns

`DeletedObject`

## List

`client.Features.List(ctx, ...)`

`GET /features` · operation `list-features`

List every feature defined in the organization. This is the organization's feature catalog (definitions), not a customer's feature access.

### Returns

`FeaturesListResult`

## Create

`client.Features.Create(ctx, ...)`

`POST /features` · operation `create-feature`

Create a new feature. Code must be lowercase alphanumeric with underscores.

### Parameters

- `Name` (`string`, required)
- `Code` (`string`, required)
- `Type` (`string`, required)
- `Description` (`string`, optional)
- `UnitName` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Feature`
