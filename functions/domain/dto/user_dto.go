package dto

import (
    "time"
    "github.com/kha0sys/nodo.social/domain/models"
)

// UserDTO representa los datos de un usuario para transferencia
type UserDTO struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"createdAt"`
}

// ToModel convierte el DTO a un modelo User
func (dto *UserDTO) ToModel() *models.User {
    return &models.User{
        ID:        dto.ID,
        Name:      dto.Name,
        Email:     dto.Email,
        CreatedAt: dto.CreatedAt,
    }
}

// FromUserModel crea un DTO a partir de un modelo User
func FromUserModel(user *models.User) *UserDTO {
    return &UserDTO{
        ID:        user.ID,
        Name:      user.Name,
        Email:     user.Email,
        CreatedAt: user.CreatedAt,
    }
}
