package services

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
)

// AchievementService maneja los logros de los usuarios
type AchievementService struct {
	client     *firestore.Client
	userRepo   repositories.UserRepository
	collection string
}

// NewAchievementService crea una nueva instancia de AchievementService
func NewAchievementService(client *firestore.Client, userRepo repositories.UserRepository) *AchievementService {
	return &AchievementService{
		client:     client,
		userRepo:   userRepo,
		collection: "achievements",
	}
}

// CheckNodeAchievements verifica y otorga logros relacionados con nodos
func (s *AchievementService) CheckNodeAchievements(ctx context.Context, userID string) error {
	// Obtener el usuario
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	// Verificar logros basados en el número de nodos creados
	nodeCount := len(user.Nodes)
	achievements := make([]models.Achievement, 0)

	switch {
	case nodeCount >= 100:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeCreation,
			Name:        "Maestro de Nodos",
			Description: "Has creado 100 nodos",
			Points:      1000,
			Conditions: []models.Condition{
				{
					Type:     "node_count",
					Value:    100,
					Operator: ">=",
				},
			},
		})
	case nodeCount >= 50:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeCreation,
			Name:        "Experto en Nodos",
			Description: "Has creado 50 nodos",
			Points:      500,
			Conditions: []models.Condition{
				{
					Type:     "node_count",
					Value:    50,
					Operator: ">=",
				},
			},
		})
	case nodeCount >= 25:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeCreation,
			Name:        "Entusiasta de Nodos",
			Description: "Has creado 25 nodos",
			Points:      250,
			Conditions: []models.Condition{
				{
					Type:     "node_count",
					Value:    25,
					Operator: ">=",
				},
			},
		})
	case nodeCount >= 10:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeCreation,
			Name:        "Creador de Nodos",
			Description: "Has creado 10 nodos",
			Points:      100,
			Conditions: []models.Condition{
				{
					Type:     "node_count",
					Value:    10,
					Operator: ">=",
				},
			},
		})
	case nodeCount >= 1:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeCreation,
			Name:        "Primer Nodo",
			Description: "Has creado tu primer nodo",
			Points:      50,
			Conditions: []models.Condition{
				{
					Type:     "node_count",
					Value:    1,
					Operator: ">=",
				},
			},
		})
	}

	// Otorgar logros
	for _, achievement := range achievements {
		if err := s.grantAchievement(ctx, userID, &achievement); err != nil {
			return fmt.Errorf("error granting achievement %s: %v", achievement.Name, err)
		}
	}

	return nil
}

// CheckInteractionAchievements verifica y otorga logros relacionados con interacciones
func (s *AchievementService) CheckInteractionAchievements(ctx context.Context, userID string) error {
	// Obtener el usuario
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	// Verificar logros basados en el número de interacciones
	interactions := user.Metrics.TotalInteractions
	achievements := make([]models.Achievement, 0)

	switch {
	case interactions >= 1000:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeShare,
			Name:        "Super Activo",
			Description: "Has realizado 1000 interacciones",
			Points:      1000,
			Conditions: []models.Condition{
				{
					Type:     "interaction_count",
					Value:    1000,
					Operator: ">=",
				},
			},
		})
	case interactions >= 500:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeShare,
			Name:        "Muy Activo",
			Description: "Has realizado 500 interacciones",
			Points:      500,
			Conditions: []models.Condition{
				{
					Type:     "interaction_count",
					Value:    500,
					Operator: ">=",
				},
			},
		})
	case interactions >= 100:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeShare,
			Name:        "Activo",
			Description: "Has realizado 100 interacciones",
			Points:      100,
			Conditions: []models.Condition{
				{
					Type:     "interaction_count",
					Value:    100,
					Operator: ">=",
				},
			},
		})
	case interactions >= 10:
		achievements = append(achievements, models.Achievement{
			Type:        models.NodeShare,
			Name:        "Principiante",
			Description: "Has realizado 10 interacciones",
			Points:      50,
			Conditions: []models.Condition{
				{
					Type:     "interaction_count",
					Value:    10,
					Operator: ">=",
				},
			},
		})
	}

	// Otorgar logros
	for _, achievement := range achievements {
		if err := s.grantAchievement(ctx, userID, &achievement); err != nil {
			return fmt.Errorf("error granting achievement %s: %v", achievement.Name, err)
		}
	}

	return nil
}

// grantAchievement otorga un logro a un usuario
func (s *AchievementService) grantAchievement(ctx context.Context, userID string, achievement *models.Achievement) error {
	now := time.Now().Unix()
	achievement.ID = fmt.Sprintf("%s_%s", userID, achievement.Type)
	achievement.CreatedAt = now
	achievement.UpdatedAt = now

	_, err := s.client.Collection(s.collection).Doc(achievement.ID).Set(ctx, achievement)
	if err != nil {
		return err
	}

	// Actualizar puntos del usuario
	userPoints := &models.UserPoints{
		UserID:    userID,
		Total:     achievement.Points,
		UpdatedAt: now,
	}

	_, err = s.client.Collection("user_points").Doc(userID).Set(ctx, userPoints, firestore.MergeAll)
	return err
}
