package models

type User struct {
    ID            string      `json:"id" firestore:"id"`
    Profile       Profile     `json:"profile" firestore:"profile"`
    FollowedNodes []string    `json:"followedNodes" firestore:"followedNodes"`
    Following     []string    `json:"following" firestore:"following"`
    StoreID       string      `json:"storeId,omitempty" firestore:"storeId,omitempty"`
    Achievements  []Achievement `json:"achievements" firestore:"achievements"`
}

type Profile struct {
    Name        string   `json:"name" firestore:"name"`
    Avatar      string   `json:"avatar" firestore:"avatar"`
    Bio         string   `json:"bio" firestore:"bio"`
    Interests   []string `json:"interests" firestore:"interests"`
}

type Achievement struct {
    ID          string `json:"id" firestore:"id"`
    Name        string `json:"name" firestore:"name"`
    Description string `json:"description" firestore:"description"`
    UnlockedAt  int64  `json:"unlockedAt" firestore:"unlockedAt"`
    Points      int    `json:"points" firestore:"points"`
}
