package models

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kha0sys/nodo.social/functions/domain/models/contact"
)

// ValidationError representa un error de validación
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateNode valida un nodo antes de su creación o actualización
func ValidateNode(node *Node) error {
	// Validar título (3-100 chars)
	if len(strings.TrimSpace(node.Title)) < 3 {
		return &ValidationError{
			Field:   "Title",
			Message: "el título debe tener al menos 3 caracteres",
		}
	}
	if len(node.Title) > 100 {
		return &ValidationError{
			Field:   "Title",
			Message: "el título no puede tener más de 100 caracteres",
		}
	}

	// Validar descripción (10-1000 chars)
	if len(strings.TrimSpace(node.Description)) < 10 {
		return &ValidationError{
			Field:   "Description",
			Message: "la descripción debe tener al menos 10 caracteres",
		}
	}
	if len(node.Description) > 1000 {
		return &ValidationError{
			Field:   "Description",
			Message: "la descripción no puede tener más de 1000 caracteres",
		}
	}

	// Validar tipo
	validTypes := map[NodeType]bool{
		Social:        true,
		Environmental: true,
		Animal:        true,
	}
	if !validTypes[node.Type] {
		return &ValidationError{
			Field:   "Type",
			Message: "tipo de nodo inválido",
		}
	}

	// Validar imágenes
	if len(node.Media) > 10 {
		return &ValidationError{
			Field:   "Media",
			Message: "no se pueden agregar más de 10 imágenes",
		}
	}

	// Validar URLs de imágenes
	for i, mediaURL := range node.Media {
		if _, err := url.ParseRequestURI(mediaURL); err != nil {
			return &ValidationError{
				Field:   fmt.Sprintf("Media[%d]", i),
				Message: "URL de imagen inválida",
			}
		}
	}

	return nil
}

// ValidateProduct valida un producto antes de su creación o actualización
func ValidateProduct(product *Product) error {
	// Validar nombre
	if len(strings.TrimSpace(product.Name)) == 0 {
		return &ValidationError{
			Field:   "Name",
			Message: "el nombre es requerido",
		}
	}

	// Validar descripción
	if len(strings.TrimSpace(product.Description)) < 10 {
		return &ValidationError{
			Field:   "Description",
			Message: "la descripción debe tener al menos 10 caracteres",
		}
	}

	// Validar precio
	if product.Price <= 0 {
		return &ValidationError{
			Field:   "Price",
			Message: "el precio debe ser mayor a 0",
		}
	}

	// Validar porcentaje de donación
	if product.DonationPercent < 1 || product.DonationPercent > 100 {
		return &ValidationError{
			Field:   "DonationPercent",
			Message: "el porcentaje de donación debe estar entre 1 y 100",
		}
	}

	// Validar imágenes
	if len(product.Images) == 0 {
		return &ValidationError{
			Field:   "Images",
			Message: "debe proporcionar al menos una imagen",
		}
	}
	if len(product.Images) > 5 {
		return &ValidationError{
			Field:   "Images",
			Message: "no se pueden agregar más de 5 imágenes",
		}
	}

	// Validar URLs de imágenes
	for i, imageURL := range product.Images {
		if _, err := url.ParseRequestURI(imageURL); err != nil {
			return &ValidationError{
				Field:   fmt.Sprintf("Images[%d]", i),
				Message: "URL de imagen inválida",
			}
		}
	}

	// Validar IDs requeridos
	if product.StoreID == "" {
		return &ValidationError{
			Field:   "StoreID",
			Message: "el ID de la tienda es requerido",
		}
	}
	if product.NodeID == "" {
		return &ValidationError{
			Field:   "NodeID",
			Message: "el ID del nodo es requerido",
		}
	}
	if product.UserID == "" {
		return &ValidationError{
			Field:   "UserID",
			Message: "el ID del usuario es requerido",
		}
	}

	// Validar contacto
	if err := ValidateContact(product.Contact); err != nil {
		return err
	}

	return nil
}

// ValidateContact valida la información de contacto
func ValidateContact(contactInfo contact.ContactInfo) error {
	// Al menos un método de contacto debe estar presente
	if contactInfo.Email == "" && contactInfo.Phone == "" && 
		contactInfo.Website == "" && contactInfo.Instagram == "" && 
		contactInfo.Facebook == "" && contactInfo.Twitter == "" {
		return &ValidationError{
			Field:   "Contact",
			Message: "debe proporcionar al menos un método de contacto",
		}
	}

	// Validar email si está presente
	if contactInfo.Email != "" {
		if !strings.Contains(contactInfo.Email, "@") || !strings.Contains(contactInfo.Email, ".") {
			return &ValidationError{
				Field:   "Contact.Email",
				Message: "email inválido",
			}
		}
	}

	// Validar teléfono si está presente
	if contactInfo.Phone != "" {
		// Eliminar caracteres no numéricos
		phone := strings.Map(func(r rune) rune {
			if r >= '0' && r <= '9' {
				return r
			}
			return -1
		}, contactInfo.Phone)

		if len(phone) < 10 {
			return &ValidationError{
				Field:   "Contact.Phone",
				Message: "número de teléfono debe tener al menos 10 dígitos",
			}
		}
	}

	// Validar website si está presente
	if contactInfo.Website != "" {
		if _, err := url.Parse(contactInfo.Website); err != nil {
			return &ValidationError{
				Field:   "Contact.Website",
				Message: "URL del sitio web inválida",
			}
		}
	}

	// Validar redes sociales
	if contactInfo.Instagram != "" {
		if !strings.HasPrefix(contactInfo.Instagram, "@") {
			return &ValidationError{
				Field:   "Contact.Instagram",
				Message: "el usuario de Instagram debe comenzar con @",
			}
		}
	}

	return nil
}
