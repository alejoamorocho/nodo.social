package services

import (
	"context"
	"fmt"
	"time"

	"github.com/kha0sys/nodo.social/functions/domain/dto"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
)

// NodeService maneja la lógica de negocio relacionada con nodos
type NodeService struct {
	nodeRepo repositories.NodeRepository
	userRepo repositories.UserRepository
	feedRepo repositories.FeedRepository
}

// NewNodeService crea una nueva instancia de NodeService
func NewNodeService(
	nodeRepo repositories.NodeRepository,
	userRepo repositories.UserRepository,
	feedRepo repositories.FeedRepository,
) *NodeService {
	return &NodeService{
		nodeRepo: nodeRepo,
		userRepo: userRepo,
		feedRepo: feedRepo,
	}
}

// CreateNode crea un nuevo nodo
func (s *NodeService) CreateNode(ctx context.Context, nodeDTO dto.NodeDTO) (*models.Node, error) {
	node := &models.Node{
		Title:       nodeDTO.Title,
		Description: nodeDTO.Description,
		Type:        models.NodeType(nodeDTO.Type),
		UserID:      nodeDTO.UserID,
		Images:      nodeDTO.Images,
		CreatedAt:   nodeDTO.CreatedAt,
		UpdatedAt:   nodeDTO.UpdatedAt,
	}

	if err := s.nodeRepo.Create(ctx, node); err != nil {
		return nil, fmt.Errorf("error creating node: %v", err)
	}

	// Crear entrada en el feed
	feedItem := &models.FeedItem{
		UserID:    node.UserID,
		NodeID:    node.ID,
		Type:      "node_created",
		Content:   node,
		CreatedAt: time.Now().Unix(),
	}
	if err := s.feedRepo.Create(ctx, feedItem); err != nil {
		return nil, fmt.Errorf("error creating feed item: %v", err)
	}

	// Actualizar lista de nodos del usuario
	user, err := s.userRepo.Get(ctx, node.UserID)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	user.Nodes = append(user.Nodes, node.ID)
	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return node, nil
}

// GetNode obtiene un nodo por su ID
func (s *NodeService) GetNode(ctx context.Context, nodeID string) (*models.Node, error) {
	node, err := s.nodeRepo.Get(ctx, nodeID)
	if err != nil {
		return nil, fmt.Errorf("error getting node: %v", err)
	}
	return node, nil
}

// GetNodeFeed obtiene el feed de nodos
func (s *NodeService) GetNodeFeed(ctx context.Context) ([]*models.Node, error) {
	nodes, err := s.nodeRepo.GetPopularNodes(ctx, 20)
	if err != nil {
		return nil, fmt.Errorf("error getting recent nodes: %v", err)
	}
	return nodes, nil
}

// FollowNode permite a un usuario seguir un nodo
func (s *NodeService) FollowNode(ctx context.Context, nodeID string, userID string) error {
	// Obtener el nodo
	node, err := s.nodeRepo.Get(ctx, nodeID)
	if err != nil {
		return fmt.Errorf("error getting node: %v", err)
	}

	// Verificar si el usuario ya sigue el nodo
	for _, followerID := range node.Followers {
		if followerID == userID {
			return nil // Usuario ya sigue el nodo
		}
	}

	// Añadir el usuario a los seguidores del nodo
	node.Followers = append(node.Followers, userID)
	node.FollowersCount++

	if err := s.nodeRepo.Update(ctx, node); err != nil {
		return fmt.Errorf("error updating node: %v", err)
	}

	// Añadir el nodo a la lista de nodos seguidos del usuario
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	user.FollowedNodes = append(user.FollowedNodes, nodeID)
	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

// AddImage añade una imagen a un nodo
func (s *NodeService) AddImage(ctx context.Context, nodeID string, imageURL string) error {
	node, err := s.nodeRepo.Get(ctx, nodeID)
	if err != nil {
		return fmt.Errorf("error getting node: %v", err)
	}

	// Añadir la imagen al nodo
	node.Images = append(node.Images, imageURL)
	
	// Actualizar el nodo
	if err := s.nodeRepo.Update(ctx, node); err != nil {
		return fmt.Errorf("error updating node: %v", err)
	}

	return nil
}

// RemoveImage elimina una imagen de un nodo
func (s *NodeService) RemoveImage(ctx context.Context, nodeID string, imageURL string) error {
	node, err := s.nodeRepo.Get(ctx, nodeID)
	if err != nil {
		return fmt.Errorf("error getting node: %v", err)
	}

	// Eliminar la imagen del nodo
	for i, img := range node.Images {
		if img == imageURL {
			node.Images = append(node.Images[:i], node.Images[i+1:]...)
			break
		}
	}

	// Actualizar el nodo
	if err := s.nodeRepo.Update(ctx, node); err != nil {
		return fmt.Errorf("error updating node: %v", err)
	}

	return nil
}

// AddMedia añade un recurso multimedia a un nodo
func (s *NodeService) AddMedia(ctx context.Context, nodeID string, mediaURL string) error {
	node, err := s.nodeRepo.Get(ctx, nodeID)
	if err != nil {
		return fmt.Errorf("error getting node: %v", err)
	}

	// Añadir el recurso multimedia al nodo
	node.Media = append(node.Media, mediaURL)
	
	// Actualizar el nodo
	if err := s.nodeRepo.Update(ctx, node); err != nil {
		return fmt.Errorf("error updating node: %v", err)
	}

	return nil
}

// RemoveMedia elimina un recurso multimedia de un nodo
func (s *NodeService) RemoveMedia(ctx context.Context, nodeID string, mediaURL string) error {
	node, err := s.nodeRepo.Get(ctx, nodeID)
	if err != nil {
		return fmt.Errorf("error getting node: %v", err)
	}

	// Eliminar el recurso multimedia del nodo
	for i, media := range node.Media {
		if media == mediaURL {
			node.Media = append(node.Media[:i], node.Media[i+1:]...)
			break
		}
	}

	// Actualizar el nodo
	if err := s.nodeRepo.Update(ctx, node); err != nil {
		return fmt.Errorf("error updating node: %v", err)
	}

	return nil
}
