package models

type NodeType string

const (
    Social      NodeType = "social"
    Environmental NodeType = "environmental"
    Animal       NodeType = "animal"
)

type Node struct {
    ID                    string            `json:"id" firestore:"id"`
    Type                  NodeType          `json:"type" firestore:"type"`
    Content               Content           `json:"content" firestore:"content"`
    Updates               []Update          `json:"updates" firestore:"updates"`
    FollowersCount        int              `json:"followersCount" firestore:"followersCount"`
    LinkedProducts        []string         `json:"linkedProducts" firestore:"linkedProducts"`
    ProductApprovalConfig ApprovalConfig   `json:"productApprovalConfig" firestore:"productApprovalConfig"`
    Metrics              InteractionMetrics `json:"metrics" firestore:"metrics"`
}

type Content struct {
    Text      string   `json:"text" firestore:"text"`
    MediaURLs []string `json:"mediaUrls" firestore:"mediaUrls"`
}

type Update struct {
    ID        string    `json:"id" firestore:"id"`
    Content   Content   `json:"content" firestore:"content"`
    Timestamp int64     `json:"timestamp" firestore:"timestamp"`
}

type ApprovalConfig struct {
    RequiresApproval bool `json:"requiresApproval" firestore:"requiresApproval"`
    AutoApprove      bool `json:"autoApprove" firestore:"autoApprove"`
}

type InteractionMetrics struct {
    Views        int `json:"views" firestore:"views"`
    Likes        int `json:"likes" firestore:"likes"`
    Shares       int `json:"shares" firestore:"shares"`
    Interactions int `json:"interactions" firestore:"interactions"`
}
