package repositories

import (
    "context"
    "fmt"
    "github.com/kha0sys/nodo.social/domain/models"
    "cloud.google.com/go/firestore"
)

// StoreRepository define las operaciones disponibles para la persistencia de tiendas
type StoreRepository interface {
    Create(ctx context.Context, store *models.Store) error
    Get(ctx context.Context, id string) (*models.Store, error)
    Update(ctx context.Context, store *models.Store) error
    Delete(ctx context.Context, id string) error
}

// FirestoreStoreRepository implementa StoreRepository usando Firestore
type FirestoreStoreRepository struct {
    client *firestore.Client
}

// NewFirestoreStoreRepository crea una nueva instancia de FirestoreStoreRepository
func NewFirestoreStoreRepository(client *firestore.Client) *FirestoreStoreRepository {
    return &FirestoreStoreRepository{
        client: client,
    }
}

func (r *FirestoreStoreRepository) Create(ctx context.Context, store *models.Store) error {
    if err := store.Validate(); err != nil {
        return fmt.Errorf("validación fallida: %w", err)
    }

    if store.ID == "" {
        store.ID = r.client.Collection("stores").NewDoc().ID
    }

    store.BeforeCreate()
    
    _, err := r.client.Collection("stores").Doc(store.ID).Set(ctx, store)
    return err
}

func (r *FirestoreStoreRepository) Get(ctx context.Context, id string) (*models.Store, error) {
    doc, err := r.client.Collection("stores").Doc(id).Get(ctx)
    if err != nil {
        return nil, err
    }

    var store models.Store
    if err := doc.DataTo(&store); err != nil {
        return nil, err
    }

    return &store, nil
}

func (r *FirestoreStoreRepository) Update(ctx context.Context, store *models.Store) error {
    if err := store.Validate(); err != nil {
        return fmt.Errorf("validación fallida: %w", err)
    }

    store.BeforeUpdate()
    
    _, err := r.client.Collection("stores").Doc(store.ID).Set(ctx, store)
    return err
}

func (r *FirestoreStoreRepository) Delete(ctx context.Context, id string) error {
    _, err := r.client.Collection("stores").Doc(id).Delete(ctx)
    return err
}
