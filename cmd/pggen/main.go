package main

import (
	"database/sql"
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
	err := postgres.Generate("./.gen",
		postgres.DBConnection{
			Host:       "localhost",
			Port:       5432,
			User:       "barkadmin",
			Password:   "",
			SslMode:    "disable",
			DBName:     "barkdog",
			SchemaName: "public",
		},
		template.Default(pgdialect.Dialect).
			UseSchema(func(schema metadata.Schema) template.Schema {
				return template.DefaultSchema(schema).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(column metadata.Column) template.TableModelField {
									return template.DefaultTableModelField(column)
								})
						}),
					)
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
