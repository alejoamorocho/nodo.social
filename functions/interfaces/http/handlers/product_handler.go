package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/kha0sys/nodo.social/domain/dto"
    "github.com/kha0sys/nodo.social/services"
)

// ProductHandler maneja las peticiones HTTP relacionadas con productos
type ProductHandler struct {
    productService *services.ProductService
}

// NewProductHandler crea una nueva instancia de ProductHandler
func NewProductHandler(productService *services.ProductService) *ProductHandler {
    return &ProductHandler{
        productService: productService,
    }
}

// RegisterRoutes registra las rutas del handler en el router
func (h *ProductHandler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/products", h.CreateProduct).Methods("POST")
    r.HandleFunc("/products/{id}", h.GetProduct).Methods("GET")
    r.HandleFunc("/products/{id}", h.UpdateProduct).Methods("PUT")
    r.HandleFunc("/products/{id}", h.DeleteProduct).Methods("DELETE")
    r.HandleFunc("/products/{id}/approve", h.ApproveProduct).Methods("POST")
    r.HandleFunc("/nodes/{nodeId}/products", h.GetProductsByNode).Methods("GET")
}

// CreateProduct maneja la creación de un nuevo producto
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var productDTO dto.ProductDTO
    if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    product, err := h.productService.CreateProduct(r.Context(), &productDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(dto.FromProductModel(product))
}

// GetProduct maneja la obtención de un producto por ID
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    product, err := h.productService.GetProduct(r.Context(), vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(dto.FromProductModel(product))
}

// UpdateProduct maneja la actualización de un producto
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var productDTO dto.ProductDTO
    if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    productDTO.ID = vars["id"]
    product, err := h.productService.UpdateProduct(r.Context(), &productDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(dto.FromProductModel(product))
}

// DeleteProduct maneja la eliminación de un producto
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    if err := h.productService.DeleteProduct(r.Context(), vars["id"]); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// ApproveProduct maneja la aprobación de un producto
func (h *ProductHandler) ApproveProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    if err := h.productService.ApproveProduct(r.Context(), vars["id"]); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// GetProductsByNode maneja la obtención de productos por nodo
func (h *ProductHandler) GetProductsByNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    products, err := h.productService.GetProductsByNode(r.Context(), vars["nodeId"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var productsDTO []dto.ProductDTO
    for _, product := range products {
        productDTO := dto.FromProductModel(product)
        productsDTO = append(productsDTO, *productDTO)
    }

    json.NewEncoder(w).Encode(productsDTO)
}
