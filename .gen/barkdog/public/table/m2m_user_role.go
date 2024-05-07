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

var M2mUserRole = newM2mUserRoleTable("public", "m2m_user_role", "")

type m2mUserRoleTable struct {
	postgres.Table

	// Columns
	SysUserID postgres.ColumnInteger
	SysRoleID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type M2mUserRoleTable struct {
	m2mUserRoleTable

	EXCLUDED m2mUserRoleTable
}

// AS creates new M2mUserRoleTable with assigned alias
func (a M2mUserRoleTable) AS(alias string) *M2mUserRoleTable {
	return newM2mUserRoleTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new M2mUserRoleTable with assigned schema name
func (a M2mUserRoleTable) FromSchema(schemaName string) *M2mUserRoleTable {
	return newM2mUserRoleTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new M2mUserRoleTable with assigned table prefix
func (a M2mUserRoleTable) WithPrefix(prefix string) *M2mUserRoleTable {
	return newM2mUserRoleTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new M2mUserRoleTable with assigned table suffix
func (a M2mUserRoleTable) WithSuffix(suffix string) *M2mUserRoleTable {
	return newM2mUserRoleTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newM2mUserRoleTable(schemaName, tableName, alias string) *M2mUserRoleTable {
	return &M2mUserRoleTable{
		m2mUserRoleTable: newM2mUserRoleTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newM2mUserRoleTableImpl("", "excluded", ""),
	}
}

func newM2mUserRoleTableImpl(schemaName, tableName, alias string) m2mUserRoleTable {
	var (
		SysUserIDColumn = postgres.IntegerColumn("sys_user_id")
		SysRoleIDColumn = postgres.IntegerColumn("sys_role_id")
		allColumns      = postgres.ColumnList{SysUserIDColumn, SysRoleIDColumn}
		mutableColumns  = postgres.ColumnList{}
	)

	return m2mUserRoleTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		SysUserID: SysUserIDColumn,
		SysRoleID: SysRoleIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}