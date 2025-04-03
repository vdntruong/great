## Install CLI

[Installation](https://docs.sqlc.dev/en/stable/overview/install.html)

### MacOS

```bash
brew install sqlc
```

## Components

### Configuration file (`sqlc.yaml` | `sqlc.yml` | `sqlc.json`)

Example:
```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "tutorial"
        out: "tutorial"
        sql_package: "pgx/v5"
```

### Schema files

Configured by `sql.schema` in the configuration file.

### Query files

Configured by `sql.queries` in the configuration file.

## Generation code

```bash
sql generate
```
