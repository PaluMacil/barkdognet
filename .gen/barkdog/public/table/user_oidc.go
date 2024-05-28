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

var UserOidc = newUserOidcTable("public", "user_oidc", "")

type userOidcTable struct {
	postgres.Table

	// Columns
	SysUserID      postgres.ColumnInteger
	OidcProviderID postgres.ColumnInteger
	Sub            postgres.ColumnString
	CreatedAt      postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserOidcTable struct {
	userOidcTable

	EXCLUDED userOidcTable
}

// AS creates new UserOidcTable with assigned alias
func (a UserOidcTable) AS(alias string) *UserOidcTable {
	return newUserOidcTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserOidcTable with assigned schema name
func (a UserOidcTable) FromSchema(schemaName string) *UserOidcTable {
	return newUserOidcTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserOidcTable with assigned table prefix
func (a UserOidcTable) WithPrefix(prefix string) *UserOidcTable {
	return newUserOidcTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserOidcTable with assigned table suffix
func (a UserOidcTable) WithSuffix(suffix string) *UserOidcTable {
	return newUserOidcTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserOidcTable(schemaName, tableName, alias string) *UserOidcTable {
	return &UserOidcTable{
		userOidcTable: newUserOidcTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newUserOidcTableImpl("", "excluded", ""),
	}
}

func newUserOidcTableImpl(schemaName, tableName, alias string) userOidcTable {
	var (
		SysUserIDColumn      = postgres.IntegerColumn("sys_user_id")
		OidcProviderIDColumn = postgres.IntegerColumn("oidc_provider_id")
		SubColumn            = postgres.StringColumn("sub")
		CreatedAtColumn      = postgres.TimestampzColumn("created_at")
		allColumns           = postgres.ColumnList{SysUserIDColumn, OidcProviderIDColumn, SubColumn, CreatedAtColumn}
		mutableColumns       = postgres.ColumnList{SubColumn, CreatedAtColumn}
	)

	return userOidcTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		SysUserID:      SysUserIDColumn,
		OidcProviderID: OidcProviderIDColumn,
		Sub:            SubColumn,
		CreatedAt:      CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
