package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
    Users        *UserRepository
    Groups       *GroupRepository
    Products     *ProductRepository
    Purchases    *PurchaseRepository
    ShoppingList *ShoppingListRepository
}

func New(pool *pgxpool.Pool) *Repositories {
    return &Repositories{
        Users:        &UserRepository{db: pool},
        Groups:       &GroupRepository{db: pool},
        Products:     &ProductRepository{db: pool},
        Purchases:    &PurchaseRepository{db: pool},
        ShoppingList: &ShoppingListRepository{db: pool},
    }
}