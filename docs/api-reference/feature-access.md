# Feature Access

API version: `2026-07-31`

## Get

`client.FeatureAccess.Get(ctx, ...)`

`GET /feature-access/{code}` · operation `get-feature-access`

Get one feature's access and current usage for a customer. To evaluate a prospective consumption, use POST /usage/check.

### Parameters

- `Code` (`string`, required)
- `CustomerID` (`string`, required)

### Returns

`FeatureAccess`

## List

`client.FeatureAccess.List(ctx, ...)`

`GET /feature-access` · operation `list-feature-access`

List a customer's feature access and current usage.

### Parameters

- `CustomerID` (`string`, required)

### Returns

`FeatureAccessListResult`
