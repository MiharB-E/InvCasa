package repositories

import (
    "context"

    "github.com/jackc/pgx/v5/pgxpool"
)

type ShoppingListRepository struct {
    db *pgxpool.Pool
}

func (r *ShoppingListRepository) Upsert(ctx context.Context, productID int64, status string) error {
    _, err := r.db.Exec(ctx, `
        INSERT INTO shopping_list_items (product_id, status)
        VALUES ($1, $2)
        ON CONFLICT (product_id) DO UPDATE SET status = EXCLUDED.status
    `, productID, status)
    return err
}