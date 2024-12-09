package services

import (
	"context"

	"github.com/kha0sys/nodo.social/domain/dto"
	"github.com/kha0sys/nodo.social/domain/models"
	"github.com/kha0sys/nodo.social/domain/repositories"
)

// StoreService encapsula la lógica de negocio relacionada con tiendas
type StoreService struct {
	storeRepo repositories.StoreRepository
}

// NewStoreService crea una nueva instancia de StoreService
func NewStoreService(storeRepo repositories.StoreRepository) *StoreService {
	return &StoreService{
		storeRepo: storeRepo,
	}
}

// CreateStore crea una nueva tienda
func (s *StoreService) CreateStore(ctx context.Context, storeDTO dto.StoreDTO) (*models.Store, error) {
	store := storeDTO.ToModel()

	if err := store.Validate(); err != nil {
		return nil, err
	}

	store.BeforeCreate()

	if err := s.storeRepo.Create(ctx, store); err != nil {
		return nil, err
	}

	return store, nil
}

// GetStore obtiene una tienda por su ID
func (s *StoreService) GetStore(ctx context.Context, id string) (*models.Store, error) {
	return s.storeRepo.Get(ctx, id)
}

// UpdateStore actualiza una tienda existente
func (s *StoreService) UpdateStore(ctx context.Context, store *models.Store) error {
	if err := store.Validate(); err != nil {
		return err
	}

	store.BeforeUpdate()
	return s.storeRepo.Update(ctx, store)
}

// DeleteStore elimina una tienda
func (s *StoreService) DeleteStore(ctx context.Context, id string) error {
	return s.storeRepo.Delete(ctx, id)
}
