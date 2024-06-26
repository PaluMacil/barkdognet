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

var M2mBlogPostTag = newM2mBlogPostTagTable("public", "m2m_blog_post_tag", "")

type m2mBlogPostTagTable struct {
	postgres.Table

	// Columns
	BlogPostID postgres.ColumnInteger
	TagID      postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type M2mBlogPostTagTable struct {
	m2mBlogPostTagTable

	EXCLUDED m2mBlogPostTagTable
}

// AS creates new M2mBlogPostTagTable with assigned alias
func (a M2mBlogPostTagTable) AS(alias string) *M2mBlogPostTagTable {
	return newM2mBlogPostTagTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new M2mBlogPostTagTable with assigned schema name
func (a M2mBlogPostTagTable) FromSchema(schemaName string) *M2mBlogPostTagTable {
	return newM2mBlogPostTagTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new M2mBlogPostTagTable with assigned table prefix
func (a M2mBlogPostTagTable) WithPrefix(prefix string) *M2mBlogPostTagTable {
	return newM2mBlogPostTagTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new M2mBlogPostTagTable with assigned table suffix
func (a M2mBlogPostTagTable) WithSuffix(suffix string) *M2mBlogPostTagTable {
	return newM2mBlogPostTagTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newM2mBlogPostTagTable(schemaName, tableName, alias string) *M2mBlogPostTagTable {
	return &M2mBlogPostTagTable{
		m2mBlogPostTagTable: newM2mBlogPostTagTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newM2mBlogPostTagTableImpl("", "excluded", ""),
	}
}

func newM2mBlogPostTagTableImpl(schemaName, tableName, alias string) m2mBlogPostTagTable {
	var (
		BlogPostIDColumn = postgres.IntegerColumn("blog_post_id")
		TagIDColumn      = postgres.IntegerColumn("tag_id")
		allColumns       = postgres.ColumnList{BlogPostIDColumn, TagIDColumn}
		mutableColumns   = postgres.ColumnList{}
	)

	return m2mBlogPostTagTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		BlogPostID: BlogPostIDColumn,
		TagID:      TagIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
