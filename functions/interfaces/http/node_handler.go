package http

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "github.com/kha0sys/nodo.social/domain/dto"
    "github.com/kha0sys/nodo.social/services"
)

// NodeHandler maneja las peticiones HTTP relacionadas con nodos
type NodeHandler struct {
    nodeService *services.NodeService
}

// NewNodeHandler crea una nueva instancia de NodeHandler
func NewNodeHandler(nodeService *services.NodeService) *NodeHandler {
    return &NodeHandler{
        nodeService: nodeService,
    }
}

// CreateNode maneja POST /nodes
func (h *NodeHandler) CreateNode(w http.ResponseWriter, r *http.Request) {
    var nodeDTO dto.NodeDTO
    if err := json.NewDecoder(r.Body).Decode(&nodeDTO); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Set timestamps
    nodeDTO.CreatedAt = time.Now()
    nodeDTO.UpdatedAt = time.Now()

    node, err := h.nodeService.CreateNode(r.Context(), nodeDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := dto.FromNodeModel(node)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// GetNode maneja GET /nodes/{id}
func (h *NodeHandler) GetNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]

    node, err := h.nodeService.GetNode(r.Context(), nodeID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := dto.FromNodeModel(node)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// GetNodeFeed maneja GET /nodes/feed
func (h *NodeHandler) GetNodeFeed(w http.ResponseWriter, r *http.Request) {
    nodes, err := h.nodeService.GetNodeFeed(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var response []dto.NodeDTO
    for _, node := range nodes {
        nodeDTO := dto.FromNodeModel(node)
        response = append(response, *nodeDTO)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// FollowNode maneja POST /nodes/{id}/follow
func (h *NodeHandler) FollowNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]
    
    // TODO: Obtener userID del contexto de autenticación
    userID := r.Header.Get("X-User-ID")

    if err := h.nodeService.FollowNode(r.Context(), nodeID, userID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// RegisterRoutes registra todas las rutas relacionadas con nodos
func (h *NodeHandler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/nodes", h.CreateNode).Methods("POST")
    r.HandleFunc("/nodes/{id}", h.GetNode).Methods("GET")
    r.HandleFunc("/nodes/feed", h.GetNodeFeed).Methods("GET")
    r.HandleFunc("/nodes/{id}/follow", h.FollowNode).Methods("POST")
}
