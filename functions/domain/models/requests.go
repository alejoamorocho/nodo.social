package models

import (
	"time"
	"github.com/kha0sys/nodo.social/functions/domain/models/contact"
)

// NodeData representa los datos para crear/actualizar un nodo
type NodeData struct {
    Type        NodeType   `json:"type"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Media       []MediaURL `json:"media"`
}

// FeedRequest representa los parámetros de la request para obtener el feed de nodos
type FeedRequest struct {
    Type      NodeType `json:"type"`      // Tipo de nodo a filtrar
    Category  string   `json:"category"`  // Categoría específica dentro del tipo
    UserID    string   `json:"userId"`    // Filtrar por usuario específico
    Following bool     `json:"following"` // Solo mostrar nodos que el usuario sigue
    Page      int      `json:"page"`      // Número de página para paginación
    PageSize  int      `json:"pageSize"`  // Tamaño de página para paginación
    LastID    string   `json:"lastId"`    // ID del último nodo para paginación por cursor
    Limit     int      `json:"limit"`     // Límite de resultados a retornar
}

// ProductData representa los datos para crear un producto
type ProductData struct {
    Name        string            `json:"name"`
    Description string            `json:"description"`
    Price       float64           `json:"price"`
    Images      []MediaURL        `json:"images"`
    Contact     contact.ContactInfo `json:"contact"`
    Percentage  int               `json:"percentage"`
    StoreID     string            `json:"storeId"`
}

// UserActivityFilters representa los filtros para la actividad del usuario
type UserActivityFilters struct {
    From          time.Time `json:"from" firestore:"from"`
    To            time.Time `json:"to" firestore:"to"`
    LastTimestamp int64     `json:"lastTimestamp" firestore:"lastTimestamp"`
    Limit         int       `json:"limit" firestore:"limit"`
    Type          string    `json:"type" firestore:"type"` // "nodes", "products", "follows"
}
