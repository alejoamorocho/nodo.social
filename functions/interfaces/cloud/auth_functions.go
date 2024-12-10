package cloud

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
	"github.com/kha0sys/nodo.social/functions/services"
)

// authEventData representa los datos del evento de autenticación
type authEventData struct {
	UID         string    `json:"uid"`
	Email       string    `json:"email"`
	DisplayName string    `json:"displayName"`
	PhotoURL    string    `json:"photoURL"`
	CreatedAt   time.Time `json:"createdAt"`
}

// initAuthClient inicializa el cliente de autenticación de Firebase
func initAuthClient(ctx context.Context) (*auth.Client, error) {
	app, err := InitFirebase(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return client, nil
}

// initFirestoreClient inicializa el cliente de Firestore
func initFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	app, err := InitFirebase(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Firestore client: %v", err)
	}

	return client, nil
}

// OnUserCreated se ejecuta cuando se crea un nuevo usuario en Firebase Auth
func OnUserCreated(ctx context.Context, event *cloudevents.Event) error {
	var eventData authEventData
	if err := event.DataAs(&eventData); err != nil {
		return fmt.Errorf("error parsing event data: %v", err)
	}

	// Obtener el cliente de Firestore
	client, err := initFirestoreClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Crear el usuario en Firestore
	userRepo := repositories.NewFirestoreUserRepository(client)
	userService := services.NewUserService(userRepo)

	user := &models.User{
		ID:          eventData.UID,
		Email:       eventData.Email,
		DisplayName: eventData.DisplayName,
		PhotoURL:    eventData.PhotoURL,
		CreatedAt:   eventData.CreatedAt,
	}

	createdUser, err := userService.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error creating user in Firestore: %v", err)
	}

	log.Printf("User created successfully: %s", createdUser.ID)
	return nil
}

// OnUserDeleted se ejecuta cuando se elimina un usuario de Firebase Auth
func OnUserDeleted(ctx context.Context, event *cloudevents.Event) error {
	var eventData authEventData
	if err := event.DataAs(&eventData); err != nil {
		return fmt.Errorf("error parsing event data: %v", err)
	}

	// Obtener el cliente de Firestore
	client, err := initFirestoreClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Eliminar el usuario de Firestore
	userRepo := repositories.NewFirestoreUserRepository(client)
	userService := services.NewUserService(userRepo)

	if err := userService.DeleteUser(ctx, eventData.UID); err != nil {
		return fmt.Errorf("error deleting user from Firestore: %v", err)
	}

	log.Printf("User deleted successfully: %s", eventData.UID)
	return nil
}

// handleAuthEvent maneja los eventos de autenticación de Firebase
func handleAuthEvent(w http.ResponseWriter, r *http.Request, handler func(context.Context, *cloudevents.Event) error) {
	event, err := cloudevents.NewEventFromHTTPRequest(r)
	if err != nil {
		log.Printf("Error parsing CloudEvent: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler(r.Context(), event); err != nil {
		log.Printf("Error handling event: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func init() {
	functions.HTTP("OnUserCreated", func(w http.ResponseWriter, r *http.Request) {
		handleAuthEvent(w, r, OnUserCreated)
	})

	functions.HTTP("OnUserDeleted", func(w http.ResponseWriter, r *http.Request) {
		handleAuthEvent(w, r, OnUserDeleted)
	})
}
