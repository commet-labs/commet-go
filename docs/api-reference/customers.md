# Customers

API version: `2026-07-31`

## RevokeCredit

`client.Customers.RevokeCredit(ctx, ...)`

`POST /customers/{id}/credits/{creditId}/revoke` · operation `revoke-customer-credit`

Revoke the unallocated remainder of a customer credit grant. Applied invoice history is unchanged.

### Parameters

- `ID` (`string`, required)
- `CreditID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CustomerCreditRevocation`

## ListCredits

`client.Customers.ListCredits(ctx, ...)`

`GET /customers/{id}/credits` · operation `list-customer-credits`

List currency-specific invoice credit grants and their remaining balances for a customer.

### Parameters

- `ID` (`string`, required)

### Returns

`CustomersListCreditsResult`

## CreateCredit

`client.Customers.CreateCredit(ctx, ...)`

`POST /customers/{id}/credits` · operation `create-customer-credit`

Grant monetary credit in one currency. Credit is applied FIFO before tax to eligible recurring invoices.

### Parameters

- `ID` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's smallest unit.
- `Currency` (`string`, required)
- `Reason` (`string`, required)
- `ExpiresAt` (`string | null`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CustomerCredit`

## RevokePlanGrant

`client.Customers.RevokePlanGrant(ctx, ...)`

`POST /customers/{id}/plan-grants/{grantId}/revoke` · operation `revoke-plan-grant`

End expanded access immediately and restore the base plan's limits. The subscription, billing cycle, invoices, and payment state remain unchanged.

### Parameters

- `ID` (`string`, required)
- `GrantID` (`string`, required)
- `Reason` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanGrant`

## UpdatePlanGrant

`client.Customers.UpdatePlanGrant(ctx, ...)`

`PATCH /customers/{id}/plan-grants/{grantId}` · operation `update-plan-grant`

Keep the overlay for a number of the subscription's existing billing cycles, set an exact deadline, or leave it active until revoked. The billing anchor is never reset.

### Parameters

- `ID` (`string`, required)
- `GrantID` (`string`, required)
- `Reason` (`string`, required)
- `Duration` (`string`, required)
- `DurationCycles` (`int`, optional)
- `ExpiresAt` (`string`, optional)

### Valid parameter combinations

- `Reason` + `Duration` + `DurationCycles`
- `Reason` + `Duration` + `ExpiresAt`
- `Reason` + `Duration`

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanGrant`

## ListPlanGrants

`client.Customers.ListPlanGrants(ctx, ...)`

`GET /customers/{id}/plan-grants` · operation `list-plan-grants`

List the independent audit timeline for paid-plan access granted without checkout or payment credentials.

### Parameters

- `ID` (`string`, required)

### Returns

`CustomersListPlanGrantsResult`

## CreatePlanGrant

`client.Customers.CreatePlanGrant(ctx, ...)`

`POST /customers/{id}/plan-grants` · operation `create-plan-grant`

Temporarily expand an active subscription's feature access using a higher plan in the same plan group. Billing, prices, periods, invoices, and the base subscription remain unchanged.

### Parameters

- `ID` (`string`, required)
- `SubscriptionID` (`string`, required)
- `PlanID` (`string`, required)
- `Reason` (`string`, required)
- `Duration` (`string`, required)
- `DurationCycles` (`int`, optional)
- `ExpiresAt` (`string`, optional)

### Valid parameter combinations

- `SubscriptionID` + `PlanID` + `Reason` + `Duration` + `DurationCycles`
- `SubscriptionID` + `PlanID` + `Reason` + `Duration` + `ExpiresAt`
- `SubscriptionID` + `PlanID` + `Reason` + `Duration`

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanGrant`

## Get

`client.Customers.Get(ctx, ...)`

`GET /customers/{id}` · operation `get-customer`

Retrieve a customer by their public ID, including subscription status and metadata.

### Parameters

- `ID` (`string`, required)

### Returns

`Customer`

## Update

`client.Customers.Update(ctx, ...)`

`PATCH /customers/{id}` · operation `update-customer`

Update a customer's name, external ID, or metadata.

### Parameters

- `ID` (`string`, required)
- `Email` (`string`, optional)
- `FullName` (`string`, optional)
- `TaxDocument` (`string`, optional)
- `ExternalID` (`string`, optional)
- `Timezone` (`Timezone`, optional)
- `Metadata` (`map[string]any`, optional)
- `Address` (`UpdateCustomerParamsAddress`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Customer`

## CreateBatch

`client.Customers.CreateBatch(ctx, ...)`

`POST /customers/batch` · operation `batch-create-customers`

Create up to 100 customers in a single request.

### Parameters

- `Customers` (`[]BatchCreateCustomersParamsCustomersItem`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CustomerBatch`

## List

`client.Customers.List(ctx, ...)`

`GET /customers` · operation `list-customers`

List customers with cursor-based pagination.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)
- `ExternalID` (`string`, optional)

### Returns

`CustomersListResult`

## Create

`client.Customers.Create(ctx, ...)`

`POST /customers` · operation `create-customer`

Create a new customer. Idempotent when customerId is provided.

### Parameters

- `ID` (`string`, optional)
- `ExternalID` (`string`, optional)
- `FullName` (`string`, optional)
- `TaxDocument` (`string`, optional)
- `Address` (`CreateCustomerParamsAddress`, optional)
- `AddressID` (`string`, optional)
- `Email` (`string`, required)
- `Timezone` (`Timezone`, optional)
- `Metadata` (`map[string]any`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Customer`
