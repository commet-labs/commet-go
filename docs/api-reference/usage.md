# Usage

API version: `2026-07-31`

## Check

`client.Usage.Check(ctx, ...)`

`POST /usage/check` · operation `check-usage-availability`

Check if a customer can consume a feature before actual consumption. Returns availability and cost estimates based on the plan's consumption model.

### Parameters

- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)
- `Quantity` (`int`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`UsageCheck`

## Track

`client.Usage.Track(ctx, ...)`

`POST /usage/events` · operation `track-usage`

Track a usage event for a metered feature. Deducts from balance/credits if applicable.

### Parameters

- `FeatureCode` (`string`, required)
- `CustomerID` (`string`, required)
- `EventID` (`string`, optional)
- `Timestamp` (`string`, optional)
- `Properties` (`[]TrackUsageParamsPropertiesItem`, optional)
- `Model` (`string`, optional)
- `InputTokens` (`int`, optional)
- `OutputTokens` (`int`, optional)
- `Value` (`float64`, optional)
- `CacheReadTokens` (`int`, optional)
- `CacheWriteTokens` (`int`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`UsageEvent`

## Set

`client.Usage.Set(ctx, ...)`

`PUT /usage` · operation `set-usage`

Set a metered feature's usage to an exact value for the current period. Use the Idempotency-Key header to make retries safe.

### Parameters

- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)
- `Value` (`int`, required)
- `Reason` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`UsageAdjustment`
