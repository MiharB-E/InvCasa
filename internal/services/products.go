package services

import (
    "context"

    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/MiharB-E/InvCasa/internal/repositories"
)

type ProductService struct {
    repos *repositories.Repositories
}

func (s *ProductService) List(ctx context.Context, groupID int64) ([]models.Product, error) {
    return s.repos.Products.ListByGroup(ctx, groupID)
}

func (s *ProductService) Create(ctx context.Context, product models.Product) (models.Product, error) {
    if product.Status == "" {
        product.Status = "ok"
    }
    return s.repos.Products.Create(ctx, product)
}

func (s *ProductService) MarkLow(ctx context.Context, productID int64) error {
    if err := s.repos.Products.UpdateStatus(ctx, productID, "low"); err != nil {
        return err
    }
    return s.repos.ShoppingList.Upsert(ctx, productID, "pending")
}

func (s *ProductService) UpdateFavorite(ctx context.Context, productID int64, isFavorite bool) error {
    return s.repos.Products.UpdateFavorite(ctx, productID, isFavorite)
}