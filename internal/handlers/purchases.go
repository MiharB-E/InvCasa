package handlers

import (
    "net/http"

    "github.com/MiharB-E/InvCasa/internal/services"
)

type PurchasesHandler struct {
    service *services.PurchaseService
}

func NewPurchasesHandler(service *services.PurchaseService) *PurchasesHandler {
    return &PurchasesHandler{service: service}
}

func (h *PurchasesHandler) Create(w http.ResponseWriter, r *http.Request) {
    var body struct {
        ProductID int64   `json:"product_id"`
        Quantity  float64 `json:"quantity"`
        Price     float64 `json:"price"`
        StoreName string  `json:"store_name"`
    }
    if err := readJSON(r, &body); err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    userID := middleware.GetUserID(r.Context())
    purchase, err := h.service.Create(r.Context(), models.Purchase{
        ProductID: body.ProductID,
        UserID:    userID,
        Quantity:  body.Quantity,
        Price:     body.Price,
        StoreName: body.StoreName,
    })
    if err != nil {
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusCreated, purchase)
}