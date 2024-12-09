package triggers

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
	"github.com/kha0sys/nodo.social/domain/models"
	"github.com/kha0sys/nodo.social/domain/repositories"
	"github.com/kha0sys/nodo.social/internal/firebase"
)

type NodeTriggers struct {
	nodeRepo repositories.NodeRepository
}

func NewNodeTriggers(nodeRepo repositories.NodeRepository) *NodeTriggers {
	return &NodeTriggers{
		nodeRepo: nodeRepo,
	}
}

// OnNodeCreate es el trigger que se ejecuta cuando se crea un nuevo nodo
// Firestore trigger: projects/{project}/databases/{database}/documents/nodes/{nodeId}
func (t *NodeTriggers) OnNodeCreate(ctx context.Context, e firebase.FirestoreEvent) error {
	if _, err := metadata.FromContext(ctx); err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	var node models.Node
	if err := e.DataTo(&node); err != nil {
		return fmt.Errorf("e.DataTo: %v", err)
	}

	// Inicializar métricas
	node.Metrics = models.InteractionMetrics{
		Views:    0,
		Likes:    0,
		Shares:   0,
		Comments: 0,
	}

	// Actualizar el nodo con las métricas inicializadas
	if err := t.nodeRepo.Update(ctx, &node); err != nil {
		return fmt.Errorf("failed to initialize metrics: %v", err)
	}

	log.Printf("Initialized metrics for node: %s", node.ID)
	return nil
}

// OnProductLink es el trigger que se ejecuta cuando se vincula un producto a un nodo
// Firestore trigger: projects/{project}/databases/{database}/documents/nodes/{nodeId}/products/{productId}
func (t *NodeTriggers) OnProductLink(ctx context.Context, e firebase.FirestoreEvent) error {
	if _, err := metadata.FromContext(ctx); err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	var product models.Product
	if err := e.DataTo(&product); err != nil {
		return fmt.Errorf("e.DataTo: %v", err)
	}

	log.Printf("Product %s linked to node", product.ID)
	return nil
}

// OnFollowNode es el trigger que se ejecuta cuando un usuario sigue a un nodo
// Firestore trigger: projects/{project}/databases/{database}/documents/nodes/{nodeId}/followers/{userId}
func (t *NodeTriggers) OnFollowNode(ctx context.Context, e firebase.FirestoreEvent) error {
	if _, err := metadata.FromContext(ctx); err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	// Extraer nodeId y userId del path del documento
	docPath := e.Value.Name
	// TODO: Implementar extracción de IDs del path

	// Actualizar métricas del nodo
	metrics := models.InteractionMetrics{
		Views:    0,
		Likes:    0,
		Shares:   0,
		Comments: 0,
	}

	if err := t.nodeRepo.UpdateMetrics(ctx, docPath, metrics); err != nil {
		return fmt.Errorf("failed to update metrics: %v", err)
	}

	log.Printf("User followed node: %s", docPath)
	return nil
}
