package services

import (
	"context"

	"github.com/kha0sys/nodo.social/domain/models"
	"github.com/kha0sys/nodo.social/domain/repositories"
	"github.com/kha0sys/nodo.social/domain/dto"
)

// ProductService encapsula la lógica de negocio relacionada con productos
type ProductService struct {
	productRepo repositories.ProductRepository
}

// NewProductService crea una nueva instancia de ProductService
func NewProductService(productRepo repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// CreateProduct crea un nuevo producto
func (s *ProductService) CreateProduct(ctx context.Context, productDTO *dto.ProductDTO) (*models.Product, error) {
	product := productDTO.ToModel()
	if err := product.Validate(); err != nil {
		return nil, err
	}

	if err := s.productRepo.Create(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

// GetProduct obtiene un producto por su ID
func (s *ProductService) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	return s.productRepo.Get(ctx, id)
}

// UpdateProduct actualiza un producto existente
func (s *ProductService) UpdateProduct(ctx context.Context, productDTO *dto.ProductDTO) (*models.Product, error) {
	product := productDTO.ToModel()
	if err := product.Validate(); err != nil {
		return nil, err
	}

	if err := s.productRepo.Update(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

// DeleteProduct elimina un producto
func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	return s.productRepo.Delete(ctx, id)
}

// ListProducts obtiene una lista de productos que coinciden con los criterios de búsqueda
func (s *ProductService) ListProducts(ctx context.Context, filters map[string]interface{}) ([]*models.Product, error) {
	return s.productRepo.FindByFilters(ctx, filters)
}

// ApproveProduct aprueba un producto para su publicación
func (s *ProductService) ApproveProduct(ctx context.Context, id string) error {
	product, err := s.productRepo.Get(ctx, id)
	if err != nil {
		return err
	}
	
	product.ApprovalStatus = "approved"
	return s.productRepo.Update(ctx, product)
}

// GetProductsByNode obtiene los productos asociados a un nodo
func (s *ProductService) GetProductsByNode(ctx context.Context, nodeID string) ([]*models.Product, error) {
	filters := map[string]interface{}{
		"nodeId": nodeID,
	}
	return s.productRepo.FindByFilters(ctx, filters)
}
