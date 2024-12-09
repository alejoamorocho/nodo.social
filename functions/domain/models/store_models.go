// Package models define las estructuras de datos principales de la aplicación
package models

import (
	"fmt"
	"time"
)

// MediaURL representa una URL multimedia con metadatos
type MediaURL struct {
	URL         string `json:"url" firestore:"url"`               // URL del recurso multimedia
	Type        string `json:"type" firestore:"type"`             // Tipo de medio (image, video, etc.)
	Description string `json:"description" firestore:"description"` // Descripción opcional del recurso
	Thumbnail   string `json:"thumbnail,omitempty" firestore:"thumbnail,omitempty"` // URL de la miniatura (opcional)
}

// ContactInfo representa la información de contacto
type ContactInfo struct {
	Email     string   `json:"email" firestore:"email"`
	Phone     string   `json:"phone" firestore:"phone"`
	Address   string   `json:"address" firestore:"address"`
	Website   string   `json:"website" firestore:"website"`
	Instagram string   `json:"instagram" firestore:"instagram"`
	Twitter   string   `json:"twitter" firestore:"twitter"`
	Facebook  string   `json:"facebook" firestore:"facebook"`
	Social    []string `json:"social" firestore:"social"`
}

// Validate verifica que los campos obligatorios de ContactInfo estén presentes
func (c *ContactInfo) Validate() error {
	if c.Email == "" && c.Phone == "" {
		return fmt.Errorf("se requiere al menos un método de contacto (email o teléfono)")
	}
	return nil
}

// Store representa una tienda en el sistema.
// Una tienda es una entidad que puede vender productos asociados a nodos sociales.
// Cada tienda tiene un usuario propietario y puede tener múltiples productos.
type Store struct {
	// ID es el identificador único de la tienda
	ID string `json:"id" firestore:"id"`
	// UserID es el ID del usuario propietario de la tienda
	UserID string `json:"userId" firestore:"userId"`
	// Name es el nombre comercial de la tienda
	Name string `json:"name" firestore:"name"`
	// Description es una descripción detallada de la tienda y sus productos
	Description string `json:"description" firestore:"description"`
	// Contact contiene la información de contacto de la tienda
	Contact ContactInfo `json:"contact" firestore:"contact"`
	// Logo es la URL del logo de la tienda
	Logo string `json:"logo" firestore:"logo"`
	// Products es una lista de IDs de productos asociados a la tienda
	Products []string `json:"products" firestore:"products"`
	// CreatedAt es la fecha de creación de la tienda
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
	// UpdatedAt es la última fecha de actualización de la tienda
	UpdatedAt time.Time `json:"updatedAt" firestore:"updatedAt"`
}

// Validate verifica que los campos obligatorios de la tienda estén presentes y sean válidos.
func (s *Store) Validate() error {
	if s.UserID == "" {
		return fmt.Errorf("el ID de usuario es obligatorio")
	}
	if s.Name == "" {
		return fmt.Errorf("el nombre de la tienda es obligatorio")
	}
	if s.Description == "" {
		return fmt.Errorf("la descripción de la tienda es obligatoria")
	}
	return s.Contact.Validate()
}

// BeforeCreate prepara la tienda para ser creada inicializando campos requeridos.
func (s *Store) BeforeCreate() {
	now := time.Now()
	s.CreatedAt = now
	s.UpdatedAt = now
	if s.Products == nil {
		s.Products = make([]string, 0)
	}
}

// BeforeUpdate actualiza la fecha de modificación de la tienda.
func (s *Store) BeforeUpdate() {
	s.UpdatedAt = time.Now()
}

// Product representa un producto en el sistema.
// Los productos están asociados a tiendas y nodos sociales.
// Los productos requieren aprobación antes de ser visibles en el sistema.
type Product struct {
	// ID es el identificador único del producto
	ID string `json:"id" firestore:"id"`
	// StoreID es el ID de la tienda a la que pertenece el producto
	StoreID string `json:"storeId" firestore:"storeId"`
	// NodeID es el ID del nodo social al que está vinculado el producto
	NodeID string `json:"nodeId" firestore:"nodeId"`
	// Name es el nombre del producto
	Name string `json:"name" firestore:"name"`
	// Description es una descripción detallada del producto
	Description string `json:"description" firestore:"description"`
	// Price es el precio del producto en la moneda predeterminada
	Price float64 `json:"price" firestore:"price"`
	// Images es una lista de URLs de imágenes del producto
	Images []string `json:"images" firestore:"images"`
	// ContactInfo contiene la información de contacto específica para este producto
	ContactInfo ContactInfo `json:"contactInfo" firestore:"contactInfo"`
	// DonationPercent es el porcentaje del precio que se donará a la causa
	DonationPercent int `json:"donationPercent" firestore:"donationPercent"`
	// ApprovalStatus indica el estado de aprobación del producto (pending, approved, rejected)
	ApprovalStatus string `json:"approvalStatus" firestore:"approvalStatus"`
	// CreatedAt es la fecha de creación del producto
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
	// UpdatedAt es la última fecha de actualización del producto
	UpdatedAt time.Time `json:"updatedAt" firestore:"updatedAt"`
}

// Validate verifica que los campos obligatorios del producto estén presentes y sean válidos.
func (p *Product) Validate() error {
	if p.StoreID == "" {
		return fmt.Errorf("el ID de la tienda es obligatorio")
	}
	if p.NodeID == "" {
		return fmt.Errorf("el ID del nodo es obligatorio")
	}
	if p.Name == "" {
		return fmt.Errorf("el nombre del producto es obligatorio")
	}
	if p.Description == "" {
		return fmt.Errorf("la descripción del producto es obligatoria")
	}
	if p.Price <= 0 {
		return fmt.Errorf("el precio debe ser mayor a 0")
	}
	if p.DonationPercent < 0 || p.DonationPercent > 100 {
		return fmt.Errorf("el porcentaje de donación debe estar entre 0 y 100")
	}
	return p.ContactInfo.Validate()
}

// BeforeCreate prepara el producto para ser creado inicializando campos requeridos.
// Establece las fechas de creación y actualización, inicializa las imágenes
// y establece el estado de aprobación inicial como "pending".
func (p *Product) BeforeCreate() {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	if p.Images == nil {
		p.Images = make([]string, 0)
	}
	if p.ApprovalStatus == "" {
		p.ApprovalStatus = "pending"
	}
}

// BeforeUpdate actualiza la fecha de modificación del producto.
func (p *Product) BeforeUpdate() {
	p.UpdatedAt = time.Now()
}
