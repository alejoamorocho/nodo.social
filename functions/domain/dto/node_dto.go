package dto

import (
    "time"
    "github.com/kha0sys/nodo.social/domain/models"
)

// NodeDTO representa los datos de un nodo para transferencia
type NodeDTO struct {
    ID             string          `json:"id"`
    Type           models.NodeType `json:"type"`
    Title          string          `json:"title"`
    Description    string          `json:"description"`
    UserID         string          `json:"userId"`
    LinkedProducts []string        `json:"linkedProducts"`
    Images         []string        `json:"images"`
    CreatedAt      time.Time       `json:"createdAt"`
    UpdatedAt      time.Time       `json:"updatedAt"`
}

// ToModel convierte el DTO a un modelo Node
func (dto *NodeDTO) ToModel() *models.Node {
    return &models.Node{
        ID:             dto.ID,
        Type:           dto.Type,
        Title:          dto.Title,
        Description:    dto.Description,
        UserID:         dto.UserID,
        LinkedProducts: dto.LinkedProducts,
        Images:         dto.Images,
        CreatedAt:      dto.CreatedAt,
        UpdatedAt:      dto.UpdatedAt,
    }
}

// FromModel crea un DTO a partir de un modelo Node
func FromNodeModel(node *models.Node) *NodeDTO {
    return &NodeDTO{
        ID:             node.ID,
        Type:           node.Type,
        Title:          node.Title,
        Description:    node.Description,
        UserID:         node.UserID,
        LinkedProducts: node.LinkedProducts,
        Images:         node.Images,
        CreatedAt:      node.CreatedAt,
        UpdatedAt:      node.UpdatedAt,
    }
}
