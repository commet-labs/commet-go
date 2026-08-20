# Api Keys

API version: `2026-07-31`

## Delete

`client.ApiKeys.Delete(ctx, ...)`

`DELETE /api-keys/{id}` · operation `delete-api-key`

Permanently revoke and delete an API key.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedObject`

## List

`client.ApiKeys.List(ctx, ...)`

`GET /api-keys` · operation `list-api-keys`

List API keys with cursor-based pagination. Keys are returned without the full secret.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)

### Returns

`ApiKeysListResult`

## Create

`client.ApiKeys.Create(ctx, ...)`

`POST /api-keys` · operation `create-api-key`

Create a new API key. The full key is only returned once in the response.

### Parameters

- `Name` (`string`, required)
- `ExpiresInDays` (`int`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CreatedApiKey`
