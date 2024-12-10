package triggers

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
	"github.com/kha0sys/nodo.social/functions/services"
)

type ScheduledTriggers struct {
	client          *firestore.Client
	nodeRepo        repositories.NodeRepository
	userRepo        repositories.UserRepository
	notificationSvc *services.NotificationService
}

func NewScheduledTriggers(
	client *firestore.Client,
	nodeRepo repositories.NodeRepository,
	userRepo repositories.UserRepository,
	notificationSvc *services.NotificationService,
) *ScheduledTriggers {
	return &ScheduledTriggers{
		client:          client,
		nodeRepo:        nodeRepo,
		userRepo:        userRepo,
		notificationSvc: notificationSvc,
	}
}

// DailyCleanup se ejecuta diariamente para realizar tareas de mantenimiento
func (t *ScheduledTriggers) DailyCleanup(ctx context.Context, _ interface{}) error {
	// 1. Limpiar notificaciones antiguas
	if err := t.cleanOldNotifications(ctx); err != nil {
		log.Printf("error cleaning notifications: %v", err)
	}

	// 2. Actualizar estadísticas diarias
	if err := t.updateDailyStats(ctx); err != nil {
		log.Printf("error updating daily stats: %v", err)
	}

	// 3. Limpiar archivos temporales
	if err := t.cleanTempFiles(ctx); err != nil {
		log.Printf("error cleaning temp files: %v", err)
	}

	return nil
}

// WeeklyDigest se ejecuta semanalmente para enviar resúmenes a usuarios
func (t *ScheduledTriggers) WeeklyDigest(ctx context.Context, _ interface{}) error {
	// Obtener usuarios activos
	activeUsersCount, err := t.userRepo.GetActiveUsers(ctx)
	if err != nil {
		return fmt.Errorf("error getting active users count: %v", err)
	}

	// Obtener todos los usuarios activos
	users := make([]*models.User, 0, activeUsersCount)
	iter := t.client.Collection("users").Where("active", "==", true).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var user models.User
		if err := doc.DataTo(&user); err != nil {
			log.Printf("error parsing user data: %v", err)
			continue
		}
		users = append(users, &user)
	}

	weekAgo := time.Now().AddDate(0, 0, -7)

	for _, user := range users {
		// Obtener actividad de la semana para el usuario
		weeklyActivity, err := t.generateWeeklyDigest(ctx, user.ID, weekAgo)
		if err != nil {
			log.Printf("error generating digest for user %s: %v", user.ID, err)
			continue
		}

		// Enviar notificación con el resumen
		notification := &models.Notification{
			UserID:      user.ID,
			Type:        "weekly_digest",
			Title:       "Tu resumen semanal",
			Description: "Aquí está tu resumen de actividad de la semana",
			CreatedAt:   time.Now(),
			Read:        false,
			Data:        weeklyActivity,
		}

		if err := t.notificationSvc.CreateNotification(ctx, notification); err != nil {
			log.Printf("error sending digest notification to user %s: %v", user.ID, err)
		}
	}

	return nil
}

func (t *ScheduledTriggers) cleanOldNotifications(ctx context.Context) error {
	// Eliminar notificaciones más antiguas de 30 días
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	return t.notificationSvc.DeleteOldNotifications(ctx, thirtyDaysAgo)
}

func (t *ScheduledTriggers) updateDailyStats(ctx context.Context) error {
	// Actualizar estadísticas globales
	stats := struct {
		TotalNodes        int     `firestore:"total_nodes"`
		ActiveUsers       int     `firestore:"active_users"`
		TotalInteractions int     `firestore:"total_interactions"`
		UpdatedAt         time.Time `firestore:"updated_at"`
	}{}

	// Contar nodos
	nodes, err := t.nodeRepo.GetTotalNodes(ctx)
	if err != nil {
		return err
	}
	stats.TotalNodes = nodes

	// Contar usuarios activos
	activeUsers, err := t.userRepo.GetActiveUsers(ctx)
	if err != nil {
		return err
	}
	stats.ActiveUsers = activeUsers

	// Guardar estadísticas
	stats.UpdatedAt = time.Now()
	_, err = t.client.Collection("statistics").Doc("daily").Set(ctx, stats)
	return err
}

func (t *ScheduledTriggers) cleanTempFiles(ctx context.Context) error {
    // Obtener lista de archivos temporales más antiguos de 24 horas
    files, err := t.client.Collection("temp_files").
        Where("created_at", "<", time.Now().Add(-24*time.Hour)).
        Documents(ctx).GetAll()
    if err != nil {
        return fmt.Errorf("error getting temp files: %v", err)
    }

    // Eliminar archivos en batch
    batch := t.client.Batch()
    for _, doc := range files {
        batch.Delete(doc.Ref)
    }

    // Commit batch
    _, err = batch.Commit(ctx)
    if err != nil {
        return fmt.Errorf("error deleting temp files: %v", err)
    }

    return nil
}

func (t *ScheduledTriggers) generateWeeklyDigest(ctx context.Context, userID string, weekAgo time.Time) (map[string]interface{}, error) {
	// Obtener actividad del usuario
	activity := make(map[string]interface{})
	
	// Obtener interacciones del usuario
	interactions, err := t.client.Collection("user_activity").
		Where("user_id", "==", userID).
		Where("created_at", ">=", weekAgo).
		Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting user activity: %v", err)
	}

	// Contar interacciones por tipo
	activityCount := make(map[string]int)
	for _, doc := range interactions {
		data := doc.Data()
		if activityType, ok := data["type"].(string); ok {
			activityCount[activityType]++
		}
	}
	activity["interactions"] = activityCount

	// Obtener nodos populares
	popularNodes, err := t.nodeRepo.GetPopularNodes(ctx, 5)
	if err != nil {
		log.Printf("error getting popular nodes: %v", err)
	}

	// Generar resumen
	digest := map[string]interface{}{
		"period": map[string]time.Time{
			"from": weekAgo,
			"to":   time.Now(),
		},
		"activity":     activity,
		"popularNodes": popularNodes,
	}

	return digest, nil
}
