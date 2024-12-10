package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    
    "cloud.google.com/go/firestore"
    firebase "firebase.google.com/go/v4"
    "github.com/gorilla/mux"
    "google.golang.org/api/iterator"
    "github.com/kha0sys/nodo.social/functions/domain/models"
    "github.com/kha0sys/nodo.social/functions/interfaces/http/middleware"
)

// NodeHandler maneja las peticiones HTTP relacionadas con nodos
type NodeHandler struct {
    app *firebase.App
}

// NewNodeHandler crea una nueva instancia de NodeHandler
func NewNodeHandler(app *firebase.App) *NodeHandler {
    return &NodeHandler{
        app: app,
    }
}

// RegisterRoutes registra las rutas del handler en el router
func (h *NodeHandler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/nodes", h.CreateNode).Methods("POST")
    r.HandleFunc("/nodes/{id}", h.GetNode).Methods("GET")
    r.HandleFunc("/nodes/{id}", h.UpdateNode).Methods("PUT")
    r.HandleFunc("/nodes/{id}", h.DeleteNode).Methods("DELETE")
    r.HandleFunc("/nodes/{id}/followers", h.GetFollowers).Methods("GET")
    r.HandleFunc("/nodes/feed", h.GetFeed).Methods("GET")
}

// CreateNode maneja la creación de un nuevo nodo
func (h *NodeHandler) CreateNode(w http.ResponseWriter, r *http.Request) {
    var node models.Node
    if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Obtener información del usuario del contexto
    userID, _, _ := middleware.GetUserFromContext(r.Context())
    if userID == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Añadir metadatos
    node.UserID = userID
    node.CreatedAt = time.Now()
    node.UpdatedAt = time.Now()

    // Crear el documento en Firestore
    client, err := h.app.Firestore(r.Context())
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer client.Close()

    docRef, _, err := client.Collection("nodes").Add(r.Context(), node)
    if err != nil {
        http.Error(w, "Error creating node", http.StatusInternalServerError)
        return
    }

    node.ID = docRef.ID
    json.NewEncoder(w).Encode(node)
}

// GetNode maneja la obtención de un nodo por ID
func (h *NodeHandler) GetNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]

    client, err := h.app.Firestore(r.Context())
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer client.Close()

    doc, err := client.Collection("nodes").Doc(nodeID).Get(r.Context())
    if err != nil {
        http.Error(w, "Node not found", http.StatusNotFound)
        return
    }

    var node models.Node
    if err := doc.DataTo(&node); err != nil {
        http.Error(w, "Error parsing node data", http.StatusInternalServerError)
        return
    }

    node.ID = doc.Ref.ID
    json.NewEncoder(w).Encode(node)
}

// UpdateNode maneja la actualización de un nodo
func (h *NodeHandler) UpdateNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]

    var updates models.Node
    if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Obtener información del usuario del contexto
    userID, _, userRole := middleware.GetUserFromContext(r.Context())
    if userID == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    client, err := h.app.Firestore(r.Context())
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer client.Close()

    // Verificar propiedad del nodo
    doc, err := client.Collection("nodes").Doc(nodeID).Get(r.Context())
    if err != nil {
        http.Error(w, "Node not found", http.StatusNotFound)
        return
    }

    var existingNode models.Node
    if err := doc.DataTo(&existingNode); err != nil {
        http.Error(w, "Error parsing node data", http.StatusInternalServerError)
        return
    }

    // Solo el creador o un admin pueden actualizar el nodo
    if existingNode.UserID != userID && userRole != "admin" {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    updates.UpdatedAt = time.Now()
    updates.UserID = existingNode.UserID // No permitir cambiar el creador
    updates.CreatedAt = existingNode.CreatedAt // No permitir cambiar la fecha de creación

    _, err = client.Collection("nodes").Doc(nodeID).Set(r.Context(), updates, firestore.MergeAll)
    if err != nil {
        http.Error(w, "Error updating node", http.StatusInternalServerError)
        return
    }

    updates.ID = nodeID
    json.NewEncoder(w).Encode(updates)
}

// DeleteNode maneja la eliminación de un nodo
func (h *NodeHandler) DeleteNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]

    // Obtener información del usuario del contexto
    userID, _, userRole := middleware.GetUserFromContext(r.Context())
    if userID == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    client, err := h.app.Firestore(r.Context())
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer client.Close()

    // Verificar propiedad del nodo
    doc, err := client.Collection("nodes").Doc(nodeID).Get(r.Context())
    if err != nil {
        http.Error(w, "Node not found", http.StatusNotFound)
        return
    }

    var node models.Node
    if err := doc.DataTo(&node); err != nil {
        http.Error(w, "Error parsing node data", http.StatusInternalServerError)
        return
    }

    // Solo el creador o un admin pueden eliminar el nodo
    if node.UserID != userID && userRole != "admin" {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    _, err = client.Collection("nodes").Doc(nodeID).Delete(r.Context())
    if err != nil {
        http.Error(w, "Error deleting node", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// GetFollowers maneja la obtención de seguidores de un nodo
func (h *NodeHandler) GetFollowers(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]

    client, err := h.app.Firestore(r.Context())
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer client.Close()

    followers := make([]string, 0)
    iter := client.Collection("nodes").Doc(nodeID).Collection("followers").Documents(r.Context())
    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            http.Error(w, "Error getting followers", http.StatusInternalServerError)
            return
        }
        followers = append(followers, doc.Ref.ID)
    }

    json.NewEncoder(w).Encode(followers)
}

// GetFeed maneja la obtención del feed de nodos
func (h *NodeHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
    limit := 10 // Número de nodos por página
    
    client, err := h.app.Firestore(r.Context())
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer client.Close()

    // Obtener los nodos más recientes
    query := client.Collection("nodes").
        OrderBy("CreatedAt", firestore.Desc).
        Limit(limit)

    docs, err := query.Documents(r.Context()).GetAll()
    if err != nil {
        http.Error(w, "Error getting feed", http.StatusInternalServerError)
        return
    }

    nodes := make([]models.Node, 0, len(docs))
    for _, doc := range docs {
        var node models.Node
        if err := doc.DataTo(&node); err != nil {
            continue // Saltar documentos con error
        }
        node.ID = doc.Ref.ID
        nodes = append(nodes, node)
    }

    json.NewEncoder(w).Encode(nodes)
}
