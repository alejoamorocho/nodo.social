package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"google.golang.org/api/iterator"
)

// StoreRepository define la interfaz para operaciones con tiendas
type StoreRepository interface {
	Create(ctx context.Context, store *models.Store) error
	Get(ctx context.Context, storeID string) (*models.Store, error)
	Update(ctx context.Context, store *models.Store) error
	Delete(ctx context.Context, storeID string) error
	GetByNode(ctx context.Context, nodeID string) ([]*models.Store, error)
	GetByUser(ctx context.Context, userID string) ([]*models.Store, error)
}

// FirestoreStoreRepository implementa StoreRepository usando Firestore
type FirestoreStoreRepository struct {
	client     *firestore.Client
	collection string
}

// NewFirestoreStoreRepository crea una nueva instancia de FirestoreStoreRepository
func NewFirestoreStoreRepository(client *firestore.Client) *FirestoreStoreRepository {
	return &FirestoreStoreRepository{
		client:     client,
		collection: "stores",
	}
}

// Create crea una nueva tienda en Firestore
func (r *FirestoreStoreRepository) Create(ctx context.Context, store *models.Store) error {
	// Si no hay ID, crear uno nuevo
	if store.ID == "" {
		doc := r.client.Collection(r.collection).NewDoc()
		store.ID = doc.ID
	}

	_, err := r.client.Collection(r.collection).Doc(store.ID).Set(ctx, store)
	return err
}

// Get obtiene una tienda por su ID
func (r *FirestoreStoreRepository) Get(ctx context.Context, storeID string) (*models.Store, error) {
	doc, err := r.client.Collection(r.collection).Doc(storeID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var store models.Store
	if err := doc.DataTo(&store); err != nil {
		return nil, err
	}

	return &store, nil
}

// Update actualiza una tienda existente
func (r *FirestoreStoreRepository) Update(ctx context.Context, store *models.Store) error {
	_, err := r.client.Collection(r.collection).Doc(store.ID).Set(ctx, store)
	return err
}

// Delete elimina una tienda
func (r *FirestoreStoreRepository) Delete(ctx context.Context, storeID string) error {
	_, err := r.client.Collection(r.collection).Doc(storeID).Delete(ctx)
	return err
}

// GetByNode obtiene las tiendas asociadas a un nodo
func (r *FirestoreStoreRepository) GetByNode(ctx context.Context, nodeID string) ([]*models.Store, error) {
	iter := r.client.Collection(r.collection).Where("nodeID", "==", nodeID).Documents(ctx)
	defer iter.Stop()

	var stores []*models.Store
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var store models.Store
		if err := doc.DataTo(&store); err != nil {
			return nil, err
		}
		stores = append(stores, &store)
	}

	return stores, nil
}

// GetByUser obtiene las tiendas de un usuario
func (r *FirestoreStoreRepository) GetByUser(ctx context.Context, userID string) ([]*models.Store, error) {
	iter := r.client.Collection(r.collection).Where("userID", "==", userID).Documents(ctx)
	defer iter.Stop()

	var stores []*models.Store
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var store models.Store
		if err := doc.DataTo(&store); err != nil {
			return nil, err
		}
		stores = append(stores, &store)
	}

	return stores, nil
}
