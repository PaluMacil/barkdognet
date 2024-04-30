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
