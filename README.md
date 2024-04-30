# Barkdog Blog

# Dev Setup

## Tool Dependencies

### Taskfile

See [Taskfile Install](https://taskfile.dev/installation/) or on MacOS:

```bash
brew install go-task
```

## Database

```sql
CREATE DATABASE barkdog;
CREATE USER barkadmin LOGIN;
ALTER DATABASE barkdog OWNER TO barkadmin;
```

## Codegen

### Database Model Generation

```bash
go build -o ./dist/pggen/pggen ./cmd/pggen
./dist/pggen/pggen
```

### Mock Generation

### API Generation

## Migrations

### Build Migrator

```bash
go build -o ./dist/migrate/migrate ./cmd/migrate
```

### Examples

```bash
./dist/migrate/migrate status
./dist/migrate/migrate create init sql
./dist/migrate/migrate create add_some_column sql
./dist/migrate/migrate create fetch_user_data go
./dist/migrate/migrate up
```