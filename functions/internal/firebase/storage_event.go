package firebase

import "time"

// StorageEvent representa un evento de Cloud Storage
type StorageEvent struct {
	Bucket         string    `json:"bucket"`
	Name           string    `json:"name"`
	Metageneration string    `json:"metageneration"`
	TimeCreated    time.Time `json:"timeCreated"`
	Updated        time.Time `json:"updated"`
	Size           int64     `json:"size"`
	MD5Hash        string    `json:"md5Hash"`
	ContentType    string    `json:"contentType"`
	ContentEncoding string    `json:"contentEncoding,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
}
