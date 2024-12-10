package repositories

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
)

// NotificationRepository define la interfaz para operaciones con notificaciones
type NotificationRepository interface {
	Create(ctx context.Context, notification *models.Notification) error
	MarkAsRead(ctx context.Context, notificationID string) error
	GetUnreadByUser(ctx context.Context, userID string) ([]*models.Notification, error)
	GetOlderThan(ctx context.Context, olderThan time.Time) ([]*models.Notification, error)
	Delete(ctx context.Context, notificationID string) error
}

// FirestoreNotificationRepository implementa NotificationRepository usando Firestore
type FirestoreNotificationRepository struct {
	client     *firestore.Client
	collection string
}

// NewFirestoreNotificationRepository crea una nueva instancia de FirestoreNotificationRepository
func NewFirestoreNotificationRepository(client *firestore.Client) *FirestoreNotificationRepository {
	return &FirestoreNotificationRepository{
		client:     client,
		collection: "notifications",
	}
}

// Create guarda una nueva notificación en Firestore
func (r *FirestoreNotificationRepository) Create(ctx context.Context, notification *models.Notification) error {
	docRef := r.client.Collection(r.collection).NewDoc()
	notification.ID = docRef.ID
	_, err := docRef.Set(ctx, notification)
	return err
}

// MarkAsRead marca una notificación como leída
func (r *FirestoreNotificationRepository) MarkAsRead(ctx context.Context, notificationID string) error {
	_, err := r.client.Collection(r.collection).Doc(notificationID).Update(ctx, []firestore.Update{
		{Path: "read", Value: true},
		{Path: "read_at", Value: firestore.ServerTimestamp},
	})
	return err
}

// GetUnreadByUser obtiene todas las notificaciones no leídas de un usuario
func (r *FirestoreNotificationRepository) GetUnreadByUser(ctx context.Context, userID string) ([]*models.Notification, error) {
	iter := r.client.Collection(r.collection).
		Where("user_id", "==", userID).
		Where("read", "==", false).
		OrderBy("created_at", firestore.Desc).
		Documents(ctx)

	var notifications []*models.Notification
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var notification models.Notification
		if err := doc.DataTo(&notification); err != nil {
			continue
		}
		notification.ID = doc.Ref.ID
		notifications = append(notifications, &notification)
	}

	return notifications, nil
}

// GetOlderThan obtiene todas las notificaciones más antiguas que la fecha especificada
func (r *FirestoreNotificationRepository) GetOlderThan(ctx context.Context, olderThan time.Time) ([]*models.Notification, error) {
	iter := r.client.Collection(r.collection).
		Where("created_at", "<", olderThan).
		Documents(ctx)

	var notifications []*models.Notification
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var notification models.Notification
		if err := doc.DataTo(&notification); err != nil {
			continue
		}
		notification.ID = doc.Ref.ID
		notifications = append(notifications, &notification)
	}

	return notifications, nil
}

// Delete elimina una notificación
func (r *FirestoreNotificationRepository) Delete(ctx context.Context, notificationID string) error {
	_, err := r.client.Collection(r.collection).Doc(notificationID).Delete(ctx)
	return err
}
