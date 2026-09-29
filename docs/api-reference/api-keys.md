# Api Keys

API version: `2026-08-27`

## Delete

`client.APIKeys.Delete(ctx, ...)`

`DELETE /api-keys/{id}` · operation `delete-api-key`

Permanently revoke and delete an API key.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedObject`

## List

`client.APIKeys.List(ctx, ...)`

`GET /api-keys` · operation `list-api-keys`

List API keys with cursor-based pagination. Keys are returned without the full secret.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)

### Returns

`ApiKeysListResult`

## Create

`client.APIKeys.Create(ctx, ...)`

`POST /api-keys` · operation `create-api-key`

Create a full-access or restricted API key. Provide permissions to restrict access; the full key is returned only once. A restricted key with api_key: write may only create restricted keys with the same or fewer permissions, and they expire no later than the key that creates them.

### Parameters

- `Name` (`string`, required)
- `ExpiresInDays` (`int`, optional)
- `Permissions` (`CreateApiKeyParamsPermissions`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`CreatedApiKey`
