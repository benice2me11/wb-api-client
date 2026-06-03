# Reports Analytics Orders FBW Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Generate Reports, Analytics, and Orders FBW clients and expose public facade coverage with tests.

**Architecture:** Add three generated categories under `internal/generated`, wire them into `client.Client`, and expose handwritten service facades that wrap generated errors through `transport.WrapResponseError`. Keep generated code untouched after generation.

**Tech Stack:** Go, OpenAPI Generator Docker image `openapitools/openapi-generator-cli:v7.13.0`, httptest, existing `transport` middleware.

---

### Task 1: Specs And Generation

**Files:**
- Create: `specs/wb/06-reports.yaml`
- Create: `specs/wb/07-analytics.yaml`
- Create: `specs/wb/08-orders-fbw.yaml`
- Modify: `scripts/generate.sh`
- Create: `internal/generated/reports/*`
- Create: `internal/generated/analytics/*`
- Create: `internal/generated/ordersfbw/*`

- [x] Copy the three pinned specs from `/tmp/wb-sdk-eslazarev/specs`.
- [x] Add `run_generate` calls for reports, analytics, and ordersfbw.
- [x] Run `make generate`.
- [x] Run `go test ./...` to catch generated compilation issues.

### Task 2: Client Wiring RED/GREEN

**Files:**
- Modify: `client/client.go`
- Modify: `client/config.go`
- Modify: `client/public_types.go`
- Test: `client/*_test.go`

- [x] Write compile/route tests that reference `Client.Reports()`, `Client.Analytics()`, and `Client.OrdersFBW()`.
- [x] Run `go test ./client` and confirm missing-method/type failures.
- [x] Add API client aliases, config base URLs, `WithReportsBaseURL`, `WithAnalyticsBaseURL`, and `WithOrdersFBWBaseURL`.
- [x] Wire generated clients in `NewClient`.
- [x] Run `go test ./client`.

### Task 3: Reports Facade

**Files:**
- Create: `client/reports.go`
- Test: `client/reports_test.go`

- [x] Add RED tests for representative Reports operations: supplier stocks, warehouse remains create/status/download, acceptance report, paid storage, and one analytics report.
- [x] Implement all generated Reports operations in `ReportsService`.
- [x] Add public aliases for request, response, and model types used by Reports methods.
- [x] Run `go test ./client`.

### Task 4: Analytics Facade

**Files:**
- Create: `client/analytics.go`
- Test: `client/analytics_test.go`

- [x] Add RED tests for sales funnel, stocks report groups/products/sizes/offices/WB warehouses, and CSV create/list/retry/download.
- [x] Implement all generated Analytics operations in `AnalyticsService`.
- [x] Add public aliases for request, response, model, and CSV download types.
- [x] Run `go test ./client`.

### Task 5: Orders FBW Facade

**Files:**
- Create: `client/orders_fbw.go`
- Test: `client/orders_fbw_test.go`

- [x] Add RED tests for acceptance options, warehouses, transit tariffs, supplies list/details/goods/package.
- [x] Implement all generated Orders FBW operations in `OrdersFBWService`.
- [x] Add public aliases for request, response, and model types used by Orders FBW methods.
- [x] Run `go test ./client`.

### Task 6: Docs And Final Verification

**Files:**
- Modify: `README.md`

- [x] Update README category and facade coverage.
- [x] Run `go test ./...`.
- [ ] Commit and push `wb-api-client`.
- [ ] Update `services/wb-integration-service` dependency in `/home/user2222/simple-ozonapi-prjct`.
- [ ] Run `GOPRIVATE=github.com/benice2me11/* go test ./...` in the consumer service.
