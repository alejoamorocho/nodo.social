package firebase

import (
    "context"
    "log"

    firebase "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/db"
    "google.golang.org/api/option"
)

type FirebaseConfig struct {
    app *firebase.App
    db  *db.Client
}

func NewFirebaseConfig(ctx context.Context, credentialsFile string) (*FirebaseConfig, error) {
    opt := option.WithCredentialsFile(credentialsFile)
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("Error initializing Firebase app: %v\n", err)
        return nil, err
    }

    client, err := app.Database(ctx)
    if err != nil {
        log.Fatalf("Error initializing Firebase database: %v\n", err)
        return nil, err
    }

    return &FirebaseConfig{
        app: app,
        db:  client,
    }, nil
}

func (fc *FirebaseConfig) GetDB() *db.Client {
    return fc.db
}

func (fc *FirebaseConfig) GetApp() *firebase.App {
    return fc.app
}

