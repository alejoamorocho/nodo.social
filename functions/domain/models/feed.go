package models

// This file has been deprecated.
// FeedFilters has been moved to requests.go

type FeedItem struct {
    ID        string      `json:"id" firestore:"id"`
    Type      string      `json:"type" firestore:"type"`
    NodeID    string      `json:"node_id" firestore:"node_id"`
    UserID    string      `json:"user_id" firestore:"user_id"`
    Content   interface{} `json:"content" firestore:"content"`
    CreatedAt int64       `json:"created_at" firestore:"created_at"`
}

type FeedFilters struct {
    UserID    string `json:"user_id"`
    NodeID    string `json:"node_id"`
    Type      string `json:"type"`
    Category  string `json:"category"`
    StartTime int64  `json:"start_time"`
    EndTime   int64  `json:"end_time"`
    Page      int    `json:"page"`
    PageSize  int    `json:"limit"`
    LastID    string `json:"last_id,omitempty"`
    Cursor    string `json:"cursor,omitempty"`
}

