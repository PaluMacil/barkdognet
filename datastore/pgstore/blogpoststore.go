package pgstore

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	. "github.com/PaluMacil/barkdognet/.gen/barkdog/public/table"
	"github.com/PaluMacil/barkdognet/datastore"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"log/slog"
)

var _ datastore.BlogPostProvider = &BlogPostStore{}

type BlogPostStore struct {
	db  *stdsql.DB
	log *slog.Logger
}

func NewBlogPostStore(db *stdsql.DB, log *slog.Logger) *BlogPostStore {
	return &BlogPostStore{
		db:  db,
		log: log.With(slog.String("DataStore", "BlogPostStore")),
	}
}

func (b BlogPostStore) Get(ctx context.Context, id int32) (*model.BlogPost, error) {
	var blogPost *model.BlogPost
	stmt := SELECT(BlogPost.AllColumns).
		FROM(BlogPost).
		WHERE(BlogPost.ID.EQ(Int32(id)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Get", slog.String("sql", stmt.DebugSql()), slog.Int("id", int(id)))
	}

	err := stmt.QueryContext(ctx, b.db, &blogPost)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			b.log.InfoContext(ctx, "not found", slog.Int("id", int(id)))
			return nil, nil
		}
		return nil, fmt.Errorf("getting blog post: %w", err)
	}
	return blogPost, nil
}

func (b BlogPostStore) Some(ctx context.Context, categoryID *int32, page, pageSize int64) ([]model.BlogPost, error) {
	var blogPosts []model.BlogPost
	stmt := SELECT(BlogPost.AllColumns).
		FROM(BlogPost).
		ORDER_BY(BlogPost.CreatedAt.DESC()).
		LIMIT(pageSize).
		OFFSET(page * pageSize)
	if categoryID != nil {
		stmt = stmt.WHERE(BlogPost.CategoryID.EQ(Int32(*categoryID)))
	}

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Some", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, &blogPosts)
	if err != nil {
		return nil, fmt.Errorf("getting some Blog Posts: %w", err)
	}
	return blogPosts, nil
}

func (b BlogPostStore) GetTags(ctx context.Context, postID int32) ([]model.BlogTag, error) {
	var tags []model.BlogTag
	stmt := SELECT(BlogTag.AllColumns).
		FROM(BlogTag.
			INNER_JOIN(M2mBlogPostTag, BlogTag.ID.EQ(M2mBlogPostTag.TagID))).
		WHERE(M2mBlogPostTag.BlogPostID.EQ(Int32(postID)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "GetTags", slog.String("sql", stmt.DebugSql()), slog.Int("postID", int(postID)))
	}

	err := stmt.QueryContext(ctx, b.db, &tags)
	if err != nil {
		b.log.ErrorContext(ctx, "GetTags", slog.Int("postID", int(postID)), slog.String("err", err.Error()))
		return nil, fmt.Errorf("getting tags for post %d: %w", postID, err)
	}

	return tags, nil
}

func (b BlogPostStore) All(ctx context.Context, categoryID *int32) ([]model.BlogPost, error) {
	var blogPosts []model.BlogPost
	stmt := SELECT(BlogPost.AllColumns).
		FROM(BlogPost).
		ORDER_BY(BlogPost.CreatedAt.DESC())
	if categoryID != nil {
		stmt = stmt.WHERE(BlogPost.CategoryID.EQ(Int32(*categoryID)))
	}

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "All", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, &blogPosts)
	if err != nil {
		return nil, fmt.Errorf("getting All Blog Posts: %w", err)
	}
	return blogPosts, nil
}

func (b BlogPostStore) Create(ctx context.Context, blogPost *model.BlogPost) error {
	stmt := BlogPost.INSERT(BlogPost.MutableColumns.Except(BlogPost.CreatedAt)).
		MODEL(blogPost).
		RETURNING(BlogPost.AllColumns)
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, blogPost)
	if err != nil {
		b.log.ErrorContext(ctx, "Create", slog.String("err", err.Error()))
		return fmt.Errorf("creating blog post: %w", err)
	}
	return nil
}

func (b BlogPostStore) Update(ctx context.Context, blogPost *model.BlogPost) error {
	stmt := BlogPost.UPDATE(BlogPost.MutableColumns.Except(BlogPost.CreatedAt)).
		MODEL(blogPost).
		WHERE(BlogPost.ID.EQ(Int32(blogPost.ID)))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Update", slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		b.log.ErrorContext(ctx, "Update", slog.String("err", err.Error()))
		return fmt.Errorf("updating blog post: %w", err)
	}
	return nil
}

func (b BlogPostStore) Delete(ctx context.Context, id int32) error {
	stmt := BlogPost.DELETE().WHERE(BlogPost.ID.EQ(Int32(id)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Delete", slog.String("sql", stmt.DebugSql()), slog.Int("id", int(id)))
	}

	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		return fmt.Errorf("deleting blog post: %w", err)
	}
	return nil
}
