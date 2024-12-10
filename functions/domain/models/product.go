package models

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kha0sys/nodo.social/functions/domain/models/contact"
)

// Product representa un producto en el sistema.
// Los productos están asociados a tiendas y nodos sociales.
// Los productos requieren aprobación antes de ser visibles en el sistema.
type Product struct {
	ID              string              `json:"id" firestore:"id"`               // ID único del producto
	StoreID         string              `json:"storeId" firestore:"storeId"`     // ID de la tienda asociada
	NodeID          string              `json:"nodeId" firestore:"nodeId"`       // ID del nodo social asociado
	UserID          string              `json:"userId" firestore:"userId"`       // ID del usuario que creó el producto
	Name            string              `json:"name" firestore:"name"`           // Nombre del producto
	Description     string              `json:"description" firestore:"description"` // Descripción detallada
	Price           float64             `json:"price" firestore:"price"`         // Precio en la moneda predeterminada
	Images          []string            `json:"images" firestore:"images"`       // URLs de las imágenes
	Contact         contact.ContactInfo `json:"contact" firestore:"contact"`     // Información de contacto
	DonationPercent int                `json:"donationPercent" firestore:"donationPercent"` // Porcentaje para donación (1-100)
	ApprovalStatus  string             `json:"approvalStatus" firestore:"approvalStatus"`   // Estado de aprobación
	Status          string             `json:"status" firestore:"status"`       // Estado del producto
	CreatedAt       time.Time          `json:"createdAt" firestore:"createdAt"` // Fecha de creación
	UpdatedAt       time.Time          `json:"updatedAt" firestore:"updatedAt"` // Última actualización
}

// BeforeCreate prepara el producto para ser creado
func (p *Product) BeforeCreate() {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	if p.Images == nil {
		p.Images = make([]string, 0)
	}
	if p.Status == "" {
		p.Status = ProductStatusPending
	}
	if p.ApprovalStatus == "" {
		p.ApprovalStatus = "pending"
	}
}

// BeforeUpdate actualiza la fecha de modificación del producto
func (p *Product) BeforeUpdate() {
	p.UpdatedAt = time.Now()
}

// Validate verifica que los campos obligatorios del producto estén presentes y sean válidos
func (p *Product) Validate() error {
	if p.StoreID == "" {
		return fmt.Errorf("el ID de la tienda es obligatorio")
	}
	if p.NodeID == "" {
		return fmt.Errorf("el ID del nodo es obligatorio")
	}
	if p.UserID == "" {
		return fmt.Errorf("el ID del usuario es obligatorio")
	}
	if p.Name == "" {
		return fmt.Errorf("el nombre del producto es obligatorio")
	}
	if len(p.Description) < 10 {
		return fmt.Errorf("la descripción debe tener al menos 10 caracteres")
	}
	if p.Price <= 0 {
		return fmt.Errorf("el precio debe ser mayor a 0")
	}
	if p.DonationPercent < 1 || p.DonationPercent > 100 {
		return fmt.Errorf("el porcentaje de donación debe estar entre 1 y 100")
	}
	if len(p.Images) == 0 {
		return fmt.Errorf("debe proporcionar al menos una imagen")
	}
	if len(p.Images) > 5 {
		return fmt.Errorf("no se pueden agregar más de 5 imágenes")
	}

	// Validar URLs de imágenes
	for _, imageURL := range p.Images {
		if _, err := url.ParseRequestURI(imageURL); err != nil {
			return fmt.Errorf("URL de imagen inválida: %s", imageURL)
		}
	}

	// Validar contacto
	if err := p.Contact.Validate(); err != nil {
		return fmt.Errorf("información de contacto inválida: %v", err)
	}

	return nil
}

// ProductStatus define los posibles estados de un producto
const (
	ProductStatusPending  = "pending"  // Producto pendiente de revisión
	ProductStatusActive   = "active"   // Producto activo y visible
	ProductStatusInactive = "inactive" // Producto inactivo temporalmente
	ProductStatusRejected = "rejected" // Producto rechazado
)

