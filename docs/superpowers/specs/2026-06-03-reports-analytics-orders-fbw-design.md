# Reports, Analytics, Orders FBW Facade Design

## Goal

Generate the full Wildberries Reports, Analytics, and Orders FBW OpenAPI sections and expose ergonomic public facade services for all useful operations in those sections.

## Source Specs

The official dev.wildberries.ru YAML endpoints are protected by browser anti-bot checks in shell access. Use the pinned specs from the `eslazarev/wildberries-sdk` mirror, which tracks the same Wildberries OpenAPI YAML set and includes:

- `specs/12-reports.yaml`
- `specs/11-analytics.yaml`
- `specs/07-orders-fbw.yaml`

Store them locally as:

- `specs/wb/06-reports.yaml`
- `specs/wb/07-analytics.yaml`
- `specs/wb/08-orders-fbw.yaml`

## Architecture

Generated code lives under `internal/generated/reports`, `internal/generated/analytics`, and `internal/generated/ordersfbw`, following the existing category pattern. Public code lives in `client/` with three new services:

- `ReportsService`
- `AnalyticsService`
- `OrdersFBWService`

`Client` owns and exposes these services through `Reports()`, `Analytics()`, and `OrdersFBW()`. Each category gets its own default base URL and `With*BaseURL` option so tests and callers can override generated operation servers consistently.

## Facade Scope

Expose every operation in each new generated category with stable, Go-friendly method names. Use query structs for GET endpoints with several query parameters and direct generated request aliases for JSON bodies. File downloads keep the generated return shape, expected to be `*os.File` for ZIP downloads.

Reports covers main reports, warehouse remains, excise, retentions, acceptance, paid storage, region sale, brand share, hidden products, and goods return endpoints.

Analytics covers sales funnel, search report, inventory report, item rating, and Seller Analytics CSV endpoints.

Orders FBW covers acceptance options, warehouses, transit tariffs, supplies list, supply details, goods, and package endpoints.

## Testing

Add httptest-based facade tests for each new service covering:

- path, method, auth header
- query and request body wiring
- expected success status codes
- wrapped API errors and rate-limit errors

Add compile tests for public type aliases and update README facade coverage.
