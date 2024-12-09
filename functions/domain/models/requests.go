package models

// NodeData representa los datos para crear/actualizar un nodo
type NodeData struct {
    Type        NodeType   `json:"type"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Media       []MediaURL `json:"media"`
}

// FeedFilters representa los filtros para el feed de nodos
type FeedFilters struct {
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
    Name        string     `json:"name"`
    Description string     `json:"description"`
    Price       float64    `json:"price"`
    Images      []MediaURL `json:"images"`
    Contact     ContactInfo `json:"contact"`
    Percentage  int        `json:"percentage"`
    StoreID     string     `json:"storeId"`
}

// UserActivityFilters representa los filtros para la actividad del usuario
type UserActivityFilters struct {
    LastTimestamp int64  `json:"lastTimestamp"`
    Limit         int    `json:"limit"`
    Type          string `json:"type"` // "nodes", "products", "follows"
}
