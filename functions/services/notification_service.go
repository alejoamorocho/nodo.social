package services

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
)

// NotificationService maneja el envío de notificaciones
type NotificationService struct {
	fcmClient    *messaging.Client
	userRepo     repositories.UserRepository
	notificationRepo repositories.NotificationRepository
}

// NewNotificationService crea una nueva instancia de NotificationService
func NewNotificationService(
	app *firebase.App,
	userRepo repositories.UserRepository,
	notificationRepo repositories.NotificationRepository,
) (*NotificationService, error) {
	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing FCM client: %v", err)
	}

	return &NotificationService{
		fcmClient:    fcmClient,
		userRepo:     userRepo,
		notificationRepo: notificationRepo,
	}, nil
}

// SendNotification envía una notificación a un usuario
func (s *NotificationService) SendNotification(ctx context.Context, notification *models.Notification) error {
	// Obtener el token FCM del usuario desde Firestore
	user, err := s.userRepo.Get(ctx, notification.UserID)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	// Guardar la notificación en Firestore
	if err := s.notificationRepo.Create(ctx, notification); err != nil {
		return fmt.Errorf("error saving notification: %v", err)
	}

	if user.FCMToken == "" {
		// Si el usuario no tiene token FCM, solo guardamos la notificación
		return nil
	}

	// Crear mensaje FCM
	message := &messaging.Message{
		Token: user.FCMToken,
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Description,
		},
		Data: map[string]string{
			"type": notification.Type,
			"id":   notification.ID,
		},
	}

	// Si hay datos adicionales, los agregamos al mensaje
	if notification.Data != nil {
		for k, v := range notification.Data {
			if str, ok := v.(string); ok {
				message.Data[k] = str
			}
		}
	}

	// Enviar notificación FCM
	_, err = s.fcmClient.Send(ctx, message)
	if err != nil {
		// Si falla el envío FCM, al menos la notificación ya está guardada
		fmt.Printf("error sending FCM notification: %v\n", err)
	}

	return nil
}

// SendMulticastNotification envía una notificación a múltiples usuarios
func (s *NotificationService) SendMulticastNotification(ctx context.Context, userIDs []string, notification *models.Notification) error {
	tokens := make([]string, 0, len(userIDs))
	notificationsByToken := make(map[string]*models.Notification)

	// Obtener tokens FCM de todos los usuarios y crear notificaciones
	for _, userID := range userIDs {
		user, err := s.userRepo.Get(ctx, userID)
		if err != nil {
			continue // Skip users that can't be found
		}

		// Crear una copia de la notificación para este usuario
		userNotification := *notification
		userNotification.UserID = userID

		// Guardar la notificación
		if err := s.notificationRepo.Create(ctx, &userNotification); err != nil {
			fmt.Printf("error saving notification for user %s: %v\n", userID, err)
			continue
		}

		if user.FCMToken != "" {
			tokens = append(tokens, user.FCMToken)
			notificationsByToken[user.FCMToken] = &userNotification
		}
	}

	if len(tokens) == 0 {
		return nil // No hay tokens FCM, pero las notificaciones se guardaron
	}

	// Crear mensaje multicast
	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Description,
		},
		Data: map[string]string{
			"type": notification.Type,
		},
	}

	// Si hay datos adicionales, los agregamos al mensaje
	if notification.Data != nil {
		for k, v := range notification.Data {
			if str, ok := v.(string); ok {
				message.Data[k] = str
			}
		}
	}

	// Enviar notificación multicast
	response, err := s.fcmClient.SendMulticast(ctx, message)
	if err != nil {
		return fmt.Errorf("error sending multicast notification: %v", err)
	}

	if response.FailureCount > 0 {
		fmt.Printf("%d notifications failed to send\n", response.FailureCount)
	}

	return nil
}

// MarkAsRead marca una notificación como leída
func (s *NotificationService) MarkAsRead(ctx context.Context, notificationID string) error {
	return s.notificationRepo.MarkAsRead(ctx, notificationID)
}

// GetUnreadNotifications obtiene todas las notificaciones no leídas de un usuario
func (s *NotificationService) GetUnreadNotifications(ctx context.Context, userID string) ([]*models.Notification, error) {
	return s.notificationRepo.GetUnreadByUser(ctx, userID)
}

// CreateNotification crea una nueva notificación
func (s *NotificationService) CreateNotification(ctx context.Context, notification *models.Notification) error {
	return s.notificationRepo.Create(ctx, notification)
}

// DeleteOldNotifications elimina las notificaciones más antiguas que la fecha especificada
func (s *NotificationService) DeleteOldNotifications(ctx context.Context, olderThan time.Time) error {
	notifications, err := s.notificationRepo.GetOlderThan(ctx, olderThan)
	if err != nil {
		return err
	}

	for _, notification := range notifications {
		if err := s.notificationRepo.Delete(ctx, notification.ID); err != nil {
			// Log error but continue deleting others
			fmt.Printf("error deleting notification %s: %v\n", notification.ID, err)
		}
	}

	return nil
}
