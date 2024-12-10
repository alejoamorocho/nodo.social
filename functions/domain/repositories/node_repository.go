package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
)

// NodeRepository define la interfaz para operaciones con nodos
type NodeRepository interface {
	Create(ctx context.Context, node *models.Node) error
	Get(ctx context.Context, nodeID string) (*models.Node, error)
	Update(ctx context.Context, node *models.Node) error
	Delete(ctx context.Context, nodeID string) error
	GetTotalNodes(ctx context.Context) (int, error)
	GetPopularNodes(ctx context.Context, limit int) ([]*models.Node, error)
}

// FirestoreNodeRepository implementa NodeRepository usando Firestore
type FirestoreNodeRepository struct {
	client     *firestore.Client
	collection string
}

// NewFirestoreNodeRepository crea una nueva instancia de FirestoreNodeRepository
func NewFirestoreNodeRepository(client *firestore.Client) *FirestoreNodeRepository {
	return &FirestoreNodeRepository{
		client:     client,
		collection: "nodes",
	}
}

// Create crea un nuevo nodo en Firestore
func (r *FirestoreNodeRepository) Create(ctx context.Context, node *models.Node) error {
	// Preparar el nodo para su creación
	node.BeforeCreate()

	// Validar el nodo
	if err := models.ValidateNode(node); err != nil {
		return err
	}

	// Si no hay ID, crear uno nuevo
	if node.ID == "" {
		doc := r.client.Collection(r.collection).NewDoc()
		node.ID = doc.ID
	}

	// Crear el documento
	_, err := r.client.Collection(r.collection).Doc(node.ID).Set(ctx, node)
	return err
}

// Get obtiene un nodo por su ID
func (r *FirestoreNodeRepository) Get(ctx context.Context, nodeID string) (*models.Node, error) {
	doc, err := r.client.Collection(r.collection).Doc(nodeID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var node models.Node
	if err := doc.DataTo(&node); err != nil {
		return nil, err
	}

	node.ID = doc.Ref.ID
	return &node, nil
}

// Update actualiza un nodo existente
func (r *FirestoreNodeRepository) Update(ctx context.Context, node *models.Node) error {
	// Preparar el nodo para su actualización
	node.BeforeUpdate()

	// Validar el nodo
	if err := models.ValidateNode(node); err != nil {
		return err
	}

	_, err := r.client.Collection(r.collection).Doc(node.ID).Set(ctx, node)
	return err
}

// Delete elimina un nodo
func (r *FirestoreNodeRepository) Delete(ctx context.Context, nodeID string) error {
	_, err := r.client.Collection(r.collection).Doc(nodeID).Delete(ctx)
	return err
}

// GetTotalNodes obtiene el número total de nodos
func (r *FirestoreNodeRepository) GetTotalNodes(ctx context.Context) (int, error) {
	docs, err := r.client.Collection(r.collection).Documents(ctx).GetAll()
	if err != nil {
		return 0, err
	}
	return len(docs), nil
}

// GetPopularNodes obtiene los nodos más populares basados en el número de seguidores
func (r *FirestoreNodeRepository) GetPopularNodes(ctx context.Context, limit int) ([]*models.Node, error) {
	docs, err := r.client.Collection(r.collection).
		OrderBy("followersCount", firestore.Desc).
		Limit(limit).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}

	var nodes []*models.Node
	for _, doc := range docs {
		var node models.Node
		if err := doc.DataTo(&node); err != nil {
			continue
		}
		node.ID = doc.Ref.ID
		nodes = append(nodes, &node)
	}

	return nodes, nil
}
