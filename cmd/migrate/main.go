// Package main provides a customization for the Goose DB migration tool.
// It retrieves configuration details from the default configuration provider
// and supports various commands for database migration.
//
// Supported commands are:
//   - up: Migrate the DB to the most recent version available
//   - up-by-one: Migrate the DB up by 1
//   - up-to VERSION: Migrate the DB to a specific VERSION
//   - down: Roll back the version by 1
//   - down-to VERSION: Roll back to a specific VERSION
//   - redo: Re-run the latest migration
//   - reset: Roll back all migrations
//   - status: Dump the migration status for the current DB
//   - version: Print the current version of the database
//   - create NAME [sql|go]: Creates new migration file with the current timestamp
//   - fix: Apply sequential ordering to migrations
//   - validate: Check migration files without running them
//
// All commands are run through the command line by providing the command as an argument,
// followed by any parameters the command might require.
package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/PaluMacil/barkdognet/configuration"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

func init() {
	sql.Register("postgres", stdlib.GetDefaultDriver())
}

var flags = flag.NewFlagSet("goose", flag.ExitOnError)

func main() {
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("goose: could not parse args: %v\n", err)
	}
	args := flags.Args()

	if len(args) < 1 {
		log.Fatalf("goose: no command given")
	}

	command := args[0]

	configProvider := configuration.DefaultProvider{}
	config, err := configProvider.Config()
	if err != nil {
		log.Fatalf("getting config: %v", err)
	}
	db, err := goose.OpenDBWithDriver("postgres", config.Database.ConnectionString())
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	var arguments []string
	if len(args) > 2 {
		arguments = append(arguments, args[1:]...)
	}

	dir := config.Database.MigrationsDir
	if err := goose.RunContext(context.Background(), command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
