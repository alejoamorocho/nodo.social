package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
)

// ProductRepository define la interfaz para operaciones con productos
type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	Get(ctx context.Context, productID string) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, productID string) error
	GetByNode(ctx context.Context, nodeID string) ([]*models.Product, error)
}

// FirestoreProductRepository implementa ProductRepository usando Firestore
type FirestoreProductRepository struct {
	client     *firestore.Client
	collection string
}

// NewFirestoreProductRepository crea una nueva instancia de FirestoreProductRepository
func NewFirestoreProductRepository(client *firestore.Client) *FirestoreProductRepository {
	return &FirestoreProductRepository{
		client:     client,
		collection: "products",
	}
}

// Create crea un nuevo producto en Firestore
func (r *FirestoreProductRepository) Create(ctx context.Context, product *models.Product) error {
	// Preparar el producto para su creación
	product.BeforeCreate()

	// Si no hay ID, crear uno nuevo
	if product.ID == "" {
		doc := r.client.Collection(r.collection).NewDoc()
		product.ID = doc.ID
	}

	// Crear el documento
	_, err := r.client.Collection(r.collection).Doc(product.ID).Set(ctx, product)
	return err
}

// Get obtiene un producto por su ID
func (r *FirestoreProductRepository) Get(ctx context.Context, productID string) (*models.Product, error) {
	doc, err := r.client.Collection(r.collection).Doc(productID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var product models.Product
	if err := doc.DataTo(&product); err != nil {
		return nil, err
	}

	product.ID = doc.Ref.ID
	return &product, nil
}

// Update actualiza un producto existente
func (r *FirestoreProductRepository) Update(ctx context.Context, product *models.Product) error {
	// Preparar el producto para su actualización
	product.BeforeUpdate()

	_, err := r.client.Collection(r.collection).Doc(product.ID).Set(ctx, product)
	return err
}

// Delete elimina un producto
func (r *FirestoreProductRepository) Delete(ctx context.Context, productID string) error {
	_, err := r.client.Collection(r.collection).Doc(productID).Delete(ctx)
	return err
}

// GetByNode obtiene todos los productos asociados a un nodo
func (r *FirestoreProductRepository) GetByNode(ctx context.Context, nodeID string) ([]*models.Product, error) {
	iter := r.client.Collection(r.collection).Where("nodeId", "==", nodeID).Documents(ctx)
	
	var products []*models.Product
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var product models.Product
		if err := doc.DataTo(&product); err != nil {
			continue
		}
		product.ID = doc.Ref.ID
		products = append(products, &product)
	}

	return products, nil
}
