package services

import (
	"context"
	"fmt"

	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
)

// ProductService maneja la lógica de negocio relacionada con productos
type ProductService struct {
	productRepo repositories.ProductRepository
	nodeRepo    repositories.NodeRepository
}

// NewProductService crea una nueva instancia de ProductService
func NewProductService(productRepo repositories.ProductRepository, nodeRepo repositories.NodeRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		nodeRepo:    nodeRepo,
	}
}

// CreateProduct crea un nuevo producto
func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	if err := s.productRepo.Create(ctx, product); err != nil {
		return fmt.Errorf("error creating product: %v", err)
	}

	// Actualizar la lista de productos del nodo
	node, err := s.nodeRepo.Get(ctx, product.NodeID)
	if err != nil {
		return fmt.Errorf("error getting node: %v", err)
	}

	node.Products = append(node.Products, product.ID)
	if err := s.nodeRepo.Update(ctx, node); err != nil {
		return fmt.Errorf("error updating node: %v", err)
	}

	return nil
}

// GetProduct obtiene un producto por su ID
func (s *ProductService) GetProduct(ctx context.Context, productID string) (*models.Product, error) {
	product, err := s.productRepo.Get(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("error getting product: %v", err)
	}
	return product, nil
}

// UpdateProduct actualiza un producto existente
func (s *ProductService) UpdateProduct(ctx context.Context, product *models.Product) error {
	if err := s.productRepo.Update(ctx, product); err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}
	return nil
}

// DeleteProduct elimina un producto
func (s *ProductService) DeleteProduct(ctx context.Context, productID string) error {
	// Obtener el producto para saber su nodo
	product, err := s.productRepo.Get(ctx, productID)
	if err != nil {
		return fmt.Errorf("error getting product: %v", err)
	}

	// Eliminar el producto de la lista de productos del nodo
	node, err := s.nodeRepo.Get(ctx, product.NodeID)
	if err != nil {
		return fmt.Errorf("error getting node: %v", err)
	}

	// Eliminar el ID del producto de la lista del nodo
	for i, id := range node.Products {
		if id == productID {
			node.Products = append(node.Products[:i], node.Products[i+1:]...)
			break
		}
	}

	if err := s.nodeRepo.Update(ctx, node); err != nil {
		return fmt.Errorf("error updating node: %v", err)
	}

	// Eliminar el producto
	if err := s.productRepo.Delete(ctx, productID); err != nil {
		return fmt.Errorf("error deleting product: %v", err)
	}

	return nil
}

// ApproveProduct aprueba un producto para su publicación
func (s *ProductService) ApproveProduct(ctx context.Context, productID string) error {
	product, err := s.productRepo.Get(ctx, productID)
	if err != nil {
		return fmt.Errorf("error getting product: %v", err)
	}

	product.Status = "approved"
	if err := s.productRepo.Update(ctx, product); err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}

	return nil
}

// GetProductsByNode obtiene todos los productos asociados a un nodo
func (s *ProductService) GetProductsByNode(ctx context.Context, nodeID string) ([]*models.Product, error) {
	products, err := s.productRepo.GetByNode(ctx, nodeID)
	if err != nil {
		return nil, fmt.Errorf("error getting products by node: %v", err)
	}
	return products, nil
}

// AddImage añade una imagen a un producto
func (s *ProductService) AddImage(ctx context.Context, productID string, imageURL string) error {
	product, err := s.productRepo.Get(ctx, productID)
	if err != nil {
		return fmt.Errorf("error getting product: %v", err)
	}

	product.Images = append(product.Images, imageURL)
	if err := s.productRepo.Update(ctx, product); err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}

	return nil
}

// RemoveImage elimina una imagen de un producto
func (s *ProductService) RemoveImage(ctx context.Context, productID string, imageURL string) error {
	product, err := s.productRepo.Get(ctx, productID)
	if err != nil {
		return fmt.Errorf("error getting product: %v", err)
	}

	for i, img := range product.Images {
		if img == imageURL {
			product.Images = append(product.Images[:i], product.Images[i+1:]...)
			break
		}
	}

	if err := s.productRepo.Update(ctx, product); err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}

	return nil
}

// UpdateImages actualiza todas las imágenes de un producto
func (s *ProductService) UpdateImages(ctx context.Context, productID string, images []string) error {
	product, err := s.productRepo.Get(ctx, productID)
	if err != nil {
		return fmt.Errorf("error getting product: %v", err)
	}

	product.Images = images
	
	if err := s.productRepo.Update(ctx, product); err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}

	return nil
}
