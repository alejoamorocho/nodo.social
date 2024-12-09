package repositories

import (
    "context"
    "fmt"

    "cloud.google.com/go/firestore"
    "github.com/kha0sys/nodo.social/domain/models"
    "google.golang.org/api/iterator"
)

// ProductRepository define las operaciones disponibles para la persistencia de productos
type ProductRepository interface {
    Create(ctx context.Context, product *models.Product) error
    Get(ctx context.Context, id string) (*models.Product, error)
    Update(ctx context.Context, product *models.Product) error
    Delete(ctx context.Context, id string) error
    FindByFilters(ctx context.Context, filters map[string]interface{}) ([]*models.Product, error)
}

// FirestoreProductRepository implementa ProductRepository usando Firestore
type FirestoreProductRepository struct {
    client *firestore.Client
}

// NewFirestoreProductRepository crea una nueva instancia de FirestoreProductRepository
func NewFirestoreProductRepository(client *firestore.Client) *FirestoreProductRepository {
    return &FirestoreProductRepository{
        client: client,
    }
}

func (r *FirestoreProductRepository) Create(ctx context.Context, product *models.Product) error {
    if err := product.Validate(); err != nil {
        return fmt.Errorf("validación fallida: %w", err)
    }

    if product.ID == "" {
        product.ID = r.client.Collection("products").NewDoc().ID
    }

    product.BeforeCreate()
    
    _, err := r.client.Collection("products").Doc(product.ID).Set(ctx, product)
    return err
}

func (r *FirestoreProductRepository) Get(ctx context.Context, id string) (*models.Product, error) {
    doc, err := r.client.Collection("products").Doc(id).Get(ctx)
    if err != nil {
        return nil, fmt.Errorf("error al obtener el producto: %w", err)
    }

    var product models.Product
    if err := doc.DataTo(&product); err != nil {
        return nil, fmt.Errorf("error al deserializar el producto: %w", err)
    }
    return &product, nil
}

func (r *FirestoreProductRepository) Update(ctx context.Context, product *models.Product) error {
    if err := product.Validate(); err != nil {
        return fmt.Errorf("validación fallida: %w", err)
    }

    product.BeforeUpdate()
    
    _, err := r.client.Collection("products").Doc(product.ID).Set(ctx, product)
    if err != nil {
        return fmt.Errorf("error al actualizar el producto: %w", err)
    }
    return nil
}

func (r *FirestoreProductRepository) Delete(ctx context.Context, id string) error {
    _, err := r.client.Collection("products").Doc(id).Delete(ctx)
    if err != nil {
        return fmt.Errorf("error al eliminar el producto: %w", err)
    }
    return nil
}

func (r *FirestoreProductRepository) FindByFilters(ctx context.Context, filters map[string]interface{}) ([]*models.Product, error) {
    query := r.client.Collection("products").Query
    
    // Apply filters to the query
    for field, value := range filters {
        query = query.Where(field, "==", value)
    }
    
    iter := query.Documents(ctx)
    var products []*models.Product

    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("error al iterar sobre los productos: %w", err)
        }

        var product models.Product
        if err := doc.DataTo(&product); err != nil {
            return nil, fmt.Errorf("error al deserializar el producto: %w", err)
        }
        products = append(products, &product)
    }

    return products, nil
}
