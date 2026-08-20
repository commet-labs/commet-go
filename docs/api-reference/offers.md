# Offers

API version: `2026-07-31`

## Get

`client.Offers.Get(ctx, ...)`

`GET /offers/{id}` · operation `get-offer`

Retrieve reusable offer terms by public ID.

### Parameters

- `ID` (`string`, required)

### Returns

`Offer`

## Update

`client.Offers.Update(ctx, ...)`

`PATCH /offers/{id}` · operation `update-offer`

Replace reusable offer terms. Existing applications keep their immutable accepted terms.

### Parameters

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Phases` (`[]UpdateOfferParamsPhasesItem`, required)
- `Metadata` (`map[string]any`, optional)
- `StartsAt` (`string | null`, optional)
- `EndsAt` (`string | null`, optional)
- `Active` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Offer`

## Delete

`client.Offers.Delete(ctx, ...)`

`DELETE /offers/{id}` · operation `delete-offer`

Soft-delete an Offer. Existing applications and their accepted terms remain available for billing and audit.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedOffer`

## List

`client.Offers.List(ctx, ...)`

`GET /offers` · operation `list-offers`

List reusable offer terms. Offers are independent from plans, prices, eligibility, and distribution channels.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)
- `Active` (`bool`, optional)

### Returns

`OffersListResult`

## Create

`client.Offers.Create(ctx, ...)`

`POST /offers` · operation `create-offer`

Create reusable offer terms without assigning a plan, price, eligibility rule, or distribution channel.

### Parameters

- `Name` (`string`, required)
- `Phases` (`[]CreateOfferParamsPhasesItem`, required)
- `Metadata` (`map[string]any`, optional)
- `StartsAt` (`string | null`, optional)
- `EndsAt` (`string | null`, optional)
- `Active` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Offer`
