# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands
- Build: `make build`
- Run: `make run`
- Test all: `make test`
- Test single test: `go test -v -run TestName ./path/to/package` (e.g., `go test -v -run TestAnthropic ./internal/api`)
- Verify (format/lint): `make verify` (runs go mod tidy, verify, vet, fmt)
- Complete check: `make check`

## Code Style Guidelines
- **Imports**: Standard Go convention with grouped imports (stdlib first, then external packages)
- **Formatting**: Use `go fmt` for consistent code formatting
- **Types**: Use explicit type definitions with clear struct definitions
- **Naming**: Follow Go conventions (CamelCase for exported, camelCase for unexported)
- **Error Handling**: Handle errors explicitly with clear error messages
- **Logging**: Use zerolog for structured logging
- **Testing**: Use `t.Helper()` for test helper methods, use pkg/assert for assertions
- **Context**: Use context with appropriate timeouts for cancellation
- **Comments**: Add comments for complex logic, particularly around concurrency

## Project Structure
- Follow standard Go project layout with cmd/, internal/, and pkg/ directories
- Use internal/secrets for Google Cloud Secret Manager integration