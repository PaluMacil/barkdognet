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

var BlogComment = newBlogCommentTable("public", "blog_comment", "")

type blogCommentTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnInteger
	AuthorID   postgres.ColumnInteger
	BlogPostID postgres.ColumnInteger
	Body       postgres.ColumnString
	Status     postgres.ColumnString
	CreatedAt  postgres.ColumnTimestampz
	TenantID   postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BlogCommentTable struct {
	blogCommentTable

	EXCLUDED blogCommentTable
}

// AS creates new BlogCommentTable with assigned alias
func (a BlogCommentTable) AS(alias string) *BlogCommentTable {
	return newBlogCommentTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BlogCommentTable with assigned schema name
func (a BlogCommentTable) FromSchema(schemaName string) *BlogCommentTable {
	return newBlogCommentTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BlogCommentTable with assigned table prefix
func (a BlogCommentTable) WithPrefix(prefix string) *BlogCommentTable {
	return newBlogCommentTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BlogCommentTable with assigned table suffix
func (a BlogCommentTable) WithSuffix(suffix string) *BlogCommentTable {
	return newBlogCommentTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBlogCommentTable(schemaName, tableName, alias string) *BlogCommentTable {
	return &BlogCommentTable{
		blogCommentTable: newBlogCommentTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newBlogCommentTableImpl("", "excluded", ""),
	}
}

func newBlogCommentTableImpl(schemaName, tableName, alias string) blogCommentTable {
	var (
		IDColumn         = postgres.IntegerColumn("id")
		AuthorIDColumn   = postgres.IntegerColumn("author_id")
		BlogPostIDColumn = postgres.IntegerColumn("blog_post_id")
		BodyColumn       = postgres.StringColumn("body")
		StatusColumn     = postgres.StringColumn("status")
		CreatedAtColumn  = postgres.TimestampzColumn("created_at")
		TenantIDColumn   = postgres.IntegerColumn("tenant_id")
		allColumns       = postgres.ColumnList{IDColumn, AuthorIDColumn, BlogPostIDColumn, BodyColumn, StatusColumn, CreatedAtColumn, TenantIDColumn}
		mutableColumns   = postgres.ColumnList{AuthorIDColumn, BlogPostIDColumn, BodyColumn, StatusColumn, CreatedAtColumn, TenantIDColumn}
	)

	return blogCommentTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		AuthorID:   AuthorIDColumn,
		BlogPostID: BlogPostIDColumn,
		Body:       BodyColumn,
		Status:     StatusColumn,
		CreatedAt:  CreatedAtColumn,
		TenantID:   TenantIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
