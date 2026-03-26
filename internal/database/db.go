package database

import (
    "context"

    "github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
    cfg, err := pgxpool.ParseConfig(dbURL)
    if err != nil {
        return nil, err
    }
    return pgxpool.NewWithConfig(ctx, cfg)
}