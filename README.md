# wb-api-client

OpenAPI-first Go client for Wildberries API (v1 scope):
- General
- Product Management
- FBS Orders
- DBW Orders
- DBS Orders

## Architecture

- `specs/wb/`: pinned Swagger/OpenAPI specs.
- `internal/generated/<category>`: generated clients (source of truth).
- `client/`: ergonomic service facade over generated clients.
- `transport/`: retry/rate-limit middleware + unified API errors.
- `internal/testkit/`: HTTP/JSON test helpers.

## Facade Coverage

- Products:
  - Create and update cards (`/content/v2/cards/upload`, `/content/v2/cards/update`, `/content/v2/cards/upload/add`)
  - Get cards list
  - Set prices and size prices (`/api/v2/upload/task`, `/api/v2/upload/task/size`)
  - Track processed price uploads (`/api/v2/history/goods/task`)
  - Get and update inventory (`/api/v3/stocks/{warehouseId}`)
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

By default, generated server mappings are preserved (including multi-host categories like Products).
`With*BaseURL` options explicitly override all operation servers in that category.

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
