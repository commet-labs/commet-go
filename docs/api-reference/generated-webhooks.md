# Webhooks

API version: `2026-07-31`

## Get

`client.Webhooks.Get(ctx, ...)`

`GET /webhooks/{id}` · operation `get-webhook-endpoint`

Retrieve a webhook endpoint by its public ID.

### Parameters

- `ID` (`string`, required)

### Returns

`Webhook`

## Update

`client.Webhooks.Update(ctx, ...)`

`PATCH /webhooks/{id}` · operation `update-webhook-endpoint`

Update a webhook endpoint. Only the provided fields change.

### Parameters

- `ID` (`string`, required)
- `URL` (`string`, optional)
- `Events` (`[]string`, optional)
- `Description` (`string | null`, optional)
- `IsActive` (`bool`, optional)
- `APIVersion` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Webhook`

## Delete

`client.Webhooks.Delete(ctx, ...)`

`DELETE /webhooks/{id}` · operation `delete-webhook-endpoint`

Permanently delete a webhook endpoint.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedObject`

## Test

`client.Webhooks.Test(ctx, ...)`

`POST /webhooks/{id}/test` · operation `test-webhook-endpoint`

Send a test event to a webhook endpoint to verify connectivity.

### Parameters

- `ID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`WebhookTest`

## List

`client.Webhooks.List(ctx, ...)`

`GET /webhooks` · operation `list-webhook-endpoints`

List webhook endpoints with cursor-based pagination.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)

### Returns

`WebhooksListResult`

## Create

`client.Webhooks.Create(ctx, ...)`

`POST /webhooks` · operation `create-webhook-endpoint`

Create a new webhook endpoint. The response includes the signing secret which is only returned once.

### Parameters

- `URL` (`string`, required)
- `Events` (`[]string`, required)
- `Description` (`string`, optional)
- `APIVersion` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CreatedWebhook`
