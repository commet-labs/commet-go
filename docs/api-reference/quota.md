# Quota

API version: `2026-07-31`

## GetAll

`client.Quota.GetAll(ctx, ...)`

`GET /usage/quota/all` · operation `get-all-quota-allowances`

Get all quota allowances for a customer across every quota feature in their plan.

### Parameters

- `CustomerID` (`string`, required)

### Returns

`QuotaGetAllResult`

## Remove

`client.Quota.Remove(ctx, ...)`

`POST /usage/quota/remove` · operation `remove-quota`

Remove from a customer's quota allowance for a feature. Defaults to 1 if count is omitted. Returns 400 insufficient_balance if the balance would go negative.

### Parameters

- `FeatureCode` (`string`, required)
- `Count` (`int`, optional)
- `CustomerID` (`string`, optional)
- `ExternalID` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`UsageQuotaEvent`

## Get

`client.Quota.Get(ctx, ...)`

`GET /usage/quota` · operation `get-quota-allowance`

Get the current quota allowance (used vs included) for a specific feature.

### Parameters

- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)

### Returns

`UsageQuota`

## Add

`client.Quota.Add(ctx, ...)`

`POST /usage/quota` · operation `add-quota`

Add to a customer's quota allowance for a feature. Defaults to 1 if count is omitted.

### Parameters

- `FeatureCode` (`string`, required)
- `Count` (`int`, optional)
- `CustomerID` (`string`, optional)
- `ExternalID` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`UsageQuotaEvent`

## Set

`client.Quota.Set(ctx, ...)`

`PUT /usage/quota` · operation `set-quota`

Set a customer's quota allowance for a feature to an exact value.

### Parameters

- `FeatureCode` (`string`, required)
- `Count` (`int`, required)
- `CustomerID` (`string`, optional)
- `ExternalID` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`UsageQuotaEvent`
