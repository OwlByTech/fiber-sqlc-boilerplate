<h1 align="center">
    Go API Bootstrap
</h1>

## Getting Started
To start the project in development mode, you must copy the `.env.example` to `.env`:
and then you should run the following command:

```bash
docker compose -f compose-dev.yml
```

Then you must run the migration and seeder inside the app container.

```bash
docker compose -f compose-dev.yml exec app  bash
```
And you must first run the migrations using the cli

```bash
 go run cmd/cli/main.go migrate up
```
and then the seeder
```bash
 go run cmd/cli/main.go seed up
```
