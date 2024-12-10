package triggers

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
	"github.com/kha0sys/nodo.social/functions/services"
	"github.com/kha0sys/nodo.social/functions/internal/firebase"
)

// NodeTriggers maneja los triggers relacionados con nodos
type NodeTriggers struct {
	client          *firestore.Client
	nodeRepo        repositories.NodeRepository
	userRepo        repositories.UserRepository
	feedRepo        repositories.FeedRepository
	notificationSvc *services.NotificationService
	achievementSvc  *services.AchievementService
}

// NewNodeTriggers crea una nueva instancia de NodeTriggers
func NewNodeTriggers(
	client *firestore.Client,
	nodeRepo repositories.NodeRepository,
	userRepo repositories.UserRepository,
	feedRepo repositories.FeedRepository,
	notificationSvc *services.NotificationService,
	achievementSvc *services.AchievementService,
) *NodeTriggers {
	return &NodeTriggers{
		client:          client,
		nodeRepo:        nodeRepo,
		userRepo:        userRepo,
		feedRepo:        feedRepo,
		notificationSvc: notificationSvc,
		achievementSvc:  achievementSvc,
	}
}

// OnCreate se ejecuta cuando se crea un nuevo nodo
func (t *NodeTriggers) OnCreate(ctx context.Context, e firebase.FirestoreEvent) error {
	var node models.Node
	if err := e.DataTo(&node); err != nil {
		return fmt.Errorf("error unmarshaling node: %v", err)
	}

	// Crear entrada en el feed
	feedItem := models.FeedItem{
		ID:        node.ID,
		Type:      "node_created",
		NodeID:    node.ID,
		UserID:    node.UserID,
		Content:   node,
		CreatedAt: time.Now().Unix(),
	}

	if err := t.feedRepo.Create(ctx, &feedItem); err != nil {
		log.Printf("error creating feed item: %v", err)
	}

	// Notificar a los seguidores
	followers, err := t.userRepo.GetFollowers(ctx, node.UserID)
	if err != nil {
		log.Printf("error getting followers: %v", err)
		return nil
	}

	notification := &models.Notification{
		Title:       fmt.Sprintf("Nuevo nodo: %s", node.Title),
		Description: fmt.Sprintf("Se ha creado un nuevo nodo: %s", node.Description),
		Type:        "node_created",
		UserID:      node.UserID,
		Data: map[string]interface{}{
			"nodeID":   node.ID,
			"nodeType": node.Type,
		},
	}

	for _, follower := range followers {
		notification.UserID = follower.ID
		if err := t.notificationSvc.CreateNotification(ctx, notification); err != nil {
			log.Printf("error sending notification to user %s: %v", follower.ID, err)
		}
	}

	return nil
}

// OnUpdate se ejecuta cuando se actualiza un nodo
func (t *NodeTriggers) OnUpdate(ctx context.Context, e firebase.FirestoreEvent) error {
	var oldNode models.Node
	if err := e.DataTo(&oldNode); err != nil {
		return fmt.Errorf("error unmarshaling old node: %v", err)
	}

	var newNode models.Node
	if err := e.DataTo(&newNode); err != nil {
		return fmt.Errorf("error unmarshaling new node: %v", err)
	}

	// Verificar cambios significativos
	if newNode.Title != oldNode.Title || newNode.Description != oldNode.Description {
		// Crear entrada en el feed
		feedItem := models.FeedItem{
			ID:        fmt.Sprintf("%s_update_%d", newNode.ID, time.Now().Unix()),
			Type:      "node_updated",
			NodeID:    newNode.ID,
			UserID:    newNode.UserID,
			Content:   newNode,
			CreatedAt: time.Now().Unix(),
		}

		if err := t.feedRepo.Create(ctx, &feedItem); err != nil {
			log.Printf("error creating feed item: %v", err)
		}

		// Notificar a los seguidores
		notification := &models.Notification{
			Title:       fmt.Sprintf("Actualización en nodo: %s", newNode.Title),
			Description: "Se han realizado cambios importantes en este nodo",
			Type:        "node_updated",
			UserID:      newNode.UserID,
			Data: map[string]interface{}{
				"nodeID":   newNode.ID,
				"nodeType": newNode.Type,
			},
		}

		for _, followerID := range newNode.Followers {
			notification.UserID = followerID
			if err := t.notificationSvc.CreateNotification(ctx, notification); err != nil {
				log.Printf("error sending notification to user %s: %v", followerID, err)
			}
		}
	}

	return nil
}

// OnDelete se ejecuta cuando se elimina un nodo
func (t *NodeTriggers) OnDelete(ctx context.Context, e firebase.FirestoreEvent) error {
	var node models.Node
	if err := e.DataTo(&node); err != nil {
		return fmt.Errorf("error unmarshaling node: %v", err)
	}

	// Eliminar las referencias en el feed
	if err := t.feedRepo.DeleteByNodeID(ctx, node.ID); err != nil {
		return fmt.Errorf("error eliminando referencias del feed: %v", err)
	}

	// Notificar a los seguidores
	notification := &models.Notification{
		Title:       fmt.Sprintf("Nodo eliminado: %s", node.Title),
		Description: "Un nodo que seguías ha sido eliminado",
		Type:        "node_deleted",
		UserID:      node.UserID,
		Data: map[string]interface{}{
			"nodeID":   node.ID,
			"nodeType": node.Type,
		},
	}

	for _, followerID := range node.Followers {
		notification.UserID = followerID
		if err := t.notificationSvc.CreateNotification(ctx, notification); err != nil {
			log.Printf("error sending notification to user %s: %v", followerID, err)
		}
	}

	return nil
}

// OnInteraction se ejecuta cuando hay una interacción con un nodo
func (t *NodeTriggers) OnInteraction(ctx context.Context, e firebase.FirestoreEvent) error {
	var node models.Node
	if err := e.DataTo(&node); err != nil {
		return fmt.Errorf("error unmarshaling node: %v", err)
	}

	// Actualizar métricas
	metrics := models.InteractionMetrics{
		Views:     node.Metrics.Views,
		Likes:     node.Metrics.Likes,
		Shares:    node.Metrics.Shares,
		Comments:  node.Metrics.Comments,
		Followers: node.FollowersCount,
	}

	// Actualizar métricas en el feed
	if err := t.feedRepo.UpdateMetrics(ctx, node.ID, metrics); err != nil {
		return fmt.Errorf("error actualizando métricas en el feed: %v", err)
	}

	// Notificar al creador sobre hitos importantes
	if metrics.Followers >= 100 {
		notification := &models.Notification{
			Title:       "¡Felicitaciones! Tu nodo ha alcanzado 100 seguidores",
			Description: fmt.Sprintf("Tu nodo '%s' está creciendo", node.Title),
			Type:        "achievement_followers",
			UserID:      node.UserID,
			Data: map[string]interface{}{
				"nodeID":     node.ID,
				"followers":  metrics.Followers,
				"milestone": 100,
			},
		}
		if err := t.notificationSvc.CreateNotification(ctx, notification); err != nil {
			log.Printf("error sending milestone notification: %v", err)
		}
	}

	return nil
}
