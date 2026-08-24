# Errors and request IDs

```go
var apiError *commet.CommetError
if errors.As(err, &apiError) {
    fmt.Println(apiError.Code)
    fmt.Println(apiError.RequestID)
    fmt.Println(apiError.DocURL)
}
```

API errors expose type, code, message, status, parameter, details, the exact server request ID, and a versioned documentation URL. The installed error reference describes retry behavior. A request ID is absent when Platform did not return one and is never fabricated locally.

Preserve the same idempotency key when retrying an allowed write.
