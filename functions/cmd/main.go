package main

import (
	"log"
	"os"

	"github.com/kha0sys/nodo.social/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Inicializar Firebase
	_, err := config.InitializeFirebase()
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
}
