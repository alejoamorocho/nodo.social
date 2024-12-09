package repositories

import (
    "context"
    "fmt"
    "time"

    "cloud.google.com/go/firestore"
    "github.com/kha0sys/nodo.social/domain/models"
    "google.golang.org/api/iterator"
)

type UserRepository interface {
    Create(ctx context.Context, user *models.User) error
    Get(ctx context.Context, id string) (*models.User, error)
    Update(ctx context.Context, user *models.User) error
    Delete(ctx context.Context, id string) error
    AddFollower(ctx context.Context, userID, followerID string) error
    RemoveFollower(ctx context.Context, userID, followerID string) error
    GetFollowers(ctx context.Context, userID string) ([]string, error)
    GetActivity(ctx context.Context, userID string, filters models.UserActivityFilters) ([]interface{}, error)
}

type FirestoreUserRepository struct {
    client *firestore.Client
}

func NewFirestoreUserRepository(client *firestore.Client) *FirestoreUserRepository {
    return &FirestoreUserRepository{
        client: client,
    }
}

func (r *FirestoreUserRepository) Create(ctx context.Context, user *models.User) error {
    _, err := r.client.Collection("users").Doc(user.ID).Set(ctx, user)
    if err != nil {
        return fmt.Errorf("error al crear usuario: %v", err)
    }
    return nil
}

func (r *FirestoreUserRepository) Get(ctx context.Context, id string) (*models.User, error) {
    doc, err := r.client.Collection("users").Doc(id).Get(ctx)
    if err != nil {
        return nil, fmt.Errorf("error al obtener usuario: %v", err)
    }
    var user models.User
    if err := doc.DataTo(&user); err != nil {
        return nil, fmt.Errorf("error al deserializar usuario: %v", err)
    }
    return &user, nil
}

func (r *FirestoreUserRepository) Update(ctx context.Context, user *models.User) error {
    _, err := r.client.Collection("users").Doc(user.ID).Set(ctx, user)
    if err != nil {
        return fmt.Errorf("error al actualizar usuario: %v", err)
    }
    return nil
}

func (r *FirestoreUserRepository) Delete(ctx context.Context, id string) error {
    _, err := r.client.Collection("users").Doc(id).Delete(ctx)
    if err != nil {
        return fmt.Errorf("error al eliminar usuario: %v", err)
    }
    return nil
}

func (r *FirestoreUserRepository) AddFollower(ctx context.Context, userID, followerID string) error {
    _, err := r.client.Collection("users").Doc(userID).Collection("followers").Doc(followerID).Set(ctx, map[string]interface{}{
        "followerID": followerID,
        "followedAt": time.Now(),
    })
    if err != nil {
        return fmt.Errorf("error al agregar seguidor: %v", err)
    }
    return nil
}

func (r *FirestoreUserRepository) RemoveFollower(ctx context.Context, userID, followerID string) error {
    _, err := r.client.Collection("users").Doc(userID).Collection("followers").Doc(followerID).Delete(ctx)
    if err != nil {
        return fmt.Errorf("error al eliminar seguidor: %v", err)
    }
    return nil
}

func (r *FirestoreUserRepository) GetFollowers(ctx context.Context, userID string) ([]string, error) {
    iter := r.client.Collection("users").Doc(userID).Collection("followers").Documents(ctx)
    var followers []string

    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("error al obtener seguidores: %v", err)
        }

        var data map[string]interface{}
        if err := doc.DataTo(&data); err != nil {
            return nil, fmt.Errorf("error al deserializar seguidor: %v", err)
        }
        if followerID, ok := data["followerID"].(string); ok {
            followers = append(followers, followerID)
        }
    }

    return followers, nil
}

func (r *FirestoreUserRepository) GetActivity(ctx context.Context, userID string, filters models.UserActivityFilters) ([]interface{}, error) {
    query := r.client.Collection("users").Doc(userID).Collection("activity").OrderBy("timestamp", firestore.Desc)
    
    if filters.LastTimestamp > 0 {
        query = query.StartAfter(time.Unix(filters.LastTimestamp, 0))
    }
    
    if filters.Type != "" && filters.Type != "all" {
        query = query.Where("type", "==", filters.Type)
    }
    
    if filters.Limit > 0 {
        query = query.Limit(filters.Limit)
    }

    iter := query.Documents(ctx)
    var activities []interface{}

    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("error al obtener actividades: %v", err)
        }

        var activity map[string]interface{}
        if err := doc.DataTo(&activity); err != nil {
            return nil, fmt.Errorf("error al deserializar actividad: %v", err)
        }
        activities = append(activities, activity)
    }

    return activities, nil
}
