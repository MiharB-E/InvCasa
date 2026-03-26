package services

import (
    "context"

    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/MiharB-E/InvCasa/internal/repositories"
)

type PurchaseService struct {
    repos *repositories.Repositories
}

func (s *PurchaseService) Create(ctx context.Context, purchase models.Purchase) (models.Purchase, error) {
    return s.repos.Purchases.Create(ctx, purchase)
}