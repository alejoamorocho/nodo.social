package services

import (
	"context"
	"fmt"

	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
)

// StoreService maneja la lógica de negocio relacionada con tiendas
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
func (s *StoreService) CreateStore(ctx context.Context, store *models.Store) error {
	if err := store.Validate(); err != nil {
		return fmt.Errorf("invalid store data: %v", err)
	}

	if err := s.storeRepo.Create(ctx, store); err != nil {
		return fmt.Errorf("error creating store: %v", err)
	}
	return nil
}

// GetStore obtiene una tienda por su ID
func (s *StoreService) GetStore(ctx context.Context, storeID string) (*models.Store, error) {
	store, err := s.storeRepo.Get(ctx, storeID)
	if err != nil {
		return nil, fmt.Errorf("error getting store: %v", err)
	}
	return store, nil
}

// UpdateStore actualiza una tienda existente
func (s *StoreService) UpdateStore(ctx context.Context, store *models.Store) error {
	if err := store.Validate(); err != nil {
		return fmt.Errorf("invalid store data: %v", err)
	}

	if err := s.storeRepo.Update(ctx, store); err != nil {
		return fmt.Errorf("error updating store: %v", err)
	}
	return nil
}

// DeleteStore elimina una tienda
func (s *StoreService) DeleteStore(ctx context.Context, storeID string) error {
	if err := s.storeRepo.Delete(ctx, storeID); err != nil {
		return fmt.Errorf("error deleting store: %v", err)
	}
	return nil
}

// GetStoresByNode obtiene las tiendas asociadas a un nodo
func (s *StoreService) GetStoresByNode(ctx context.Context, nodeID string) ([]*models.Store, error) {
	stores, err := s.storeRepo.GetByNode(ctx, nodeID)
	if err != nil {
		return nil, fmt.Errorf("error getting stores by node: %v", err)
	}
	return stores, nil
}

// GetStoresByUser obtiene las tiendas de un usuario
func (s *StoreService) GetStoresByUser(ctx context.Context, userID string) ([]*models.Store, error) {
	stores, err := s.storeRepo.GetByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting stores by user: %v", err)
	}
	return stores, nil
}
