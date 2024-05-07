# Barkdog Blog

# Dev Setup

## Tool Dependencies

### Taskfile

See [Taskfile Install](https://taskfile.dev/installation/) or on MacOS:

```bash
brew install go-task
```

### UI Tooling

- yarn
- Karma CLI `npm install -g karma-cli`

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

## TODO

### Feature Areas

 - Mockery/Tests
 - ESLint
 - OpenAPI codegen
 - SMTP Configuration
 - OidcConfig table

### Security

 - Failed count on user table for failed logins
 - Delay to login attempt if Failed is too high (use configuration)
 - Verify Email
 - On create, allow Roles to be designated as Security Roles which means they have m2m to an OidcConfig
