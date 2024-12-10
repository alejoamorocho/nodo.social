package main

import (
    "context"
    "log"
    "net/http"
    "os"

    firebase "firebase.google.com/go/v4"
    "github.com/gorilla/mux"
    "github.com/kha0sys/nodo.social/functions/interfaces/http/middleware"
    "github.com/kha0sys/nodo.social/functions/interfaces/http/handlers"
    "google.golang.org/api/option"
)

func main() {
    // Inicializar Firebase
    ctx := context.Background()
    opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("Error initializing Firebase app: %v\n", err)
    }

    // Crear router y middleware
    router := mux.NewRouter()
    authMiddleware := middleware.NewAuthMiddleware(app)

    // Configurar rutas
    apiRouter := router.PathPrefix("/api").Subrouter()
    
    // Rutas públicas
    apiRouter.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
    
    // Rutas protegidas
    protected := apiRouter.PathPrefix("").Subrouter()
    protected.Use(func(next http.Handler) http.Handler {
        return authMiddleware.Authenticate(next)
    })
    
    // Iniciar servidor
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Printf("Server starting on port %s\n", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
        log.Fatal(err)
    }
}
