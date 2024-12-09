package repositories

import (
    "context"
    "github.com/nodo-social/functions/domain/models"
)

type NodeRepository interface {
    Create(ctx context.Context, node *models.Node) error
    Update(ctx context.Context, node *models.Node) error
    Get(ctx context.Context, id string) (*models.Node, error)
    Delete(ctx context.Context, id string) error
    AddFollower(ctx context.Context, nodeID, userID string) error
    RemoveFollower(ctx context.Context, nodeID, userID string) error
    GetFollowers(ctx context.Context, nodeID string) ([]string, error)
    GetFeed(ctx context.Context, nodeID string) ([]models.Update, error)
    GetLinkedProducts(ctx context.Context, nodeID string) ([]models.Product, error)
    UpdateMetrics(ctx context.Context, nodeID string, metrics models.InteractionMetrics) error
}
