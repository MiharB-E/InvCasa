package services

import (
    "github.com/MiharB-E/InvCasa/internal/repositories"
)

type Services struct {
    Auth     *AuthService
    Groups   *GroupService
    Products *ProductService
    Purchases *PurchaseService
}

func New(repos *repositories.Repositories, jwtSecret string) *Services {
    return &Services{
        Auth:     &AuthService{repos: repos, jwtSecret: jwtSecret},
        Groups:   &GroupService{repos: repos},
        Products: &ProductService{repos: repos},
        Purchases: &PurchaseService{repos: repos},
    }
}