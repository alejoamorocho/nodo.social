package repositories

import (
    "context"
    "fmt"
    "github.com/kha0sys/nodo.social/domain/models"
)

// NodeRepository define las operaciones disponibles para la persistencia de nodos
type NodeRepository interface {
    Create(ctx context.Context, node *models.Node) error
    Get(ctx context.Context, id string) (*models.Node, error)
    Update(ctx context.Context, node *models.Node) error
    Delete(ctx context.Context, id string) error
    GetFeed(ctx context.Context, filters models.FeedFilters) ([]models.Node, error)
    GetLinkedProducts(ctx context.Context, nodeID string) ([]models.Product, error)
    AddFollower(ctx context.Context, nodeID, userID string) error
    RemoveFollower(ctx context.Context, nodeID, userID string) error
    GetFollowers(ctx context.Context, nodeID string) ([]string, error)
    UpdateMetrics(ctx context.Context, nodeID string, metrics models.InteractionMetrics) error
}

// FeedFilters define los filtros disponibles para el feed de nodos
type FeedFilters struct {
    Type      models.NodeType
    Category  string
    UserID    string
    Page      int
    PageSize  int
}

// NodeValidationError representa un error de validación de nodo
type NodeValidationError struct {
    Field   string
    Message string
}

func (e *NodeValidationError) Error() string {
    return fmt.Sprintf("error de validación en campo %s: %s", e.Field, e.Message)
}

// ValidateNode valida los campos requeridos de un nodo
func ValidateNode(node *models.Node) error {
    if node.Name == "" {
        return &NodeValidationError{Field: "Name", Message: "el nombre es requerido"}
    }
    if node.Type == "" {
        return &NodeValidationError{Field: "Type", Message: "el tipo es requerido"}
    }
    if node.Description == "" {
        return &NodeValidationError{Field: "Description", Message: "la descripción es requerida"}
    }
    // Validar que el tipo sea uno de los permitidos
    validTypes := map[models.NodeType]bool{
        models.Social:        true,
        models.Environmental: true,
        models.Animal:        true,
    }
    if !validTypes[node.Type] {
        return &NodeValidationError{Field: "Type", Message: "tipo de nodo inválido"}
    }
    return nil
}

// ValidateMetrics valida las métricas de interacción
func ValidateMetrics(metrics models.InteractionMetrics) error {
    if metrics.Views < 0 {
        return &NodeValidationError{Field: "Views", Message: "las vistas no pueden ser negativas"}
    }
    if metrics.Likes < 0 {
        return &NodeValidationError{Field: "Likes", Message: "los likes no pueden ser negativos"}
    }
    if metrics.Shares < 0 {
        return &NodeValidationError{Field: "Shares", Message: "los compartidos no pueden ser negativos"}
    }
    if metrics.Comments < 0 {
        return &NodeValidationError{Field: "Comments", Message: "los comentarios no pueden ser negativos"}
    }
    return nil
}

// ValidateFeedFilters valida los filtros del feed
func ValidateFeedFilters(filters FeedFilters) error {
    if filters.Page < 0 {
        return &NodeValidationError{Field: "Page", Message: "la página no puede ser negativa"}
    }
    if filters.PageSize <= 0 {
        return &NodeValidationError{Field: "PageSize", Message: "el tamaño de página debe ser positivo"}
    }
    if filters.PageSize > 100 {
        return &NodeValidationError{Field: "PageSize", Message: "el tamaño de página no puede ser mayor a 100"}
    }
    if filters.Type != "" {
        validTypes := map[models.NodeType]bool{
            models.Social:        true,
            models.Environmental: true,
            models.Animal:        true,
        }
        if !validTypes[filters.Type] {
            return &NodeValidationError{Field: "Type", Message: "tipo de nodo inválido en filtros"}
        }
    }
    return nil
}
