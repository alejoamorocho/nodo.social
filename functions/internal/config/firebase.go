package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// ServiceAccount represents the Firebase service account structure
type ServiceAccount struct {
	Type                    string `json:"type"`
	ProjectID              string `json:"project_id"`
	PrivateKeyID           string `json:"private_key_id"`
	PrivateKey             string `json:"private_key"`
	ClientEmail            string `json:"client_email"`
	ClientID               string `json:"client_id"`
	AuthURI                string `json:"auth_uri"`
	TokenURI               string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL      string `json:"client_x509_cert_url"`
}

// InitializeFirebase initializes the Firebase Admin SDK using environment variables
func InitializeFirebase() (*firebase.App, error) {
	ctx := context.Background()

	// Create service account from environment variables
	privateKey := os.Getenv("FIREBASE_PRIVATE_KEY")
	// Remove quotes if present
	privateKey = strings.Trim(privateKey, "\"")
	// Replace literal \n with actual newlines
	privateKey = strings.ReplaceAll(privateKey, "\\n", "\n")

	sa := ServiceAccount{
		Type:                    "service_account",
		ProjectID:              os.Getenv("FIREBASE_PROJECT_ID"),
		PrivateKeyID:           os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:             privateKey,
		ClientEmail:            os.Getenv("FIREBASE_CLIENT_EMAIL"),
		ClientID:               os.Getenv("FIREBASE_CLIENT_ID"),
		AuthURI:                os.Getenv("FIREBASE_AUTH_URI"),
		TokenURI:               os.Getenv("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertURL: os.Getenv("FIREBASE_AUTH_PROVIDER_CERT_URL"),
		ClientX509CertURL:      os.Getenv("FIREBASE_CLIENT_CERT_URL"),
	}

	// Validate required fields
	if sa.ProjectID == "" || sa.PrivateKey == "" || sa.ClientEmail == "" {
		return nil, fmt.Errorf("missing required environment variables for Firebase initialization")
	}

	// Convert to JSON
	credJSON, err := json.Marshal(sa)
	if err != nil {
		return nil, fmt.Errorf("error marshaling service account to JSON: %v", err)
	}

	opt := option.WithCredentialsJSON(credJSON)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	log.Println("Firebase initialized successfully")
	return app, nil
}
