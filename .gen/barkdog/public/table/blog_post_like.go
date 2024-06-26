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

var BlogPostLike = newBlogPostLikeTable("public", "blog_post_like", "")

type blogPostLikeTable struct {
	postgres.Table

	// Columns
	UserID    postgres.ColumnInteger
	PostID    postgres.ColumnInteger
	CreatedAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BlogPostLikeTable struct {
	blogPostLikeTable

	EXCLUDED blogPostLikeTable
}

// AS creates new BlogPostLikeTable with assigned alias
func (a BlogPostLikeTable) AS(alias string) *BlogPostLikeTable {
	return newBlogPostLikeTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BlogPostLikeTable with assigned schema name
func (a BlogPostLikeTable) FromSchema(schemaName string) *BlogPostLikeTable {
	return newBlogPostLikeTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BlogPostLikeTable with assigned table prefix
func (a BlogPostLikeTable) WithPrefix(prefix string) *BlogPostLikeTable {
	return newBlogPostLikeTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BlogPostLikeTable with assigned table suffix
func (a BlogPostLikeTable) WithSuffix(suffix string) *BlogPostLikeTable {
	return newBlogPostLikeTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBlogPostLikeTable(schemaName, tableName, alias string) *BlogPostLikeTable {
	return &BlogPostLikeTable{
		blogPostLikeTable: newBlogPostLikeTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newBlogPostLikeTableImpl("", "excluded", ""),
	}
}

func newBlogPostLikeTableImpl(schemaName, tableName, alias string) blogPostLikeTable {
	var (
		UserIDColumn    = postgres.IntegerColumn("user_id")
		PostIDColumn    = postgres.IntegerColumn("post_id")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		allColumns      = postgres.ColumnList{UserIDColumn, PostIDColumn, CreatedAtColumn}
		mutableColumns  = postgres.ColumnList{CreatedAtColumn}
	)

	return blogPostLikeTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:    UserIDColumn,
		PostID:    PostIDColumn,
		CreatedAt: CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
