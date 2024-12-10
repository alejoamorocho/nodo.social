package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"google.golang.org/api/iterator"
)

// FeedRepository define la interfaz para operaciones con el feed
type FeedRepository interface {
	Create(ctx context.Context, item *models.FeedItem) error
	Delete(ctx context.Context, itemID string) error
	GetUserFeed(ctx context.Context, userID string, limit int, lastTimestamp int64) ([]*models.FeedItem, error)
	DeleteByNodeID(ctx context.Context, nodeID string) error
	UpdateMetrics(ctx context.Context, nodeID string, metrics models.InteractionMetrics) error
}

// FirestoreFeedRepository implementa FeedRepository usando Firestore
type FirestoreFeedRepository struct {
	client     *firestore.Client
	collection string
}

// NewFirestoreFeedRepository crea una nueva instancia de FirestoreFeedRepository
func NewFirestoreFeedRepository(client *firestore.Client) *FirestoreFeedRepository {
	return &FirestoreFeedRepository{
		client:     client,
		collection: "feed",
	}
}

// Create crea un nuevo item en el feed
func (r *FirestoreFeedRepository) Create(ctx context.Context, item *models.FeedItem) error {
	_, err := r.client.Collection(r.collection).Doc(item.ID).Set(ctx, item)
	return err
}

// Delete elimina un item del feed
func (r *FirestoreFeedRepository) Delete(ctx context.Context, itemID string) error {
	_, err := r.client.Collection(r.collection).Doc(itemID).Delete(ctx)
	return err
}

// GetUserFeed obtiene el feed de un usuario
func (r *FirestoreFeedRepository) GetUserFeed(ctx context.Context, userID string, limit int, lastTimestamp int64) ([]*models.FeedItem, error) {
	query := r.client.Collection(r.collection).
		Where("userID", "==", userID).
		OrderBy("createdAt", firestore.Desc).
		Limit(limit)

	if lastTimestamp > 0 {
		query = query.Where("createdAt", "<", lastTimestamp)
	}

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var items []*models.FeedItem
	for _, doc := range docs {
		var item models.FeedItem
		if err := doc.DataTo(&item); err != nil {
			return nil, err
		}
		item.ID = doc.Ref.ID
		items = append(items, &item)
	}

	return items, nil
}

// DeleteByNodeID elimina todos los items del feed relacionados con un nodo
func (r *FirestoreFeedRepository) DeleteByNodeID(ctx context.Context, nodeID string) error {
	// Obtener todos los items relacionados con el nodo
	iter := r.client.Collection(r.collection).Where("nodeID", "==", nodeID).Documents(ctx)
	batch := r.client.Batch()
	
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		batch.Delete(doc.Ref)
	}

	_, err := batch.Commit(ctx)
	return err
}

// UpdateMetrics actualiza las métricas de todos los items del feed relacionados con un nodo
func (r *FirestoreFeedRepository) UpdateMetrics(ctx context.Context, nodeID string, metrics models.InteractionMetrics) error {
	// Obtener todos los items relacionados con el nodo
	iter := r.client.Collection(r.collection).Where("nodeID", "==", nodeID).Documents(ctx)
	batch := r.client.Batch()
	
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		// Actualizar las métricas en el contenido del item
		var item models.FeedItem
		if err := doc.DataTo(&item); err != nil {
			return err
		}

		if node, ok := item.Content.(models.Node); ok {
			node.Metrics = metrics
			item.Content = node
			batch.Set(doc.Ref, item)
		}
	}

	_, err := batch.Commit(ctx)
	return err
}
