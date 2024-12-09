package repositories

import (
    "context"
    "fmt"
    "time"

    "cloud.google.com/go/firestore"
    "github.com/kha0sys/nodo.social/domain/models"
    "google.golang.org/api/iterator"
)

type FirestoreNodeRepository struct {
    client *firestore.Client
}

func NewFirestoreNodeRepository(client *firestore.Client) *FirestoreNodeRepository {
    return &FirestoreNodeRepository{
        client: client,
    }
}

func (r *FirestoreNodeRepository) Create(ctx context.Context, node *models.Node) error {
    if err := ValidateNode(node); err != nil {
        return fmt.Errorf("validación fallida: %w", err)
    }

    if node.ID == "" {
        node.ID = r.client.Collection("nodes").NewDoc().ID
    }

    node.BeforeCreate()
    
    _, err := r.client.Collection("nodes").Doc(node.ID).Set(ctx, node)
    if err != nil {
        return fmt.Errorf("error al crear nodo: %v", err)
    }
    return nil
}

func (r *FirestoreNodeRepository) Update(ctx context.Context, node *models.Node) error {
    if err := ValidateNode(node); err != nil {
        return fmt.Errorf("validación fallida: %w", err)
    }

    node.BeforeUpdate()
    
    _, err := r.client.Collection("nodes").Doc(node.ID).Set(ctx, node)
    if err != nil {
        return fmt.Errorf("error al actualizar nodo: %v", err)
    }
    return nil
}

func (r *FirestoreNodeRepository) Get(ctx context.Context, id string) (*models.Node, error) {
    if id == "" {
        return nil, &NodeValidationError{Field: "ID", Message: "el ID es requerido"}
    }

    doc, err := r.client.Collection("nodes").Doc(id).Get(ctx)
    if err != nil {
        return nil, fmt.Errorf("error al obtener nodo: %v", err)
    }

    var node models.Node
    if err := doc.DataTo(&node); err != nil {
        return nil, fmt.Errorf("error al deserializar nodo: %v", err)
    }
    return &node, nil
}

func (r *FirestoreNodeRepository) Delete(ctx context.Context, id string) error {
    if id == "" {
        return &NodeValidationError{Field: "ID", Message: "el ID es requerido"}
    }

    _, err := r.client.Collection("nodes").Doc(id).Delete(ctx)
    if err != nil {
        return fmt.Errorf("error al eliminar nodo: %v", err)
    }
    return nil
}

func (r *FirestoreNodeRepository) GetFeed(ctx context.Context, filters models.FeedFilters) ([]models.Node, error) {
    // Convert models.FeedFilters to repositories.FeedFilters for validation
    repoFilters := FeedFilters{
        Type:     models.NodeType(filters.Type),
        Category: filters.Category,
        UserID:   filters.UserID,
        Page:     filters.Page,
        PageSize: filters.PageSize,
    }

    if err := ValidateFeedFilters(repoFilters); err != nil {
        return nil, fmt.Errorf("validación de filtros fallida: %w", err)
    }

    query := r.client.Collection("nodes").OrderBy("CreatedAt", firestore.Desc)
    
    if filters.Type != "" {
        query = query.Where("Type", "==", filters.Type)
    }
    
    if filters.LastID != "" {
        lastNode, err := r.Get(ctx, filters.LastID)
        if err != nil {
            return nil, err
        }
        query = query.StartAfter(lastNode.CreatedAt)
    }
    
    if filters.Limit > 0 {
        query = query.Limit(filters.Limit)
    }

    iter := query.Documents(ctx)
    var nodes []models.Node

    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("error al iterar nodos: %v", err)
        }

        var node models.Node
        if err := doc.DataTo(&node); err != nil {
            return nil, fmt.Errorf("error al deserializar nodo: %v", err)
        }
        nodes = append(nodes, node)
    }

    return nodes, nil
}

func (r *FirestoreNodeRepository) AddFollower(ctx context.Context, nodeID, userID string) error {
    if nodeID == "" || userID == "" {
        return &NodeValidationError{Field: "ID", Message: "el ID del nodo y del usuario son requeridos"}
    }

    _, err := r.client.Collection("nodes").Doc(nodeID).Collection("followers").Doc(userID).Set(ctx, map[string]interface{}{
        "userID": userID,
        "followedAt": time.Now(),
    })
    if err != nil {
        return fmt.Errorf("error al agregar seguidor: %v", err)
    }
    return nil
}

func (r *FirestoreNodeRepository) RemoveFollower(ctx context.Context, nodeID, userID string) error {
    if nodeID == "" || userID == "" {
        return &NodeValidationError{Field: "ID", Message: "el ID del nodo y del usuario son requeridos"}
    }

    _, err := r.client.Collection("nodes").Doc(nodeID).Collection("followers").Doc(userID).Delete(ctx)
    if err != nil {
        return fmt.Errorf("error al eliminar seguidor: %v", err)
    }
    return nil
}

func (r *FirestoreNodeRepository) GetFollowers(ctx context.Context, nodeID string) ([]string, error) {
    if nodeID == "" {
        return nil, &NodeValidationError{Field: "NodeID", Message: "el ID del nodo es requerido"}
    }

    iter := r.client.Collection("nodes").Doc(nodeID).Collection("followers").Documents(ctx)
    var followers []string

    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("error al iterar seguidores: %v", err)
        }

        var data map[string]interface{}
        if err := doc.DataTo(&data); err != nil {
            return nil, fmt.Errorf("error al deserializar seguidor: %v", err)
        }
        if userID, ok := data["userID"].(string); ok {
            followers = append(followers, userID)
        }
    }

    return followers, nil
}

func (r *FirestoreNodeRepository) UpdateMetrics(ctx context.Context, nodeID string, metrics models.InteractionMetrics) error {
    if nodeID == "" {
        return &NodeValidationError{Field: "NodeID", Message: "el ID del nodo es requerido"}
    }

    if err := ValidateMetrics(metrics); err != nil {
        return fmt.Errorf("validación de métricas fallida: %w", err)
    }

    _, err := r.client.Collection("nodes").Doc(nodeID).Set(ctx, map[string]interface{}{
        "metrics": metrics,
    }, firestore.MergeAll)
    if err != nil {
        return fmt.Errorf("error al actualizar métricas: %v", err)
    }
    return nil
}

// GetLinkedProducts obtiene los productos vinculados a un nodo
func (r *FirestoreNodeRepository) GetLinkedProducts(ctx context.Context, nodeID string) ([]models.Product, error) {
    node, err := r.Get(ctx, nodeID)
    if err != nil {
        return nil, fmt.Errorf("error getting node: %v", err)
    }

    if len(node.LinkedProducts) == 0 {
        return []models.Product{}, nil
    }

    var products []models.Product
    for _, productID := range node.LinkedProducts {
        doc, err := r.client.Collection("products").Doc(productID).Get(ctx)
        if err != nil {
            continue // Skip if product not found
        }

        var product models.Product
        if err := doc.DataTo(&product); err != nil {
            continue // Skip if unable to parse product
        }
        products = append(products, product)
    }

    return products, nil
}
