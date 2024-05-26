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

var _ datastore.BlogTagProvider = &BlogTagStore{}

type BlogTagStore struct {
	db  *stdsql.DB
	log *slog.Logger
}

func NewBlogTagStore(db *stdsql.DB, log *slog.Logger) *BlogTagStore {
	return &BlogTagStore{
		db:  db,
		log: log.With(slog.String("DataStore", "BlogTagStore")),
	}
}

func (b BlogTagStore) Get(ctx context.Context, id int32) (*model.BlogTag, error) {
	var tag *model.BlogTag
	stmt := SELECT(BlogTag.AllColumns).
		FROM(BlogTag).
		WHERE(BlogTag.ID.EQ(Int32(id)))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Get", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &tag)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			b.log.InfoContext(ctx, "not found", slog.Int("id", int(id)))
			return nil, nil
		}
		return nil, fmt.Errorf("getting blog tag: %w", err)
	}
	return tag, nil
}

func (b BlogTagStore) AllForPost(ctx context.Context, postID int32) ([]model.BlogTag, error) {
	var tags []model.BlogTag
	stmt := SELECT(BlogTag.AllColumns).
		FROM(BlogTag.
			INNER_JOIN(M2mBlogPostTag, BlogTag.ID.EQ(M2mBlogPostTag.TagID)).
			INNER_JOIN(BlogPost, BlogPost.ID.EQ(M2mBlogPostTag.BlogPostID)),
		).
		ORDER_BY(BlogTag.DisplayName.ASC())
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "AllForPost", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &tags)
	if err != nil {
		return nil, fmt.Errorf("getting blog post tags: %w", err)
	}
	return tags, nil
}

func (b BlogTagStore) All(ctx context.Context) ([]model.BlogTag, error) {
	var tags []model.BlogTag
	stmt := SELECT(BlogTag.AllColumns).
		FROM(BlogTag).
		ORDER_BY(BlogTag.DisplayName.ASC())
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "All", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &tags)
	if err != nil {
		return nil, fmt.Errorf("getting all blog tags: %w", err)
	}
	return tags, nil
}

func (b BlogTagStore) Create(ctx context.Context, tag *model.BlogTag) error {
	stmt := BlogTag.INSERT(BlogTag.MutableColumns.Except(BlogTag.CreatedAt)).
		MODEL(tag).
		RETURNING(BlogTag.AllColumns)
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, tag)
	if err != nil {
		b.log.ErrorContext(ctx, "Create", slog.String("err", err.Error()))
		return fmt.Errorf("creating blog tag: %w", err)
	}
	return nil
}

func (b BlogTagStore) Update(ctx context.Context, tag *model.BlogTag) error {
	stmt := BlogTag.UPDATE(BlogTag.MutableColumns.Except(BlogTag.CreatedAt)).
		MODEL(tag).
		WHERE(BlogTag.ID.EQ(Int32(tag.ID)))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Update", slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		b.log.ErrorContext(ctx, "Update", slog.String("err", err.Error()))
		return fmt.Errorf("updating blog tag: %w", err)
	}
	return nil
}

func (b BlogTagStore) Delete(ctx context.Context, id int32) error {
	stmt := BlogTag.DELETE().WHERE(BlogTag.ID.EQ(Int32(id)))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Delete", slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		return fmt.Errorf("deleting blog tag: %w", err)
	}
	return nil
}
