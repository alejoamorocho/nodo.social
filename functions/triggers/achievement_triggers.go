package triggers

import (
	"context"
	"log"

	"cloud.google.com/go/functions/metadata"
	"github.com/kha0sys/nodo.social/functions/internal/firebase"
)

type AchievementTriggers struct {
	// TODO: Add achievement repository
}

func NewAchievementTriggers() *AchievementTriggers {
	return &AchievementTriggers{}
}

// OnAchievement es el trigger que se ejecuta cuando un usuario obtiene un logro
// Firestore trigger: projects/{project}/databases/{database}/documents/users/{userId}/achievements/{achievementId}
func (t *AchievementTriggers) OnAchievement(ctx context.Context, e firebase.FirestoreEvent) error {
	// Verificar metadata del contexto (útil para logging y debugging)
	if _, err := metadata.FromContext(ctx); err != nil {
		log.Printf("Warning: Could not get metadata: %v", err)
	}

	userID := e.Value.Fields["userId"].StringValue
	achievementID := e.Value.Fields["achievementId"].StringValue

	// TODO: Implementar lógica de rankings
	// 1. Obtener el valor de puntos del logro
	// 2. Actualizar el ranking del usuario
	// 3. Recalcular posiciones si es necesario

	log.Printf("Achievement %s unlocked for user %s", achievementID, userID)
	return nil
}

// Achievement representa un logro en el sistema
type Achievement struct {
	ID          string `json:"id" firestore:"id"`
	Name        string `json:"name" firestore:"name"`
	Description string `json:"description" firestore:"description"`
	Points      int    `json:"points" firestore:"points"`
	Category    string `json:"category" firestore:"category"`
}

// UserRanking representa el ranking de un usuario
type UserRanking struct {
	UserID   string `json:"userId" firestore:"userId"`
	Points   int    `json:"points" firestore:"points"`
	Position int    `json:"position" firestore:"position"`
	Level    int    `json:"level" firestore:"level"`
}

