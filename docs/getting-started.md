# Getting started

Install the SDK:

```bash
go get github.com/commet-labs/commet-go/v9
```

Create one server-side client. Never expose an API key to browser code.

```go
import commet "github.com/commet-labs/commet-go/v9"

client, err := commet.New("ck_xxx")
```

Every resource and method in this release is generated from the versioned OpenAPI contract. Use the installed API reference instead of relying on remembered method names.
