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

var BlogCommentLike = newBlogCommentLikeTable("public", "blog_comment_like", "")

type blogCommentLikeTable struct {
	postgres.Table

	// Columns
	UserID    postgres.ColumnInteger
	CommentID postgres.ColumnInteger
	CreatedAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BlogCommentLikeTable struct {
	blogCommentLikeTable

	EXCLUDED blogCommentLikeTable
}

// AS creates new BlogCommentLikeTable with assigned alias
func (a BlogCommentLikeTable) AS(alias string) *BlogCommentLikeTable {
	return newBlogCommentLikeTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BlogCommentLikeTable with assigned schema name
func (a BlogCommentLikeTable) FromSchema(schemaName string) *BlogCommentLikeTable {
	return newBlogCommentLikeTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BlogCommentLikeTable with assigned table prefix
func (a BlogCommentLikeTable) WithPrefix(prefix string) *BlogCommentLikeTable {
	return newBlogCommentLikeTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BlogCommentLikeTable with assigned table suffix
func (a BlogCommentLikeTable) WithSuffix(suffix string) *BlogCommentLikeTable {
	return newBlogCommentLikeTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBlogCommentLikeTable(schemaName, tableName, alias string) *BlogCommentLikeTable {
	return &BlogCommentLikeTable{
		blogCommentLikeTable: newBlogCommentLikeTableImpl(schemaName, tableName, alias),
		EXCLUDED:             newBlogCommentLikeTableImpl("", "excluded", ""),
	}
}

func newBlogCommentLikeTableImpl(schemaName, tableName, alias string) blogCommentLikeTable {
	var (
		UserIDColumn    = postgres.IntegerColumn("user_id")
		CommentIDColumn = postgres.IntegerColumn("comment_id")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		allColumns      = postgres.ColumnList{UserIDColumn, CommentIDColumn, CreatedAtColumn}
		mutableColumns  = postgres.ColumnList{CreatedAtColumn}
	)

	return blogCommentLikeTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:    UserIDColumn,
		CommentID: CommentIDColumn,
		CreatedAt: CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
