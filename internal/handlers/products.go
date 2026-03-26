package handlers

import (
    "net/http"
    "strconv"

    "github.com/MiharB-E/InvCasa/internal/middleware"
    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/MiharB-E/InvCasa/internal/services"
    "github.com/go-chi/chi/v5"
)

type ProductsHandler struct {
    service *services.ProductService
}

func NewProductsHandler(service *services.ProductService) *ProductsHandler {
    return &ProductsHandler{service: service}
}

type productRequest struct {
    Name       string  `json:"name"`
    CategoryID *int64  `json:"category_id"`
    Quantity   float64 `json:"quantity"`
    Unit       string  `json:"unit"`
    Status     string  `json:"status"`
    IsFavorite bool    `json:"is_favorite"`
}

func (h *ProductsHandler) List(w http.ResponseWriter, r *http.Request) {
    groupID := middleware.GetGroupID(r.Context())
    items, err := h.service.List(r.Context(), groupID)
    if err != nil {
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    writeJSON(w, http.StatusOK, items)
}

func (h *ProductsHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req productRequest
    if err := readJSON(r, &req); err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    groupID := middleware.GetGroupID(r.Context())
    product, err := h.service.Create(r.Context(), models.Product{
        Name:       req.Name,
        CategoryID: req.CategoryID,
        Quantity:   req.Quantity,
        Unit:       req.Unit,
        Status:     req.Status,
        GroupID:    groupID,
        IsFavorite: req.IsFavorite,
    })
    if err != nil {
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusCreated, product)
}

func (h *ProductsHandler) MarkLow(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id inválido"})
        return
    }
    if err := h.service.MarkLow(r.Context(), id); err != nil {
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *ProductsHandler) UpdateFavorite(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id inválido"})
        return
    }
    var body struct {
        IsFavorite bool `json:"is_favorite"`
    }
    if err := readJSON(r, &body); err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }
    if err := h.service.UpdateFavorite(r.Context(), id, body.IsFavorite); err != nil {
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}