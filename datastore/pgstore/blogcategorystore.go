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

var _ datastore.BlogCategoryProvider = BlogCategoryStore{}

type BlogCategoryStore struct {
	db  *stdsql.DB
	log *slog.Logger
}

func NewBlogCategoryStore(db *stdsql.DB, log *slog.Logger) *BlogCategoryStore {
	return &BlogCategoryStore{
		db:  db,
		log: log.With(slog.String("DataStore", "BlogCategoryStore")),
	}
}

func (b BlogCategoryStore) Get(ctx context.Context, id int32) (*model.BlogCategory, error) {
	var category *model.BlogCategory
	stmt := SELECT(BlogCategory.AllColumns).
		FROM(BlogCategory).
		WHERE(BlogCategory.ID.EQ(Int32(id)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Get", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, category)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			b.log.InfoContext(ctx, "not found", slog.Int("id", int(id)))
			return nil, nil
		}
		return nil, fmt.Errorf("getting blog category %d: %w", id, err)
	}
	return category, nil
}

func (b BlogCategoryStore) All(ctx context.Context) ([]model.BlogCategory, error) {
	var categories []model.BlogCategory
	stmt := SELECT(BlogCategory.AllColumns).
		FROM(BlogCategory).
		ORDER_BY(BlogCategory.CategoryName.ASC())

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "All", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, &categories)
	if err != nil {
		return nil, fmt.Errorf("getting all blog categories: %w", err)
	}
	return categories, nil
}

func (b BlogCategoryStore) Create(ctx context.Context, blogCategory model.BlogCategory) (*model.BlogCategory, error) {
	stmt := BlogCategory.INSERT(BlogCategory.AllColumns).
		MODEL(&blogCategory).
		RETURNING(BlogCategory.AllColumns)

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, b.db, &blogCategory)
	if err != nil {
		return nil, fmt.Errorf("creating blog category: %w", err)
	}
	return &blogCategory, nil
}

func (b BlogCategoryStore) Update(ctx context.Context, blogCategory model.BlogCategory) (*model.BlogCategory, error) {
	stmt := BlogCategory.UPDATE(BlogCategory.AllColumns).
		MODEL(&blogCategory).
		WHERE(BlogCategory.ID.EQ(Int32(blogCategory.ID)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Update", slog.String("sql", stmt.DebugSql()))
	}

	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		return nil, fmt.Errorf("updating blog category: %w", err)
	}
	return &blogCategory, nil
}

func (b BlogCategoryStore) Delete(ctx context.Context, id model.BlogCategory) error {
	stmt := BlogCategory.DELETE().
		WHERE(BlogCategory.ID.EQ(Int32(id.ID)))

	if b.log.Enabled(ctx, slog.LevelDebug) {
		b.log.DebugContext(ctx, "Delete", slog.String("sql", stmt.DebugSql()))
	}

	_, err := stmt.ExecContext(ctx, b.db)
	if err != nil {
		return fmt.Errorf("deleting blog category %d: %w", id.ID, err)
	}
	return nil
}
