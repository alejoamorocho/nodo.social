package firestore

import (
    "context"
    "time"
    
    "cloud.google.com/go/firestore"
    "github.com/kha0sys/nodo.social/functions/domain/models"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type AchievementRepository struct {
    client *firestore.Client
}

func NewAchievementRepository(client *firestore.Client) *AchievementRepository {
    return &AchievementRepository{
        client: client,
    }
}

func (r *AchievementRepository) CreateAchievement(ctx context.Context, achievement *models.Achievement) error {
    achievement.CreatedAt = time.Now().Unix()
    achievement.UpdatedAt = achievement.CreatedAt
    
    _, err := r.client.Collection("achievements").Doc(achievement.ID).Set(ctx, achievement)
    return err
}

func (r *AchievementRepository) GetAchievement(ctx context.Context, id string) (*models.Achievement, error) {
    doc, err := r.client.Collection("achievements").Doc(id).Get(ctx)
    if err != nil {
        if status.Code(err) == codes.NotFound {
            return nil, nil
        }
        return nil, err
    }

    var achievement models.Achievement
    if err := doc.DataTo(&achievement); err != nil {
        return nil, err
    }

    return &achievement, nil
}

func (r *AchievementRepository) ListAchievements(ctx context.Context) ([]*models.Achievement, error) {
    docs, err := r.client.Collection("achievements").Documents(ctx).GetAll()
    if err != nil {
        return nil, err
    }

    achievements := make([]*models.Achievement, 0, len(docs))
    for _, doc := range docs {
        var achievement models.Achievement
        if err := doc.DataTo(&achievement); err != nil {
            return nil, err
        }
        achievements = append(achievements, &achievement)
    }

    return achievements, nil
}

func (r *AchievementRepository) UnlockAchievement(ctx context.Context, userAchievement *models.UserAchievement) error {
    // Verificar si ya está desbloqueado
    docID := userAchievement.UserID + "_" + userAchievement.AchievementID
    doc, err := r.client.Collection("user_achievements").Doc(docID).Get(ctx)
    if err == nil && doc.Exists() {
        return nil // Ya está desbloqueado
    }

    // Desbloquear logro
    _, err = r.client.Collection("user_achievements").Doc(docID).Set(ctx, userAchievement)
    if err != nil {
        return err
    }

    // Actualizar puntos
    return r.UpdateUserPoints(ctx, userAchievement.UserID, userAchievement.Points)
}

func (r *AchievementRepository) GetUserAchievements(ctx context.Context, userID string) ([]*models.UserAchievement, error) {
    docs, err := r.client.Collection("user_achievements").Where("user_id", "==", userID).Documents(ctx).GetAll()
    if err != nil {
        return nil, err
    }

    achievements := make([]*models.UserAchievement, 0, len(docs))
    for _, doc := range docs {
        var achievement models.UserAchievement
        if err := doc.DataTo(&achievement); err != nil {
            return nil, err
        }
        achievements = append(achievements, &achievement)
    }

    return achievements, nil
}

func (r *AchievementRepository) UpdateUserPoints(ctx context.Context, userID string, points int) error {
    docRef := r.client.Collection("user_points").Doc(userID)
    
    err := r.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
        doc, err := tx.Get(docRef)
        if err != nil {
            if status.Code(err) != codes.NotFound {
                return err
            }
            // Crear nuevo registro de puntos
            return tx.Set(docRef, &models.UserPoints{
                UserID:    userID,
                Total:     points,
                UpdatedAt: time.Now().Unix(),
            })
        }

        var userPoints models.UserPoints
        if err := doc.DataTo(&userPoints); err != nil {
            return err
        }

        // Actualizar puntos existentes
        userPoints.Total += points
        userPoints.UpdatedAt = time.Now().Unix()
        
        return tx.Set(docRef, userPoints)
    })

    return err
}

func (r *AchievementRepository) GetUserPoints(ctx context.Context, userID string) (*models.UserPoints, error) {
    doc, err := r.client.Collection("user_points").Doc(userID).Get(ctx)
    if err != nil {
        if status.Code(err) == codes.NotFound {
            return &models.UserPoints{
                UserID: userID,
                Total:  0,
            }, nil
        }
        return nil, err
    }

    var points models.UserPoints
    if err := doc.DataTo(&points); err != nil {
        return nil, err
    }

    return &points, nil
}

