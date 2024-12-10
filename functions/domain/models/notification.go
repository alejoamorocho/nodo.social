package models

import "time"

type Notification struct {
    ID          string                 `json:"id" firestore:"id"`
    Type        string                 `json:"type" firestore:"type"`
    UserID      string                 `json:"user_id" firestore:"user_id"`
    Title       string                 `json:"title" firestore:"title"`
    Description string                 `json:"description" firestore:"description"`
    Data        map[string]interface{} `json:"data,omitempty" firestore:"data,omitempty"`
    CreatedAt   time.Time             `json:"created_at" firestore:"created_at"`
    Read        bool                  `json:"read" firestore:"read"`
    ReadAt      *time.Time            `json:"read_at,omitempty" firestore:"read_at,omitempty"`
}
