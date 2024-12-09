// Package models define las estructuras de datos principales de la aplicación
package models

import (
	"time"
)

// NodeType representa el tipo de nodo social
type NodeType string

// Constantes que definen los tipos de nodos disponibles
const (
	// Social representa un nodo de impacto social
	Social NodeType = "social"
	// Environmental representa un nodo de impacto ambiental
	Environmental NodeType = "environmental"
	// Animal representa un nodo de protección animal
	Animal NodeType = "animal"
)

// Node representa un nodo en el sistema.
// Un nodo es una entidad que representa una causa social, ambiental o animal.
// Los nodos pueden tener seguidores, productos asociados y actualizaciones.
type Node struct {
	// ID es el identificador único del nodo
	ID string `json:"id" firestore:"id"`
	// Name es el nombre descriptivo del nodo
	Name string `json:"name" firestore:"name"`
	// Type indica el tipo de causa que representa el nodo
	Type NodeType `json:"type" firestore:"type"`
	// Title es el título principal del nodo
	Title string `json:"title" firestore:"title"`
	// Description es una descripción detallada de la causa
	Description string `json:"description" firestore:"description"`
	// UserID es el identificador del usuario que creó el nodo
	UserID string `json:"userId" firestore:"userId"`
	// Media es una lista de recursos multimedia asociados al nodo
	Media []string `json:"media" firestore:"media"`
	// Updates es una lista de actualizaciones sobre la causa
	Updates []Update `json:"updates" firestore:"updates"`
	// Followers es una lista de IDs de usuarios que siguen el nodo
	Followers []string `json:"followers" firestore:"followers"`
	// FollowersCount es el número total de seguidores
	FollowersCount int `json:"followersCount" firestore:"followersCount"`
	// LinkedProducts es una lista de IDs de productos asociados al nodo
	LinkedProducts []string `json:"linkedProducts" firestore:"linkedProducts"`
	// Products es una lista de IDs de productos asociados al nodo
	Products []string `json:"products" firestore:"products"`
	// Images es una lista de recursos multimedia asociados al nodo
	Images []string `json:"images" firestore:"images"`
	// ApprovalConfig contiene la configuración de aprobación para productos
	ApprovalConfig ApprovalConfig `json:"approvalConfig" firestore:"approvalConfig"`
	// Metrics contiene las métricas de interacción del nodo
	Metrics InteractionMetrics `json:"metrics" firestore:"metrics"`
	// CreatedAt es la fecha de creación del nodo
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
	// UpdatedAt es la última fecha de actualización del nodo
	UpdatedAt time.Time `json:"updatedAt" firestore:"updatedAt"`
}

// BeforeCreate prepara el nodo para ser creado inicializando campos requeridos.
// Esta función debe ser llamada antes de crear un nuevo nodo en la base de datos.
func (n *Node) BeforeCreate() {
	now := time.Now()
	n.CreatedAt = now
	n.UpdatedAt = now
	if n.Media == nil {
		n.Media = make([]string, 0)
	}
	if n.Updates == nil {
		n.Updates = make([]Update, 0)
	}
	if n.Followers == nil {
		n.Followers = make([]string, 0)
	}
	if n.Products == nil {
		n.Products = make([]string, 0)
	}
	if n.LinkedProducts == nil {
		n.LinkedProducts = make([]string, 0)
	}
	if n.Images == nil {
		n.Images = make([]string, 0)
	}
	n.FollowersCount = 0
	n.Metrics = InteractionMetrics{
		Views:    0,
		Likes:    0,
		Shares:   0,
		Comments: 0,
	}
}

// BeforeUpdate actualiza la fecha de modificación del nodo.
// Esta función debe ser llamada antes de actualizar un nodo en la base de datos.
func (n *Node) BeforeUpdate() {
	n.UpdatedAt = time.Now()
}

// Update representa una actualización de un nodo.
// Las actualizaciones son usadas para mantener informados a los seguidores sobre el progreso de la causa.
type Update struct {
	// ID es el identificador único de la actualización
	ID string `json:"id" firestore:"id"`
	// Title es el título de la actualización
	Title string `json:"title" firestore:"title"`
	// Description es el contenido detallado de la actualización
	Description string `json:"description" firestore:"description"`
	// Media es una lista de recursos multimedia asociados a la actualización
	Media []string `json:"media" firestore:"media"`
	// CreatedAt es la fecha de creación de la actualización
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
}

// ApprovalConfig define la configuración de aprobación para productos en un nodo.
type ApprovalConfig struct {
	// RequiresApproval indica si los productos requieren aprobación manual
	RequiresApproval bool `json:"requiresApproval" firestore:"requiresApproval"`
	// AutoApprove indica si los productos se aprueban automáticamente
	AutoApprove bool `json:"autoApprove" firestore:"autoApprove"`
}

// InteractionMetrics representa las métricas de interacción de un nodo.
type InteractionMetrics struct {
	Views    int `json:"views" firestore:"views"`
	Likes    int `json:"likes" firestore:"likes"`
	Shares   int `json:"shares" firestore:"shares"`
	Comments int `json:"comments" firestore:"comments"`
}
