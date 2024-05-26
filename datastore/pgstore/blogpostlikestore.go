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

var _ datastore.BlogPostLikeProvider = &BlogPostLikeStore{}

type BlogPostLikeStore struct {
	db  *stdsql.DB
	log *slog.Logger
}

func NewBlogPostLikeStore(db *stdsql.DB, log *slog.Logger) *BlogPostLikeStore {
	return &BlogPostLikeStore{
		db:  db,
		log: log.With(slog.String("DataStore", "BlogPostLikeStore")),
	}
}

func (b BlogPostLikeStore) Get(ctx context.Context, userID, postID int32) (*model.BlogPostLike, error) {
	var like *model.BlogPostLike
	stmt := SELECT(BlogPostLike.AllColumns).
		FROM(BlogPostLike).
		WHERE(BlogPostLike.UserID.EQ(Int32(userID)).AND(BlogPostLike.PostID.EQ(Int32(postID))))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Get", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &like)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			b.log.InfoContext(ctx, "not found", slog.Int("userID", int(userID)), slog.Int("postID", int(postID)))
			return nil, nil
		}
		return nil, fmt.Errorf("getting blog post like: %w", err)
	}
	return like, nil
}

func (b BlogPostLikeStore) AllForPost(ctx context.Context, postID int32) ([]model.BlogPostLike, error) {
	var likes []model.BlogPostLike
	stmt := SELECT(BlogPostLike.AllColumns).
		FROM(BlogPostLike).
		WHERE(BlogPostLike.PostID.EQ(Int32(postID)))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "AllForPost", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &likes)
	if err != nil {
		return nil, fmt.Errorf("getting all blog post likes: %w", err)
	}
	return likes, nil
}

func (b BlogPostLikeStore) CountForPost(ctx context.Context, postID int32) (int, error) {
	var count int
	stmt := SELECT(COUNT(BlogPostLike.UserID)).
		FROM(BlogPostLike).
		WHERE(BlogPostLike.PostID.EQ(Int32(postID)))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "CountForPost", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &count)
	if err != nil {
		b.log.ErrorContext(ctx, "CountForPost", slog.String("err", err.Error()))
		return 0, fmt.Errorf("counting blog post likes: %w", err)
	}
	return count, nil
}

func (b BlogPostLikeStore) Create(ctx context.Context, like *model.BlogPostLike) error {
	stmt := BlogPostLike.INSERT(BlogPostLike.MutableColumns.Except(BlogPostLike.CreatedAt)).
		MODEL(like).
		RETURNING(BlogPostLike.AllColumns)
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, like)
	if err != nil {
		b.log.ErrorContext(ctx, "Create", slog.String("err", err.Error()))
		return fmt.Errorf("creating blog post like: %w", err)
	}
	return nil
}

func (b BlogPostLikeStore) Delete(ctx context.Context, userID, postID int32) error {
	stmt := BlogPostLike.DELETE().
		WHERE(BlogPostLike.UserID.EQ(Int32(userID)).AND(BlogPostLike.PostID.EQ(Int32(postID))))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Delete", slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		return fmt.Errorf("deleting blog post like: %w", err)
	}
	return nil
}
