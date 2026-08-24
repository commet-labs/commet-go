# Seats

API version: `2026-07-31`

## GetBalance

`client.Seats.GetBalance(ctx, ...)`

`GET /seats/balance` · operation `get-seat-balance`

Get current balance for a specific seat type.

### Parameters

- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)

### Returns

`SeatBalance`

## GetAllBalances

`client.Seats.GetAllBalances(ctx, ...)`

`GET /seats/balances` · operation `get-all-seat-balances`

Get the current balance for all seat types in a customer's subscription.

### Parameters

- `CustomerID` (`string`, required)

### Returns

`SeatBalanceCollection`

## SetAll

`client.Seats.SetAll(ctx, ...)`

`PUT /seats/bulk` · operation `bulk-set-seats`

Set all seat types at once.

### Parameters

- `CustomerID` (`string`, required)
- `Seats` (`map[string]int`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`SeatsSetAllResult`

## Remove

`client.Seats.Remove(ctx, ...)`

`POST /seats/remove` · operation `remove-seats`

Remove seats from a customer's subscription. Takes effect at the end of the billing period.

### Parameters

- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)
- `Count` (`int`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`SeatEvent`

## Add

`client.Seats.Add(ctx, ...)`

`POST /seats` · operation `add-seats`

Add seats to a customer's subscription. Prorates charges for the current billing period.

### Parameters

- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)
- `Count` (`int`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`SeatEvent`

## Set

`client.Seats.Set(ctx, ...)`

`PUT /seats` · operation `set-seats`

Set seats to an exact count.

### Parameters

- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)
- `Count` (`int`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`SeatEvent`
