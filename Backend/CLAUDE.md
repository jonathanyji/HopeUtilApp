# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Layout

This is a monorepo with three top-level directories:

- `Backend/` — Go/Gin HTTP API (this directory, module `hope-backend-gin`)
- `Frontend/` — placeholder, not yet implemented
- `DevOps/localdev/` — Docker Compose stack (PostgreSQL + pgAdmin) for local development

## Local Dev Setup

**Start the database stack** (run from repo root or `DevOps/localdev/`):

```bash
cd DevOps/localdev
cp .env.example .env   # first time only; edit credentials if needed
docker compose up -d
```

Default local credentials (from `.env.example`):
- `POSTGRES_USER=user`, `POSTGRES_PASSWORD=admin`, `POSTGRES_DB=hope_church_db`
- pgAdmin: http://localhost:5050 — `admin@local.dev` / `admin`

> Apple Silicon: the `platform: linux/amd64` line in `docker-compose.yml` is required. Remove it on Intel/Linux.

**Stop:**
```bash
docker compose down
```

## Backend Commands

All commands run from `Backend/`:

```bash
go run .          # run the server (listens on 0.0.0.0:8080 by default)
go build .        # compile binary
go test ./...     # run all tests
go test ./... -run TestFooBar   # run a single test
go vet ./...      # static analysis
```

## API Response Shape

All endpoints return a consistent JSON envelope (defined in `main.go`):

```json
{ "success": true,  "data": {...} }
{ "success": false, "error": { "code": "...", "message": "..." } }
```

Use the `OK(c, data)` and `Fail(c, status, code, message)` helpers for all handler responses — do not call `c.JSON` directly.

## Key Dependencies

- **gin** — HTTP router/framework
- **mongo-driver v2** — MongoDB client (imported in go.mod, not yet wired up)
- **quic-go** — HTTP/3 support (imported in go.mod, not yet wired up)
