# HopeUtilApp — Local Dev Setup

## Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running

## Steps

1. Copy the env file and set your credentials:

   ```bash
   cd DevOps/localdev
   cp .env.example .env
   ```

   Edit `.env` with your own username, password, and database name. Then update `servers.json` to match — the `Username`, `Password`, and `MaintenanceDB` fields must be the same.

2. Start the services (from inside the `DevOps/localdev` folder):

   ```bash
   docker compose up -d
   ```

3. Open pgAdmin at [http://localhost:5050](http://localhost:5050)
   - Email: `admin@local.dev`
   - Password: `admin`

The PostgreSQL connection will already be configured in the sidebar — no extra setup needed.

Credentials for local dev only
POSTGRES_USER=user
POSTGRES_PASSWORD=admin
POSTGRES_DB=hope_church_db

## Stop

```bash
docker compose down
```

> **Apple Silicon users:** The `platform: linux/amd64` line in `docker-compose.yml` is required. Remove it if you are on Intel/Linux.
