package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"google.golang.org/api/iterator"
)

// UserRepository define la interfaz para operaciones con usuarios
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, userID string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID string) error
	GetFollowers(ctx context.Context, userID string) ([]*models.User, error)
	GetTotalUsers(ctx context.Context) (int, error)
	GetActiveUsers(ctx context.Context) (int, error)
}

// FirestoreUserRepository implementa UserRepository usando Firestore
type FirestoreUserRepository struct {
	client     *firestore.Client
	collection string
}

// NewFirestoreUserRepository crea una nueva instancia de FirestoreUserRepository
func NewFirestoreUserRepository(client *firestore.Client) *FirestoreUserRepository {
	return &FirestoreUserRepository{
		client:     client,
		collection: "users",
	}
}

// Create crea un nuevo usuario en Firestore
func (r *FirestoreUserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.client.Collection(r.collection).Doc(user.ID).Set(ctx, user)
	return err
}

// Delete elimina un usuario de Firestore
func (r *FirestoreUserRepository) Delete(ctx context.Context, userID string) error {
	_, err := r.client.Collection(r.collection).Doc(userID).Delete(ctx)
	return err
}

// Get obtiene un usuario de Firestore por su ID
func (r *FirestoreUserRepository) Get(ctx context.Context, userID string) (*models.User, error) {
	doc, err := r.client.Collection(r.collection).Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	user.ID = doc.Ref.ID
	return &user, nil
}

// Update actualiza un usuario existente en Firestore
func (r *FirestoreUserRepository) Update(ctx context.Context, user *models.User) error {
	_, err := r.client.Collection(r.collection).Doc(user.ID).Set(ctx, user, firestore.MergeAll)
	return err
}

// GetFollowers obtiene los seguidores de un usuario
func (r *FirestoreUserRepository) GetFollowers(ctx context.Context, userID string) ([]*models.User, error) {
	var followers []*models.User
	iter := r.client.Collection(r.collection).Where("following", "array-contains", userID).Documents(ctx)
	
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var user models.User
		if err := doc.DataTo(&user); err != nil {
			return nil, err
		}
		user.ID = doc.Ref.ID
		followers = append(followers, &user)
	}

	return followers, nil
}

// GetTotalUsers obtiene el número total de usuarios
func (r *FirestoreUserRepository) GetTotalUsers(ctx context.Context) (int, error) {
	docs, err := r.client.Collection(r.collection).Documents(ctx).GetAll()
	if err != nil {
		return 0, err
	}
	return len(docs), nil
}

// GetActiveUsers obtiene el número de usuarios activos
func (r *FirestoreUserRepository) GetActiveUsers(ctx context.Context) (int, error) {
	docs, err := r.client.Collection(r.collection).
		Where("active", "==", true).
		Documents(ctx).
		GetAll()
	if err != nil {
		return 0, err
	}
	return len(docs), nil
}
