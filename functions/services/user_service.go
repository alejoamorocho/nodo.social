package services

import (
	"context"
	"fmt"

	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
)

// UserService maneja la lógica de negocio relacionada con usuarios
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
func (s *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	return user, nil
}

// GetUser obtiene un usuario por su ID
func (s *UserService) GetUser(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return user, nil
}

// UpdateUser actualiza un usuario existente
func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

// DeleteUser elimina un usuario
func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}

// GetFollowers obtiene los seguidores de un usuario
func (s *UserService) GetFollowers(ctx context.Context, userID string) ([]*models.User, error) {
	followers, err := s.userRepo.GetFollowers(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting followers: %v", err)
	}
	return followers, nil
}

// GetTotalUsers obtiene el número total de usuarios
func (s *UserService) GetTotalUsers(ctx context.Context) (int, error) {
	count, err := s.userRepo.GetTotalUsers(ctx)
	if err != nil {
		return 0, fmt.Errorf("error getting total users: %v", err)
	}
	return count, nil
}

// GetActiveUsers obtiene el número de usuarios activos
func (s *UserService) GetActiveUsers(ctx context.Context) (int, error) {
	count, err := s.userRepo.GetActiveUsers(ctx)
	if err != nil {
		return 0, fmt.Errorf("error getting active users: %v", err)
	}
	return count, nil
}
