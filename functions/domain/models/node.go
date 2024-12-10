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
	ID string `firestore:"-" json:"id"`
	// Title es el título principal del nodo
	Title string `firestore:"title" json:"title"`
	// Description es una descripción detallada de la causa
	Description string `firestore:"description" json:"description"`
	// Type es el tipo de nodo (social, ambiental, animal)
	Type NodeType `firestore:"type" json:"type"`
	// UserID es el identificador del usuario que creó el nodo
	UserID string `firestore:"userId" json:"userId"`
	// Media es una lista de recursos multimedia asociados al nodo
	Media []string `firestore:"media" json:"media"`
	// Updates es una lista de actualizaciones sobre la causa
	Updates []Update `firestore:"updates" json:"updates"`
	// Followers es una lista de IDs de usuarios que siguen el nodo
	Followers []string `firestore:"followers" json:"followers"`
	// FollowersCount es el número total de seguidores
	FollowersCount int `firestore:"followersCount" json:"followersCount"`
	// LinkedProducts es una lista de IDs de productos asociados al nodo
	LinkedProducts []string `firestore:"linkedProducts" json:"linkedProducts"`
	// Products es una lista de IDs de productos asociados al nodo
	Products []string `firestore:"products" json:"products"`
	// Images es una lista de recursos multimedia asociados al nodo
	Images []string `firestore:"images" json:"images"`
	// ApprovalConfig contiene la configuración de aprobación para productos
	ApprovalConfig ApprovalConfig `firestore:"approvalConfig" json:"approvalConfig"`
	// Metrics contiene las métricas de interacción del nodo
	Metrics InteractionMetrics `firestore:"metrics" json:"metrics"`
	// CreatedAt es la fecha de creación del nodo
	CreatedAt time.Time `firestore:"createdAt" json:"createdAt"`
	// UpdatedAt es la fecha de última actualización del nodo
	UpdatedAt time.Time `firestore:"updatedAt" json:"updatedAt"`
	// Name es el nombre descriptivo del nodo
	Name string `firestore:"name" json:"name"`
	// Tags es una lista de etiquetas asociadas al nodo
	Tags []string `firestore:"tags,omitempty" json:"tags,omitempty"`
	// Status es el estado actual del nodo
	Status string `firestore:"status" json:"status"`
}

// BeforeCreate prepara el nodo para ser creado inicializando campos requeridos.
// Esta función debe ser llamada antes de crear un nuevo nodo en la base de datos.
func (n *Node) BeforeCreate() {
	now := time.Now()
	n.CreatedAt = now
	n.UpdatedAt = now

	// Inicializar slices vacíos
	if n.Media == nil {
		n.Media = make([]string, 0)
	}
	if n.Updates == nil {
		n.Updates = make([]Update, 0)
	}
	if n.Followers == nil {
		n.Followers = make([]string, 0)
	}
	if n.LinkedProducts == nil {
		n.LinkedProducts = make([]string, 0)
	}
	if n.Products == nil {
		n.Products = make([]string, 0)
	}
	if n.Images == nil {
		n.Images = make([]string, 0)
	}

	// Inicializar contadores
	n.FollowersCount = 0
	n.Metrics = InteractionMetrics{
		Views:     0,
		Likes:     0,
		Shares:    0,
		Comments:  0,
		Followers: 0,
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
	ID          string    `json:"id" firestore:"id"`
	Title       string    `json:"title" firestore:"title"`
	Description string    `json:"description" firestore:"description"`
	Media       []string  `json:"media" firestore:"media"`
	CreatedAt   time.Time `json:"createdAt" firestore:"createdAt"`
}

// ApprovalConfig define la configuración de aprobación para productos en un nodo.
type ApprovalConfig struct {
	RequiresApproval bool `json:"requiresApproval" firestore:"requiresApproval"`
	AutoApprove      bool `json:"autoApprove" firestore:"autoApprove"`
}

// InteractionMetrics representa las métricas de interacción de un nodo.
type InteractionMetrics struct {
	Views     int `json:"views" firestore:"views"`
	Likes     int `json:"likes" firestore:"likes"`
	Shares    int `json:"shares" firestore:"shares"`
	Comments  int `json:"comments" firestore:"comments"`
	Followers int `json:"followers" firestore:"followers"`
}
