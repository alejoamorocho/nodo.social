package contact

import (
	"fmt"
	"net/url"
)

// ContactInfo representa la información de contacto de una entidad
type ContactInfo struct {
	Email     string `json:"email,omitempty" firestore:"email,omitempty"`
	Phone     string `json:"phone,omitempty" firestore:"phone,omitempty"`
	Website   string `json:"website,omitempty" firestore:"website,omitempty"`
	Instagram string `json:"instagram,omitempty" firestore:"instagram,omitempty"`
	Facebook  string `json:"facebook,omitempty" firestore:"facebook,omitempty"`
	Twitter   string `json:"twitter,omitempty" firestore:"twitter,omitempty"`
}

// Validate verifica que los campos de ContactInfo sean válidos
func (c *ContactInfo) Validate() error {
	if c.Website != "" {
		if _, err := url.Parse(c.Website); err != nil {
			return fmt.Errorf("website URL inválida: %v", err)
		}
	}

	// Aquí podrías agregar más validaciones para otros campos si es necesario
	return nil
}

