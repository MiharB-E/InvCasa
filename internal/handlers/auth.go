package handlers

import (
    "net/http"

    "github.com/MiharB-E/InvCasa/internal/services"
)

type AuthHandler struct {
    service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
    return &AuthHandler{service: service}
}

type registerRequest struct {
    Email      string `json:"email"`
    Password   string `json:"password"`
    Name       string `json:"name"`
    InviteCode string `json:"invite_code"`
    GroupName  string `json:"group_name"`
}

type authResponse struct {
    Token string `json:"token"`
    User  interface{} `json:"user"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req registerRequest
    if err := readJSON(r, &req); err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    user, token, err := h.service.Register(r.Context(), services.RegisterInput{
        Email:      req.Email,
        Password:   req.Password,
        Name:       req.Name,
        InviteCode: req.InviteCode,
        GroupName:  req.GroupName,
    })
    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusCreated, authResponse{Token: token, User: user})
}

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req loginRequest
    if err := readJSON(r, &req); err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    user, token, err := h.service.Login(r.Context(), req.Email, req.Password)
    if err != nil {
        writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusOK, authResponse{Token: token, User: user})
}