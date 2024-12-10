// Package models define las estructuras de datos principales de la aplicación
package models

import (
	"fmt"
	"time"

	"github.com/kha0sys/nodo.social/functions/domain/models/contact"
)

// MediaURL representa una URL multimedia con metadatos
type MediaURL struct {
	URL         string `json:"url" firestore:"url"`               // URL del recurso multimedia
	Type        string `json:"type" firestore:"type"`             // Tipo de medio (image, video, etc.)
	Description string `json:"description" firestore:"description"` // Descripción opcional del recurso
	Thumbnail   string `json:"thumbnail,omitempty" firestore:"thumbnail,omitempty"` // URL de la miniatura (opcional)
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
	Contact contact.ContactInfo `json:"contact" firestore:"contact"`
	// Logo es la URL del logo de la tienda
	Logo string `json:"logo" firestore:"logo"`
	// Status indica el estado actual de la tienda (activa, inactiva, etc.)
	Status string `json:"status" firestore:"status"`
	// CreatedAt es la fecha de creación de la tienda
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
	// UpdatedAt es la fecha de la última actualización de la tienda
	UpdatedAt time.Time `json:"updatedAt" firestore:"updatedAt"`
	// Products es una lista de IDs de productos asociados a la tienda
	Products []string `json:"products" firestore:"products"`
}

// Validate verifica que los campos obligatorios de la tienda estén presentes y sean válidos.
func (s *Store) Validate() error {
	if s.UserID == "" {
		return fmt.Errorf("el ID del usuario es obligatorio")
	}
	if s.Name == "" {
		return fmt.Errorf("el nombre de la tienda es obligatorio")
	}
	if len(s.Description) < 10 {
		return fmt.Errorf("la descripción debe tener al menos 10 caracteres")
	}
	if err := s.Contact.Validate(); err != nil {
		return fmt.Errorf("información de contacto inválida: %v", err)
	}
	return nil
}

// BeforeCreate prepara la tienda para ser creada inicializando campos requeridos.
func (s *Store) BeforeCreate() {
	now := time.Now()
	s.CreatedAt = now
	s.UpdatedAt = now
	if s.Status == "" {
		s.Status = "active"
	}
	if s.Products == nil {
		s.Products = make([]string, 0)
	}
}

// BeforeUpdate actualiza la fecha de modificación de la tienda.
func (s *Store) BeforeUpdate() {
	s.UpdatedAt = time.Now()
}

