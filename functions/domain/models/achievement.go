package models

// AchievementType representa el tipo de logro
type AchievementType string

const (
    NodeCreation  AchievementType = "NODE_CREATION"
    NodeFollow    AchievementType = "NODE_FOLLOW"
    ProductLink   AchievementType = "PRODUCT_LINK"
    NodeUpdate    AchievementType = "NODE_UPDATE"
    NodeShare     AchievementType = "NODE_SHARE"
)

// Achievement representa un logro que puede ser desbloqueado por un usuario
type Achievement struct {
    ID          string          `json:"id" firestore:"id"`
    Type        AchievementType `json:"type" firestore:"type"`
    Name        string          `json:"name" firestore:"name"`
    Description string          `json:"description" firestore:"description"`
    Points      int             `json:"points" firestore:"points"`
    Conditions  []Condition    `json:"conditions" firestore:"conditions"`
    CreatedAt   int64          `json:"created_at" firestore:"created_at"`
    UpdatedAt   int64          `json:"updated_at" firestore:"updated_at"`
}

// UserAchievement has been moved to user.go

type Condition struct {
    Type      string      `json:"type" firestore:"type"`
    Value     interface{} `json:"value" firestore:"value"`
    Operator  string      `json:"operator" firestore:"operator"`
}

type UserPoints struct {
    UserID    string `json:"user_id" firestore:"user_id"`
    Total     int    `json:"total" firestore:"total"`
    UpdatedAt int64  `json:"updated_at" firestore:"updated_at"`
}
