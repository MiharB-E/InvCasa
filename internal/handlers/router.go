package handlers

import (
    "net/http"

    "github.com/MiharB-E/InvCasa/internal/middleware"
    "github.com/MiharB-E/InvCasa/internal/services"
    "github.com/go-chi/chi/v5"
    chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(services *services.Services, jwtSecret string) http.Handler {
    r := chi.NewRouter()
    r.Use(chimiddleware.RequestID)
    r.Use(chimiddleware.RealIP)
    r.Use(chimiddleware.Logger)
    r.Use(chimiddleware.Recoverer)

    authHandler := NewAuthHandler(services.Auth)
    productsHandler := NewProductsHandler(services.Products)
    groupsHandler := NewGroupsHandler(services.Groups)
    purchasesHandler := NewPurchasesHandler(services.Purchases)

    r.Post("/register", authHandler.Register)
    r.Post("/login", authHandler.Login)

    r.Route("/", func(r chi.Router) {
        r.Use(middleware.JWTAuth(jwtSecret))

        r.Get("/products", productsHandler.List)
        r.Post("/products", productsHandler.Create)
        r.Patch("/products/{id}/low", productsHandler.MarkLow)
        r.Patch("/products/{id}/favorite", productsHandler.UpdateFavorite)

        r.Post("/groups", groupsHandler.Create)
        r.Post("/groups/join", groupsHandler.Join)
        r.Get("/groups/me", groupsHandler.GetMe)

        r.Post("/purchases", purchasesHandler.Create)
    })

    return r
}