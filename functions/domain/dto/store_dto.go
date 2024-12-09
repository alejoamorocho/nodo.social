package dto

import (
	"time"

	"github.com/kha0sys/nodo.social/domain/models"
)

// StoreDTO representa los datos de una tienda para transferencia
// entre capas de la aplicación
type StoreDTO struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Contact     models.ContactInfo `json:"contact"`
	UserID      string            `json:"userId"`
	Logo        string            `json:"logo"`
	Products    []string          `json:"products"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

// ToModel convierte el DTO a un modelo Store
func (dto *StoreDTO) ToModel() *models.Store {
	return &models.Store{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Contact:     dto.Contact,
		UserID:      dto.UserID,
		Logo:        dto.Logo,
		Products:    dto.Products,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

// FromModel crea un DTO a partir de un modelo Store
func FromStoreModel(store *models.Store) *StoreDTO {
	return &StoreDTO{
		ID:          store.ID,
		Name:        store.Name,
		Description: store.Description,
		Contact:     store.Contact,
		UserID:      store.UserID,
		Logo:        store.Logo,
		Products:    store.Products,
		CreatedAt:   store.CreatedAt,
		UpdatedAt:   store.UpdatedAt,
	}
}
