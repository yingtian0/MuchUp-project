package postgres

import (
	"context"
	"fmt"

	"MuchUp/app/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	// pgxpool.NewWithConfigで接続数などの設定ができる
	// とりあえず繋げるためにNewを使う
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// 意思疎通の確認
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, err
}
