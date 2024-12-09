package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/kha0sys/nodo.social/domain/dto"
    "github.com/kha0sys/nodo.social/services"
)

// UserHandler maneja las peticiones HTTP relacionadas con usuarios
type UserHandler struct {
    userService *services.UserService
}

// NewUserHandler crea una nueva instancia de UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

// CreateUser maneja la creación de un nuevo usuario
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var userDTO dto.UserDTO
    if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := h.userService.CreateUser(r.Context(), userDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(dto.FromUserModel(user))
}

// GetUser maneja la obtención de un usuario por ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user, err := h.userService.GetUser(r.Context(), vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(dto.FromUserModel(user))
}

// UpdateUser maneja la actualización de un usuario
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var userDTO dto.UserDTO
    if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user := userDTO.ToModel()
    user.ID = vars["id"]

    if err := h.userService.UpdateUser(r.Context(), user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// DeleteUser maneja la eliminación de un usuario
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    if err := h.userService.DeleteUser(r.Context(), vars["id"]); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
