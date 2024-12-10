package router

import (
    "net/http"
    firebase "firebase.google.com/go/v4"
    "github.com/gorilla/mux"
    "github.com/kha0sys/nodo.social/functions/interfaces/http/handlers"
    "github.com/kha0sys/nodo.social/functions/interfaces/http/middleware"
)

// Router maneja la configuración de rutas de la aplicación
type Router struct {
    router *mux.Router
    auth   *middleware.AuthMiddleware
    app    *firebase.App
}

// NewRouter crea una nueva instancia del router
func NewRouter(app *firebase.App) *Router {
    return &Router{
        router: mux.NewRouter(),
        auth:   middleware.NewAuthMiddleware(app),
        app:    app,
    }
}

// SetupRoutes configura todas las rutas de la aplicación
func (r *Router) SetupRoutes() http.Handler {
    // Crear handlers
    nodeHandler := handlers.NewNodeHandler(r.app)

    // API Router
    api := r.router.PathPrefix("/api").Subrouter()

    // Rutas públicas
    api.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

    // Rutas de nodos (protegidas)
    nodes := api.PathPrefix("/nodes").Subrouter()
    nodes.Use(func(next http.Handler) http.Handler {
        return r.auth.Authenticate(next)
    })

    // Registrar rutas de nodos
    nodeHandler.RegisterRoutes(nodes)

    return r.router
}

// GetRouter retorna el router subyacente
func (r *Router) GetRouter() *mux.Router {
    return r.router
}
