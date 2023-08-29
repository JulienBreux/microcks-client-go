## 🔌 Microcks GO client

[![GoDoc](https://godoc.org/github.com/JulienBreux/microcks-client-go?status.svg)](http://godoc.org/github.com/JulienBreux/microcks-client-go)
[![GitHub tag](https://img.shields.io/github/tag/JulienBreux/microcks-client-go.svg)](Tag)
[![Go version](https://img.shields.io/badge/go-v1.21-blue)](https://golang.org/dl/#stable)
[![License](https://img.shields.io/github/license/JulienBreux/microcks-client-go.svg)](https://github.com/JulienBreux/microcks-client-go/blob/main/LICENSE)
[![Releases](https://img.shields.io/github/downloads/JulienBreux/microcks-client-go/total.svg)](https://github.com/JulienBreux/microcks-client-go/releases)

---

### How to get it

To get the latest version, use go1.21+ and fetch using the `go get` command. For example:

```
go get github.com/JulienBreux/microcks-client-go@latest
```

To get a specific version, use go1.21+ and fetch the desired version using the `go get` command. For example:

```
go get github.com/JulienBreux/microcks-client-go@v0.1.0
```

### How to use it

```go
package main

import (
    "github.com/JulienBreux/microcks-client-go"
)

const server = "http://localhost:8080/api"

fun main() {
	// Create client
	client, err := api.NewClientWithResponses(server)
	if err != nil {
		fatal(err)
	}

	// Request the Keycloak configuration
	ctx := context.Background()
	resp, err := c.actions.GetKeycloakConfigWithResponse(ctx)
	if err != nil {
		fatal(err)
	}

	// Print the response
	fmt.Printf("%+v\n\n", resp.JSON200)
}
```

### License

The MIT License (MIT) - see [`LICENSE.md`](https://github.com/JulienBreux/microcks-client-go/blob/main/LICENSE.md) for more details
