package triggers

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
	"github.com/kha0sys/nodo.social/functions/services"
)

// AuthTriggers maneja los eventos de autenticación
type AuthTriggers struct {
	userRepo        repositories.UserRepository
	notificationSvc *services.NotificationService
}

// NewAuthTriggers crea una nueva instancia de AuthTriggers
func NewAuthTriggers(
	userRepo repositories.UserRepository,
	notificationSvc *services.NotificationService,
) *AuthTriggers {
	return &AuthTriggers{
		userRepo:        userRepo,
		notificationSvc: notificationSvc,
	}
}

// OnUserCreated se ejecuta cuando se crea un nuevo usuario en Firebase Auth
func (t *AuthTriggers) OnUserCreated(ctx context.Context, user *auth.UserRecord) error {
	// Crear perfil de usuario en Firestore
	newUser := &models.User{
		ID:          user.UID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		PhotoURL:    user.PhotoURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Points:      0,
		Role:        "user",
		Profile: models.Profile{
			Bio:       "",
			Interests: []string{},
		},
		FollowedNodes: []string{},
		Following:     []string{},
		Achievements:  []models.UserAchievement{},
	}

	if err := t.userRepo.Create(ctx, newUser); err != nil {
		return fmt.Errorf("error creating user profile: %v", err)
	}

	// Enviar notificación de bienvenida
	welcomeNotification := &models.Notification{
		Type:        "welcome",
		UserID:      user.UID,
		Title:       "¡Bienvenido a Nodo Social!",
		Description: "Gracias por unirte a nuestra comunidad. Explora las causas que te interesan y comienza a generar impacto.",
		CreatedAt:   time.Now(),
		Read:        false,
	}

	if err := t.notificationSvc.SendNotification(ctx, welcomeNotification); err != nil {
		// Log el error pero no lo retornamos para no afectar el flujo principal
		fmt.Printf("error sending welcome notification: %v\n", err)
	}

	return nil
}

// OnUserDeleted se ejecuta cuando se elimina un usuario de Firebase Auth
func (t *AuthTriggers) OnUserDeleted(ctx context.Context, user *auth.UserRecord) error {
	// Eliminar perfil de usuario de Firestore
	if err := t.userRepo.Delete(ctx, user.UID); err != nil {
		return fmt.Errorf("error deleting user profile: %v", err)
	}

	return nil
}
