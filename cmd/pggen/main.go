// Package main provides a customized Jet codegen tool for generating models from the live
// application database in a database first approach. This means that the migration should
// run before this codegen runs. Database connection values come from the default config
// provider, and customizations to the codegen can be made here.
package main

import (
	"database/sql"
	"github.com/PaluMacil/barkdognet/configuration"
	"github.com/PaluMacil/barkdognet/datastore/types"
	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	pgdialect "github.com/go-jet/jet/v2/postgres"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"log"
)

func init() {
	sql.Register("postgres", stdlib.GetDefaultDriver())
}

func main() {
	configProvider := configuration.DefaultProvider{}
	config, err := configProvider.Config()
	if err != nil {
		log.Fatalf("getting config: %v", err)
	}
	err = postgres.Generate("./.gen",
		postgres.DBConnection{
			Host:       config.Database.Host,
			Port:       config.Database.PortInt(),
			User:       config.Database.User,
			Password:   config.Database.Password,
			SslMode:    "disable",
			DBName:     config.Database.Database,
			SchemaName: "public",
		},
		template.Default(pgdialect.Dialect).
			UseSchema(func(schema metadata.Schema) template.Schema {
				return template.DefaultSchema(schema).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(column metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(column)

									if column.DataType.Name == "text[]" {
										defaultTableModelField.Type = template.NewType(types.TextArray{})
									}
									return defaultTableModelField
								})
						}),
					)
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
