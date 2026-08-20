# Addons

API version: `2026-07-31`

## ListActive

`client.Addons.ListActive(ctx, ...)`

`GET /active-addons` · operation `list-active-addons`

List all active add-ons for a customer's subscription.

### Parameters

- `CustomerID` (`string`, required)

### Returns

`AddonsListActiveResult`

## Get

`client.Addons.Get(ctx, ...)`

`GET /addons/{id}` · operation `get-addon`

Retrieve an add-on by its public ID or slug.

### Parameters

- `ID` (`string`, required)

### Returns

`Addon`

## Update

`client.Addons.Update(ctx, ...)`

`PATCH /addons/{id}` · operation `update-addon`

Update an add-on's name, description, or pricing.

### Parameters

- `ID` (`string`, required)
- `Name` (`string`, optional)
- `Description` (`string`, optional)
- `BasePrice` (`int`, optional)
- `IncludedUnits` (`int`, optional)
- `OverageRate` (`int`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Addon`

## Delete

`client.Addons.Delete(ctx, ...)`

`DELETE /addons/{id}` · operation `delete-addon`

Soft-delete an add-on. Fails if the add-on has active subscriptions.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedObject`

## List

`client.Addons.List(ctx, ...)`

`GET /addons` · operation `list-addons`

List all add-ons with cursor-based pagination.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)

### Returns

`AddonsListResult`

## Create

`client.Addons.Create(ctx, ...)`

`POST /addons` · operation `create-addon`

Create a new add-on linked to a feature. Each feature can only be assigned to one add-on.

### Parameters

- `Name` (`string`, required)
- `Description` (`string`, optional)
- `BasePrice` (`int`, required)
- `FeatureID` (`string`, required)
- `ConsumptionModel` (`string`, required)
- `IncludedUnits` (`int`, optional)
- `OverageRate` (`int`, optional)
- `CreditCost` (`int`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Addon`
