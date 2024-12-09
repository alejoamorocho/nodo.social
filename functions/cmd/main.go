package main

import (
    "context"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "github.com/joho/godotenv"
    "github.com/kha0sys/nodo.social/domain/repositories"
    "github.com/kha0sys/nodo.social/interfaces/http/handlers"
    "github.com/kha0sys/nodo.social/internal/config"
    "github.com/kha0sys/nodo.social/services"
)

func main() {
    // Cargar variables de entorno
    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: No .env file found: %v", err)
    }

    // Cargar configuración
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Inicializar Firebase
    app, err := config.InitializeFirebase()
    if err != nil {
        log.Fatalf("Failed to initialize Firebase: %v", err)
    }

    // Inicializar cliente de Firestore
    client, err := app.Firestore(context.Background())
    if err != nil {
        log.Fatalf("Failed to create Firestore client: %v", err)
    }
    defer client.Close()

    // Configurar el router
    router := mux.NewRouter()

    // Configurar CORS
    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
        MaxAge:         300,
    })

    // Inicializar repositorios
    nodeRepo := repositories.NewFirestoreNodeRepository(client)
    storeRepo := repositories.NewFirestoreStoreRepository(client)
    productRepo := repositories.NewFirestoreProductRepository(client)
    userRepo := repositories.NewFirestoreUserRepository(client)

    // Inicializar servicios
    nodeService := services.NewNodeService(nodeRepo)
    storeService := services.NewStoreService(storeRepo)
    productService := services.NewProductService(productRepo)
    userService := services.NewUserService(userRepo)

    // Inicializar handlers
    nodeHandler := handlers.NewNodeHandler(nodeService)
    storeHandler := handlers.NewStoreHandler(storeService)
    productHandler := handlers.NewProductHandler(productService)
    userHandler := handlers.NewUserHandler(userService)

    // Registrar rutas
    router.HandleFunc("/nodes", nodeHandler.CreateNode).Methods("POST")
    router.HandleFunc("/nodes/{id}", nodeHandler.GetNode).Methods("GET")
    router.HandleFunc("/nodes/{id}", nodeHandler.UpdateNode).Methods("PUT")
    router.HandleFunc("/nodes/{id}", nodeHandler.DeleteNode).Methods("DELETE")
    router.HandleFunc("/nodes/{id}/followers", nodeHandler.GetFollowers).Methods("GET")
    router.HandleFunc("/nodes/{id}/followers/{userId}", nodeHandler.AddFollower).Methods("POST")
    router.HandleFunc("/nodes/{id}/followers/{userId}", nodeHandler.RemoveFollower).Methods("DELETE")
    router.HandleFunc("/nodes/{id}/products", nodeHandler.GetLinkedProducts).Methods("GET")
    router.HandleFunc("/feed", nodeHandler.GetFeed).Methods("GET")

    // Registrar rutas de store
    router.HandleFunc("/stores", storeHandler.CreateStore).Methods("POST")
    router.HandleFunc("/stores/{id}", storeHandler.GetStore).Methods("GET")
    router.HandleFunc("/stores/{id}", storeHandler.UpdateStore).Methods("PUT")
    router.HandleFunc("/stores/{id}", storeHandler.DeleteStore).Methods("DELETE")

    // Registrar rutas de product
    router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
    router.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
    router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
    router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

    // Registrar rutas de user
    router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
    router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

    // Aplicar middleware CORS
    handler := corsHandler.Handler(router)

    // Iniciar el servidor
    server := &http.Server{
        Addr:         ":" + cfg.Server.Port,
        Handler:      handler,
        ReadTimeout:  cfg.Server.ReadTimeout,
        WriteTimeout: cfg.Server.WriteTimeout,
    }

    log.Printf("Server starting on port %s", cfg.Server.Port)
    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
