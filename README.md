# wb-api-client

[![tests](https://img.shields.io/github/actions/workflow/status/benice2me11/wb-api-client/tests.yml?branch=main&label=tests)](https://github.com/benice2me11/wb-api-client/actions/workflows/tests.yml)
[![coverage](https://img.shields.io/codecov/c/github/benice2me11/wb-api-client/main?label=coverage)](https://codecov.io/gh/benice2me11/wb-api-client)

OpenAPI-first Go client for Wildberries API (v1 scope):
- General
- Product Management
- FBS Orders
- DBW Orders
- DBS Orders
- Reports
- Analytics
- Orders FBW

## Architecture

- `specs/wb/`: pinned Swagger/OpenAPI specs.
- `internal/generated/<category>`: generated clients (source of truth).
- `client/`: ergonomic service facade over generated clients.
- `transport/`: retry/rate-limit middleware + unified API errors.
- `internal/testkit/`: HTTP/JSON test helpers.

## Facade Coverage

- Products:
  - Create and update cards (`/content/v2/cards/upload`, `/content/v2/cards/update`, `/content/v2/cards/upload/add`)
  - Move cards to trash (`/content/v2/cards/delete/trash`)
  - Get cards list
  - Get products with prices (`/api/v2/list/goods/filter`, GET and POST)
  - Get size prices by article (`/api/v2/list/goods/size/nm`)
  - Set prices and size prices (`/api/v2/upload/task`, `/api/v2/upload/task/size`)
  - Track processed price uploads (`/api/v2/history/goods/task`)
  - Get, update, reset-to-zero and delete inventory (`/api/v3/stocks/{warehouseId}`)
  - Get offices for warehouse creation (`/api/v3/offices`)
  - List, create, update and delete warehouses (`/api/v3/warehouses`)
  - Get and update DBW warehouse contacts (`/api/v3/dbw/warehouses/{warehouseId}/contacts`)
- FBS Orders:
  - Get new orders and orders list
  - Get order statuses and status history
  - Cancel order
- DBW Orders:
  - Get new orders and completed orders list
  - Get order statuses
  - Confirm, assemble and cancel orders
- DBS Orders:
  - Get new orders and completed orders list
  - Get order statuses (`status/info`)
  - Confirm, deliver, receive, reject and cancel orders
- Reports:
  - Create, check and download warehouse remains reports
  - Supplier stocks, orders and sales reports
  - Excise, acceptance and paid storage report tasks
  - Antifraud, banned products, brand share, goods labeling, goods return and region sale reports
  - Deductions, measurement penalties and warehouse measurements reports
- Analytics:
  - Sales funnel products, product history and grouped history
  - Search report summary, groups, details, product search texts and product orders
  - Stocks reports for WB warehouses, product groups, products, sizes and offices
  - Item rating
  - CSV report list, create, retry and file download, including stock history CSV requests
- Orders FBW:
  - Acceptance options
  - Warehouses and transit tariffs
  - Supplies list, supply details, supply goods and supply package

By default, generated server mappings are preserved (including multi-host categories like Products).
`With*BaseURL` options explicitly override all operation servers in that category.

## Public Types

External projects should import only `github.com/benice2me11/wb-api-client/client`.
The facade exposes request/response aliases in `client` (for example, `client.ContentV2GetCardsListPostRequest`), so there is no need to import `internal/generated/*`.

## Usage Example

A simple example on how to use this library:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	wb "github.com/benice2me11/wb-api-client/client"
)

func main() {
	// Create a client with your Wildberries API token.
	// Authorization header is "Bearer <token>" by default.
	opts := []wb.Option{
		wb.WithToken("api-token"),
	}
	c := wb.NewClient(opts...)

	// Send request through the typed facade.
	warehouses, resp, err := c.Products().Warehouses(context.Background())
	if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
		status := "no response"
		if resp != nil {
			status = resp.Status
		}
		log.Fatalf("error when getting warehouses: %v (status: %s)", err, status)
	}

	// Do some stuff.
	for _, warehouse := range warehouses {
		fmt.Printf("Warehouse %d: %s\n", warehouse.GetId(), warehouse.GetName())
	}
}
```

## Regenerate Client Code

Requirements:
- Docker

Run:

```bash
make generate
```

The generator is pinned in `scripts/generate.sh` (`openapitools/openapi-generator-cli:v7.13.0`).

To ensure generated files are up-to-date:

```bash
make verify-generated
```
