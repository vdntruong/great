## Install CLI

[Installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)

### MacOS
```bash
brew install golang-migrate
```

## Create migrations

```bash
migrate create -ext sql -dir PATH_TO_YOUR_MIGRATIONS -seq YOUR_MIGRATION_NAME
```

Example:
```bash
migrate create -ext sql -dir db/migrations -seq create_products_table
```

## Run migrations

```bash
migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up
```

## Forcing our database version

```bash
migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force VERSION
```
