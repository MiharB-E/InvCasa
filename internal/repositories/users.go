package repositories

import (
    "context"

    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
    db *pgxpool.Pool
}

func (r *UserRepository) Create(ctx context.Context, user models.User) (models.User, error) {
    err := r.db.QueryRow(ctx, `
        INSERT INTO users (email, password, name, group_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, user.Email, user.Password, user.Name, user.GroupID).Scan(&user.ID)
    return user, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (models.User, error) {
    var user models.User
    err := r.db.QueryRow(ctx, `
        SELECT id, email, password, name, group_id
        FROM users
        WHERE email = $1
    `, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.GroupID)
    return user, err
}

func (r *UserRepository) UpdateGroup(ctx context.Context, userID, groupID int64) error {
    _, err := r.db.Exec(ctx, `UPDATE users SET group_id = $1 WHERE id = $2`, groupID, userID)
    return err
}