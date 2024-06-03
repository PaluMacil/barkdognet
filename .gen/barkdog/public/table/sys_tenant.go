//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var SysTenant = newSysTenantTable("public", "sys_tenant", "")

type sysTenantTable struct {
	postgres.Table

	// Columns
	ID           postgres.ColumnInteger
	DisplayName  postgres.ColumnString
	APISubdomain postgres.ColumnString
	UIDomain     postgres.ColumnString
	CreatedAt    postgres.ColumnTimestampz
	Active       postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SysTenantTable struct {
	sysTenantTable

	EXCLUDED sysTenantTable
}

// AS creates new SysTenantTable with assigned alias
func (a SysTenantTable) AS(alias string) *SysTenantTable {
	return newSysTenantTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SysTenantTable with assigned schema name
func (a SysTenantTable) FromSchema(schemaName string) *SysTenantTable {
	return newSysTenantTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SysTenantTable with assigned table prefix
func (a SysTenantTable) WithPrefix(prefix string) *SysTenantTable {
	return newSysTenantTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SysTenantTable with assigned table suffix
func (a SysTenantTable) WithSuffix(suffix string) *SysTenantTable {
	return newSysTenantTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSysTenantTable(schemaName, tableName, alias string) *SysTenantTable {
	return &SysTenantTable{
		sysTenantTable: newSysTenantTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newSysTenantTableImpl("", "excluded", ""),
	}
}

func newSysTenantTableImpl(schemaName, tableName, alias string) sysTenantTable {
	var (
		IDColumn           = postgres.IntegerColumn("id")
		DisplayNameColumn  = postgres.StringColumn("display_name")
		APISubdomainColumn = postgres.StringColumn("api_subdomain")
		UIDomainColumn     = postgres.StringColumn("ui_domain")
		CreatedAtColumn    = postgres.TimestampzColumn("created_at")
		ActiveColumn       = postgres.BoolColumn("active")
		allColumns         = postgres.ColumnList{IDColumn, DisplayNameColumn, APISubdomainColumn, UIDomainColumn, CreatedAtColumn, ActiveColumn}
		mutableColumns     = postgres.ColumnList{DisplayNameColumn, APISubdomainColumn, UIDomainColumn, CreatedAtColumn, ActiveColumn}
	)

	return sysTenantTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		DisplayName:  DisplayNameColumn,
		APISubdomain: APISubdomainColumn,
		UIDomain:     UIDomainColumn,
		CreatedAt:    CreatedAtColumn,
		Active:       ActiveColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}