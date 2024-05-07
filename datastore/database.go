package datastore

import (
	"context"
	"fmt"
	"github.com/PaluMacil/barkdognet/configuration"
	"github.com/PaluMacil/barkdognet/datastore/pgstore"
	"github.com/PaluMacil/barkdognet/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"time"
)

func NewDatabase(ctx context.Context, configProvider configuration.Provider) (*Database, error) {
	config, err := configProvider.Config()
	if err != nil {
		return nil, fmt.Errorf("getting database configuration: %w", err)
	}

	dsn := config.Database.ConnectionString()
	connConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	connConfig.MaxConns = 10
	connConfig.MaxConnIdleTime = time.Minute
	connConfig.MaxConnLifetimeJitter = 5 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDBFromPool(pool)
	log := logger.NewLogger(config.Env)

	database := &Database{
		Users:          pgstore.NewUserStore(db, log),
		Roles:          pgstore.NewRoleStore(db, log),
		Sessions:       pgstore.NewSessionStore(db, log),
		BlogCategories: pgstore.NewBlogCategoryStore(db, log),
		BlogPosts:      pgstore.NewBlogPostStore(db, log),
		BlogComments:   pgstore.NewBlogCommentStore(db, log),
	}

	return database, nil
}

type Database struct {
	Users          UserProvider
	Roles          RoleProvider
	Sessions       SessionProvider
	BlogCategories BlogCategoryProvider
	BlogPosts      BlogPostProvider
	BlogComments   BlogCommentProvider
}
