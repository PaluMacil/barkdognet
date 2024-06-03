package configuration

import (
	"cmp"
	"fmt"
	"strconv"
)

type Database struct {
	Host          string `koanf:"host"`
	Port          string `koanf:"port"`
	User          string `koanf:"user"`
	Password      string `koanf:"password"`
	Database      string `koanf:"database"`
	MigrationsDir string `koanf:"migrations_dir"`
}

func (dbc *Database) ConnectionString() string {
	host := cmp.Or(dbc.Host, "localhost")
	port := cmp.Or(dbc.Port, "5432")
	user := cmp.Or(dbc.User, "barkadmin")
	password := cmp.Or(dbc.Password, "")
	database := cmp.Or(dbc.Database, "barkdog")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
}

func (dbc *Database) PortInt() int {
	i, _ := strconv.Atoi(dbc.Port)
	if i == 0 {
		return 5432
	}
	return i
}
