package repositories

import (
    "context"

    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/jackc/pgx/v5/pgxpool"
)

type PurchaseRepository struct {
    db *pgxpool.Pool
}

func (r *PurchaseRepository) Create(ctx context.Context, purchase models.Purchase) (models.Purchase, error) {
    err := r.db.QueryRow(ctx, `
        INSERT INTO purchases (product_id, user_id, quantity, price, store_name)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `, purchase.ProductID, purchase.UserID, purchase.Quantity, purchase.Price, purchase.StoreName).Scan(&purchase.ID)
    return purchase, err
}