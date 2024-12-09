package services

import (
    "context"
    "github.com/kha0sys/nodo.social/domain/models"
    "github.com/kha0sys/nodo.social/domain/repositories"
    "github.com/kha0sys/nodo.social/domain/dto"
)

// NodeService encapsula la lógica de negocio relacionada con nodos
type NodeService struct {
    nodeRepo repositories.NodeRepository
}

// NewNodeService crea una nueva instancia de NodeService
func NewNodeService(nodeRepo repositories.NodeRepository) *NodeService {
    return &NodeService{
        nodeRepo: nodeRepo,
    }
}

// CreateNode crea un nuevo nodo
func (s *NodeService) CreateNode(ctx context.Context, nodeDTO dto.NodeDTO) (*models.Node, error) {
    node := nodeDTO.ToModel()
    if err := s.nodeRepo.Create(ctx, node); err != nil {
        return nil, err
    }
    return node, nil
}

// GetNode obtiene un nodo por su ID
func (s *NodeService) GetNode(ctx context.Context, id string) (*models.Node, error) {
    node, err := s.nodeRepo.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return node, nil
}

// GetNodeFeed obtiene el feed de nodos
func (s *NodeService) GetNodeFeed(ctx context.Context) ([]*models.Node, error) {
    // Aquí podrías implementar lógica adicional como paginación, filtros, etc.
    // Por ahora es un placeholder
    return []*models.Node{}, nil
}

// FollowNode permite a un usuario seguir a un nodo
func (s *NodeService) FollowNode(ctx context.Context, nodeID, userID string) error {
    if err := s.nodeRepo.AddFollower(ctx, nodeID, userID); err != nil {
        return err
    }
    return nil
}

// UpdateNodeMetrics actualiza las métricas de interacción de un nodo
func (s *NodeService) UpdateNodeMetrics(ctx context.Context, nodeID string, metrics models.InteractionMetrics) error {
    return s.nodeRepo.UpdateMetrics(ctx, nodeID, metrics)
}

// UpdateNode actualiza un nodo existente
func (s *NodeService) UpdateNode(ctx context.Context, node *models.Node) error {
    return s.nodeRepo.Update(ctx, node)
}

// DeleteNode elimina un nodo
func (s *NodeService) DeleteNode(ctx context.Context, id string) error {
    return s.nodeRepo.Delete(ctx, id)
}

// AddFollower añade un seguidor a un nodo
func (s *NodeService) AddFollower(ctx context.Context, nodeID, userID string) error {
    return s.nodeRepo.AddFollower(ctx, nodeID, userID)
}

// RemoveFollower elimina un seguidor de un nodo
func (s *NodeService) RemoveFollower(ctx context.Context, nodeID, userID string) error {
    return s.nodeRepo.RemoveFollower(ctx, nodeID, userID)
}

// GetFollowers obtiene los seguidores de un nodo
func (s *NodeService) GetFollowers(ctx context.Context, nodeID string) ([]string, error) {
    return s.nodeRepo.GetFollowers(ctx, nodeID)
}

// GetLinkedProducts obtiene los productos vinculados a un nodo
func (s *NodeService) GetLinkedProducts(ctx context.Context, nodeID string) ([]models.Product, error) {
    return s.nodeRepo.GetLinkedProducts(ctx, nodeID)
}

// GetFeed obtiene el feed de nodos según los filtros especificados
func (s *NodeService) GetFeed(ctx context.Context, filters models.FeedFilters) ([]models.Node, error) {
    return s.nodeRepo.GetFeed(ctx, filters)
}
