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

var BlogTag = newBlogTagTable("public", "blog_tag", "")

type blogTagTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	DisplayName postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BlogTagTable struct {
	blogTagTable

	EXCLUDED blogTagTable
}

// AS creates new BlogTagTable with assigned alias
func (a BlogTagTable) AS(alias string) *BlogTagTable {
	return newBlogTagTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BlogTagTable with assigned schema name
func (a BlogTagTable) FromSchema(schemaName string) *BlogTagTable {
	return newBlogTagTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BlogTagTable with assigned table prefix
func (a BlogTagTable) WithPrefix(prefix string) *BlogTagTable {
	return newBlogTagTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BlogTagTable with assigned table suffix
func (a BlogTagTable) WithSuffix(suffix string) *BlogTagTable {
	return newBlogTagTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBlogTagTable(schemaName, tableName, alias string) *BlogTagTable {
	return &BlogTagTable{
		blogTagTable: newBlogTagTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newBlogTagTableImpl("", "excluded", ""),
	}
}

func newBlogTagTableImpl(schemaName, tableName, alias string) blogTagTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		DisplayNameColumn = postgres.StringColumn("display_name")
		allColumns        = postgres.ColumnList{IDColumn, DisplayNameColumn}
		mutableColumns    = postgres.ColumnList{DisplayNameColumn}
	)

	return blogTagTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		DisplayName: DisplayNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
