package repositories

import (
    "context"

    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepository struct {
    db *pgxpool.Pool
}

func (r *GroupRepository) Create(ctx context.Context, group models.Group) (models.Group, error) {
    err := r.db.QueryRow(ctx, `
        INSERT INTO groups (name, invite_code)
        VALUES ($1, $2)
        RETURNING id
    `, group.Name, group.InviteCode).Scan(&group.ID)
    return group, err
}

func (r *GroupRepository) GetByInviteCode(ctx context.Context, code string) (models.Group, error) {
    var group models.Group
    err := r.db.QueryRow(ctx, `
        SELECT id, name, invite_code
        FROM groups
        WHERE invite_code = $1
    `, code).Scan(&group.ID, &group.Name, &group.InviteCode)
    return group, err
}

func (r *GroupRepository) GetByID(ctx context.Context, id int64) (models.Group, error) {
    var group models.Group
    err := r.db.QueryRow(ctx, `
        SELECT id, name, invite_code
        FROM groups
        WHERE id = $1
    `, id).Scan(&group.ID, &group.Name, &group.InviteCode)
    return group, err
}