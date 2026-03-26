package handlers

import (
    "net/http"

    "github.com/MiharB-E/InvCasa/internal/middleware"
    "github.com/MiharB-E/InvCasa/internal/services"
)

type GroupsHandler struct {
    service *services.GroupService
}

func NewGroupsHandler(service *services.GroupService) *GroupsHandler {
    return &GroupsHandler{service: service}
}

func (h *GroupsHandler) Create(w http.ResponseWriter, r *http.Request) {
    var body struct {
        Name string `json:"name"`
    }
    if err := readJSON(r, &body); err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }
    group, err := h.service.Create(r.Context(), body.Name)
    if err != nil {
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    writeJSON(w, http.StatusCreated, group)
}

func (h *GroupsHandler) Join(w http.ResponseWriter, r *http.Request) {
    var body struct {
        InviteCode string `json:"invite_code"`
    }
    if err := readJSON(r, &body); err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    userID := middleware.GetUserID(r.Context())
    group, err := h.service.Join(r.Context(), userID, body.InviteCode)
    if err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
        return
    }

    writeJSON(w, http.StatusOK, group)
}

func (h *GroupsHandler) GetMe(w http.ResponseWriter, r *http.Request) {
    groupID := middleware.GetGroupID(r.Context())
    group, err := h.service.Get(r.Context(), groupID)
    if err != nil {
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    writeJSON(w, http.StatusOK, group)
}