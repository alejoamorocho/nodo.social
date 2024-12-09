package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/kha0sys/nodo.social/domain/models"
    "github.com/kha0sys/nodo.social/domain/dto"
    "github.com/kha0sys/nodo.social/services"
)

// StoreHandler maneja las peticiones HTTP relacionadas con tiendas
type StoreHandler struct {
    storeService *services.StoreService
}

// NewStoreHandler crea una nueva instancia de StoreHandler
func NewStoreHandler(storeService *services.StoreService) *StoreHandler {
    return &StoreHandler{
        storeService: storeService,
    }
}

// RegisterRoutes registra las rutas del handler en el router
func (h *StoreHandler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/stores", h.CreateStore).Methods("POST")
    r.HandleFunc("/stores/{id}", h.GetStore).Methods("GET")
    r.HandleFunc("/stores/{id}", h.UpdateStore).Methods("PUT")
    r.HandleFunc("/stores/{id}", h.DeleteStore).Methods("DELETE")
}

// CreateStore maneja la creación de una nueva tienda
func (h *StoreHandler) CreateStore(w http.ResponseWriter, r *http.Request) {
    var storeDTO dto.StoreDTO
    if err := json.NewDecoder(r.Body).Decode(&storeDTO); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    store, err := h.storeService.CreateStore(r.Context(), storeDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(dto.FromStoreModel(store))
}

// GetStore maneja la obtención de una tienda por ID
func (h *StoreHandler) GetStore(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    store, err := h.storeService.GetStore(r.Context(), vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(store)
}

// UpdateStore maneja la actualización de una tienda
func (h *StoreHandler) UpdateStore(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var store models.Store
    if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    store.ID = vars["id"]
    if err := h.storeService.UpdateStore(r.Context(), &store); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(store)
}

// DeleteStore maneja la eliminación de una tienda
func (h *StoreHandler) DeleteStore(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    if err := h.storeService.DeleteStore(r.Context(), vars["id"]); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
