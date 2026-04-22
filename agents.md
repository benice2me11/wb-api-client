# Wildberries API Client Prompts

## Prompt 1: Project Bootstrap (OpenAPI-first)
You are a senior Go engineer. Build a new library named `wildberries-go-client`.
Use an OpenAPI-first architecture similar to generated SDK workflows.

Requirements:
- Keep generated code in `internal/generated/<category>`.
- Keep handwritten facade code in `client/`.
- Add reusable transport middleware in `transport/`.
- Add `scripts/generate.sh` to regenerate clients from `swagger.yaml`.
- Add `Makefile` targets: `generate`, `test`, `lint`.
- Do not handwrite endpoint code that can be generated from OpenAPI.

Deliverables:
- Initial directory structure.
- Minimal compilable Go module.
- Clear README section: "How to regenerate client code".

## Prompt 2: Add a New WB API Category
Add a new Wildberries API category to an existing Go project from a provided `swagger.yaml`.

Requirements:
- Generate code into `internal/generated/<category>`.
- Never edit generated files manually.
- Create a typed facade in `client/<category>.go` for ergonomic usage.
- Add unit tests using mock transport for 200, 400, 401, and 429 responses.
- Keep compatibility with existing project style and lint rules.

Output format:
- List of changed files.
- Commands to run generation and tests.
- Short note about any assumptions.

## Prompt 3: Rate Limit and Retry Middleware
Implement HTTP transport middleware for Wildberries API with robust retry behavior.

Requirements:
- Retry only for `429` and `5xx` responses.
- If present, prioritize `X-Ratelimit-Retry` header for wait duration.
- Fallback to exponential backoff with jitter.
- Respect context cancellation and request deadlines.
- Make max retry attempts configurable.
- Log request id, status code, attempt number, and wait duration.

Testing requirements:
- Add table-driven tests for:
  - `429 -> 200` successful retry.
  - `429 -> 429` until max retries reached.
  - `5xx` retry path.
  - context timeout/cancel behavior.
