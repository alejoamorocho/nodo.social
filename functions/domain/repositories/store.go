package repositories

import (
    "context"
    "github.com/nodo-social/functions/domain/models"
)

type StoreRepository interface {
    Create(ctx context.Context, store *models.Store) error
    Update(ctx context.Context, store *models.Store) error
    Get(ctx context.Context, id string) (*models.Store, error)
    Delete(ctx context.Context, id string) error
    AddProduct(ctx context.Context, storeID string, product *models.Product) error
    UpdateProduct(ctx context.Context, storeID string, product *models.Product) error
    DeleteProduct(ctx context.Context, storeID, productID string) error
    GetProducts(ctx context.Context, storeID string) ([]models.Product, error)
    UpdateProductStatus(ctx context.Context, productID, status string) error
}
