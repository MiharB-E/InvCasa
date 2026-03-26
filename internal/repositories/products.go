package repositories

import (
    "context"

    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
    db *pgxpool.Pool
}

func (r *ProductRepository) ListByGroup(ctx context.Context, groupID int64) ([]models.Product, error) {
    rows, err := r.db.Query(ctx, `
        SELECT id, name, category_id, quantity, unit, status, group_id, is_favorite
        FROM products
        WHERE group_id = $1
        ORDER BY id DESC
    `, groupID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []models.Product
    for rows.Next() {
        var p models.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Quantity, &p.Unit, &p.Status, &p.GroupID, &p.IsFavorite); err != nil {
            return nil, err
        }
        items = append(items, p)
    }
    return items, rows.Err()
}

func (r *ProductRepository) Create(ctx context.Context, product models.Product) (models.Product, error) {
    err := r.db.QueryRow(ctx, `
        INSERT INTO products (name, category_id, quantity, unit, status, group_id, is_favorite)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `, product.Name, product.CategoryID, product.Quantity, product.Unit, product.Status, product.GroupID, product.IsFavorite).Scan(&product.ID)
    return product, err
}

func (r *ProductRepository) UpdateStatus(ctx context.Context, productID int64, status string) error {
    _, err := r.db.Exec(ctx, `UPDATE products SET status = $1 WHERE id = $2`, status, productID)
    return err
}

func (r *ProductRepository) UpdateFavorite(ctx context.Context, productID int64, isFavorite bool) error {
    _, err := r.db.Exec(ctx, `UPDATE products SET is_favorite = $1 WHERE id = $2`, isFavorite, productID)
    return err
}