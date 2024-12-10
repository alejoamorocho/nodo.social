// Package cloud provides Firebase and cloud function implementations
package cloud

import (
	"context"
	"fmt"
	"log"
	"sync"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	firebaseApp *firebase.App
	initOnce    sync.Once
)

// InitFirebase inicializa la aplicación de Firebase de forma segura usando singleton
func InitFirebase(ctx context.Context) (*firebase.App, error) {
	var initErr error

	initOnce.Do(func() {
		// En producción, las credenciales se obtienen automáticamente del entorno
		// En desarrollo, necesitamos especificar el archivo de credenciales
		opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
		config := &firebase.Config{ProjectID: "nodo-social"}
		
		firebaseApp, initErr = firebase.NewApp(ctx, config, opt)
		if initErr != nil {
			log.Printf("Error initializing Firebase app: %v\n", initErr)
			return
		}
	})

	if initErr != nil {
		return nil, fmt.Errorf("error initializing firebase: %v", initErr)
	}

	return firebaseApp, nil
}
