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

On Linux or Mac, `task resetdb` can do this.

```sql
CREATE DATABASE barkdog;
CREATE USER barkadmin LOGIN;
ALTER DATABASE barkdog OWNER TO barkadmin;
```

You can use the database with no password if this fits your security posture:

- Set the connection to trust (find pg_hba.conf with `SHOW hba_file;` using an admin)
- Restart Postgres
  - Win:
  ```
  net stop postgresql-x64-13
  net start postgresql-x64-13
  ```
  - Mac (if installed via Brew):
  ```
  brew services restart postgresql
  ```
  - Linux
  ```
  sudo systemctl restart postgresql
  ```

Otherwise, set a password:

```sql
ALTER USER barkadmin WITH PASSWORD 'localdevpw';
```

You will need to specify the password in [barkconf.yaml](barkconf.yaml) or in env var `BARKDOG_DATABASE_PASSWORD`.

If using Windows and WSL2, if you install Postgres on the Windows side, you should run configure the host in the config or set this env var: `BARKDOG_DATABASE_HOST=host.docker.internal` inside WSL2.

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
