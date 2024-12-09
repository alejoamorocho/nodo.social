package services

import (
    "context"
    "github.com/kha0sys/nodo.social/domain/models"
    "github.com/kha0sys/nodo.social/domain/repositories"
    "github.com/kha0sys/nodo.social/domain/dto"
)

// UserService encapsula la lógica de negocio relacionada con usuarios
type UserService struct {
    userRepo repositories.UserRepository
}

// NewUserService crea una nueva instancia de UserService
func NewUserService(userRepo repositories.UserRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
    }
}

// CreateUser crea un nuevo usuario
func (s *UserService) CreateUser(ctx context.Context, userDTO dto.UserDTO) (*models.User, error) {
    user := &models.User{
        ID:          userDTO.ID,
        Name:        userDTO.Name,
        Email:       userDTO.Email,
        CreatedAt:   userDTO.CreatedAt,
    }

    if err := s.userRepo.Create(ctx, user); err != nil {
        return nil, err
    }

    return user, nil
}

// GetUser obtiene un usuario por su ID
func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
    return s.userRepo.Get(ctx, id)
}

// UpdateUser actualiza un usuario existente
func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
    return s.userRepo.Update(ctx, user)
}

// DeleteUser elimina un usuario
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
    return s.userRepo.Delete(ctx, id)
}
