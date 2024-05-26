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

var _ datastore.BlogCommentLikeProvider = &BlogCommentLikeStore{}

type BlogCommentLikeStore struct {
	db  *stdsql.DB
	log *slog.Logger
}

func NewBlogCommentLikeStore(db *stdsql.DB, log *slog.Logger) *BlogCommentLikeStore {
	return &BlogCommentLikeStore{
		db:  db,
		log: log.With(slog.String("DataStore", "BlogCommentLikeStore")),
	}
}

func (b BlogCommentLikeStore) Get(ctx context.Context, userID, commentID int32) (*model.BlogCommentLike, error) {
	var like *model.BlogCommentLike
	stmt := SELECT(BlogCommentLike.AllColumns).
		FROM(BlogCommentLike).
		WHERE(BlogCommentLike.UserID.EQ(Int32(userID)).AND(BlogCommentLike.CommentID.EQ(Int32(commentID))))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Get", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &like)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			b.log.InfoContext(ctx, "not found", slog.Int("userID", int(userID)), slog.Int("commentID", int(commentID)))
			return nil, nil
		}
		return nil, fmt.Errorf("getting blog comment like: %w", err)
	}
	return like, nil
}

func (b BlogCommentLikeStore) AllForComment(ctx context.Context, commentID int32) ([]model.BlogCommentLike, error) {
	var likes []model.BlogCommentLike
	stmt := SELECT(BlogCommentLike.AllColumns).
		FROM(BlogCommentLike).
		WHERE(BlogCommentLike.CommentID.EQ(Int32(commentID))).
		ORDER_BY(BlogCommentLike.CreatedAt.DESC())
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "AllForComment", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &likes)
	if err != nil {
		return nil, fmt.Errorf("getting all blog comment likes: %w", err)
	}
	return likes, nil
}

func (b BlogCommentLikeStore) CountForComment(ctx context.Context, commentID int32) (int, error) {
	var count int
	stmt := SELECT(COUNT(BlogCommentLike.UserID)).
		FROM(BlogCommentLike).
		WHERE(BlogCommentLike.CommentID.EQ(Int32(commentID)))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "CountForComment", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, &count)
	if err != nil {
		b.log.ErrorContext(ctx, "CountForComment", slog.String("err", err.Error()))
		return 0, fmt.Errorf("counting blog comment likes: %w", err)
	}
	return count, nil
}

func (b BlogCommentLikeStore) Create(ctx context.Context, like *model.BlogCommentLike) error {
	stmt := BlogCommentLike.INSERT(BlogCommentLike.MutableColumns.Except(BlogCommentLike.CreatedAt)).
		MODEL(like).
		RETURNING(BlogCommentLike.AllColumns)
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, b.db, like)
	if err != nil {
		b.log.ErrorContext(ctx, "Create", slog.String("err", err.Error()))
		return fmt.Errorf("creating blog comment like: %w", err)
	}
	return nil
}

func (b BlogCommentLikeStore) Delete(ctx context.Context, userID, commentID int32) error {
	stmt := BlogCommentLike.DELETE().
		WHERE(BlogCommentLike.UserID.EQ(Int32(userID)).AND(BlogCommentLike.CommentID.EQ(Int32(commentID))))
	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Delete", slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		return fmt.Errorf("deleting blog comment like: %w", err)
	}
	return nil
}
