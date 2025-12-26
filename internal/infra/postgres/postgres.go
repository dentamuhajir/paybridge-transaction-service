package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(dsn string) (*pgxpool.Pool, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
