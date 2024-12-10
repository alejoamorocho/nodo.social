package models

import (
	"time"
)

type User struct {
    ID            string           `json:"id" firestore:"id"`
    Email         string           `json:"email" firestore:"email"`
    DisplayName   string           `json:"displayName" firestore:"displayName"`
    PhotoURL      string           `json:"photoUrl" firestore:"photoUrl"`
    Profile       Profile          `json:"profile" firestore:"profile"`
    FollowedNodes []string         `json:"followedNodes" firestore:"followedNodes"`
    Following     []string         `json:"following" firestore:"following"`
    StoreID       string           `json:"storeId,omitempty" firestore:"storeId,omitempty"`
    Achievements  []UserAchievement `json:"achievements" firestore:"achievements"`
    Points        int              `json:"points" firestore:"points"`
    Role          string           `json:"role" firestore:"role"`
    FCMToken      string           `json:"fcmToken,omitempty" firestore:"fcmToken,omitempty"`
    CreatedAt     time.Time        `json:"createdAt" firestore:"createdAt"`
    UpdatedAt     time.Time        `json:"updatedAt" firestore:"updatedAt"`
    Nodes         []string         `json:"nodes" firestore:"nodes"`
    Metrics       UserMetrics      `json:"metrics" firestore:"metrics"`
}

type Profile struct {
    Bio         string   `json:"bio" firestore:"bio"`
    Interests   []string `json:"interests" firestore:"interests"`
}

type UserAchievement struct {
    ID            string          `json:"id" firestore:"id"`
    UserID        string          `json:"userId" firestore:"userId"`
    AchievementID string          `json:"achievementId" firestore:"achievementId"`
    Type          AchievementType `json:"type" firestore:"type"`
    Points        int             `json:"points" firestore:"points"`
    UnlockedAt    time.Time       `json:"unlockedAt" firestore:"unlockedAt"`
}

type UserMetrics struct {
    TotalInteractions int `json:"totalInteractions" firestore:"totalInteractions"`
    Likes            int `json:"likes" firestore:"likes"`
    Comments         int `json:"comments" firestore:"comments"`
    Shares           int `json:"shares" firestore:"shares"`
    Views            int `json:"views" firestore:"views"`
}
