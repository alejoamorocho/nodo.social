package models

import "time"

// UserActivityFilters has been moved to requests.go
// This file is kept for backwards compatibility

type NodeFilters struct {
	CreatorID string    `json:"creatorId,omitempty" firestore:"creatorId,omitempty"`
	Tags      []string  `json:"tags,omitempty" firestore:"tags,omitempty"`
	From      time.Time `json:"from,omitempty" firestore:"from,omitempty"`
	To        time.Time `json:"to,omitempty" firestore:"to,omitempty"`
	Limit     int       `json:"limit,omitempty" firestore:"limit,omitempty"`
	LastID    string    `json:"lastId,omitempty" firestore:"lastId,omitempty"`
}
