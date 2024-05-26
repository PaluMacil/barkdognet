package pgstore

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	. "github.com/PaluMacil/barkdognet/.gen/barkdog/public/table"
	"github.com/PaluMacil/barkdognet/datastore"
	"github.com/PaluMacil/barkdognet/datastore/identifier"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"log/slog"
)

var _ datastore.BlogCommentProvider = BlogCommentStore{}

type BlogCommentStore struct {
	db  *stdsql.DB
	log *slog.Logger
}

func NewBlogCommentStore(db *stdsql.DB, log *slog.Logger) *BlogCommentStore {
	return &BlogCommentStore{
		db:  db,
		log: log.With(slog.String("DataStore", "BlogCommentStore")),
	}
}

func (b BlogCommentStore) Get(ctx context.Context, id int32) (*model.BlogComment, error) {
	var comment model.BlogComment
	stmt := SELECT(BlogComment.AllColumns).
		FROM(BlogComment).
		WHERE(BlogComment.ID.EQ(Int32(id)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Get", slog.String("sql", stmt.DebugSql()), slog.Int("id", int(id)))
	}

	err := stmt.QueryContext(ctx, b.db, &comment)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			b.log.InfoContext(ctx, "not found", slog.Int("id", int(id)))
			return nil, nil
		}
		b.log.ErrorContext(ctx, "Get", slog.Int("id", int(id)), slog.String("err", err.Error()))
		return nil, fmt.Errorf("getting comment %d: %w", id, err)
	}

	return &comment, nil
}

func (b BlogCommentStore) SomeForBlogPost(ctx context.Context, blogPostID int32, page, pageSize int64) ([]model.BlogComment, error) {
	var comments []model.BlogComment
	stmt := SELECT(BlogComment.AllColumns).
		FROM(BlogComment).
		WHERE(BlogComment.BlogPostID.EQ(Int32(blogPostID))).
		ORDER_BY(BlogComment.CreatedAt.ASC()).
		LIMIT(pageSize).
		OFFSET(page * pageSize)

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "SomeForBlogPost", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, &comments)
	if err != nil {
		b.log.ErrorContext(ctx, "SomeForBlogPost", slog.String("err", err.Error()))
		return nil, fmt.Errorf("getting some blog post comments: %w", err)
	}

	return comments, nil
}

func (b BlogCommentStore) AllForBlogPost(ctx context.Context, blogPostID int32) ([]model.BlogComment, error) {
	var comments []model.BlogComment
	stmt := SELECT(BlogComment.AllColumns).
		FROM(BlogComment).
		WHERE(BlogComment.BlogPostID.EQ(Int32(blogPostID))).
		ORDER_BY(BlogComment.CreatedAt.ASC())

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "AllForBlogPost", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, &comments)
	if err != nil {
		b.log.ErrorContext(ctx, "AllForBlogPost", slog.String("err", err.Error()))
		return nil, fmt.Errorf("getting all blog post comments: %w", err)
	}

	return comments, nil
}

func (b BlogCommentStore) SomeForUser(ctx context.Context, userIden identifier.User, page, pageSize int64) ([]model.BlogComment, error) {
	var comments []model.BlogComment
	stmt := SELECT(BlogComment.AllColumns).
		FROM(BlogComment).
		WHERE(BlogComment.AuthorID.EQ(Int32(*userIden.ID))).
		ORDER_BY(BlogComment.CreatedAt.ASC()).
		LIMIT(pageSize).
		OFFSET(page * pageSize)

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "SomeForUser", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, &comments)
	if err != nil {
		b.log.ErrorContext(ctx, "SomeForUser", slog.String("err", err.Error()))
		return nil, fmt.Errorf("getting some comments for user: %w", err)
	}

	return comments, nil
}

func (b BlogCommentStore) Create(ctx context.Context, blogComment *model.BlogComment) error {
	stmt := BlogComment.INSERT(BlogComment.MutableColumns.Except(BlogComment.CreatedAt)).
		MODEL(blogComment).
		RETURNING(BlogComment.AllColumns)

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, blogComment)
	if err != nil {
		b.log.ErrorContext(ctx, "Create", slog.String("err", err.Error()))
		return fmt.Errorf("creating blog comment: %w", err)
	}

	return nil
}

func (b BlogCommentStore) Update(ctx context.Context, blogComment *model.BlogComment) error {
	stmt := BlogComment.UPDATE(BlogComment.MutableColumns.Except(BlogComment.CreatedAt)).
		MODEL(blogComment).
		WHERE(BlogComment.ID.EQ(Int32(blogComment.ID)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Update", slog.String("sql", stmt.DebugSql()))
	}

	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		b.log.ErrorContext(ctx, "Update", slog.String("err", err.Error()))
		return fmt.Errorf("updating blog comment: %w", err)
	}

	return nil
}

func (b BlogCommentStore) Delete(ctx context.Context, id int32) error {
	stmt := BlogComment.DELETE().
		WHERE(BlogComment.ID.EQ(Int32(id)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Delete", slog.String("sql", stmt.DebugSql()))
	}

	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		b.log.ErrorContext(ctx, "Delete", slog.String("err", err.Error()))
		return fmt.Errorf("deleting blog comment: %w", err)
	}

	return nil
}
