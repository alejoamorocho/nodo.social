package cloud

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
    "github.com/kha0sys/nodo.social/functions/domain/dto"
    "github.com/kha0sys/nodo.social/functions/services"
)

// NodeFunctions maneja las Cloud Functions relacionadas con nodos
type NodeFunctions struct {
    nodeService *services.NodeService
}

// NewNodeFunctions crea una nueva instancia de NodeFunctions
func NewNodeFunctions(nodeService *services.NodeService) *NodeFunctions {
    return &NodeFunctions{
        nodeService: nodeService,
    }
}

// CreateNode es una Cloud Function para crear nodos
func (f *NodeFunctions) CreateNode(w http.ResponseWriter, r *http.Request) {
    var nodeDTO dto.NodeDTO
    if err := json.NewDecoder(r.Body).Decode(&nodeDTO); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Set timestamps
    nodeDTO.CreatedAt = time.Now()
    nodeDTO.UpdatedAt = time.Now()

    node, err := f.nodeService.CreateNode(r.Context(), nodeDTO)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := dto.FromNodeModel(node)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// GetNodeFeed es una Cloud Function para obtener el feed de nodos
func (f *NodeFunctions) GetNodeFeed(w http.ResponseWriter, r *http.Request) {
    nodes, err := f.nodeService.GetNodeFeed(r.Context())
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

// FollowNode es una Cloud Function para seguir a un nodo
func (f *NodeFunctions) FollowNode(w http.ResponseWriter, r *http.Request) {
    // Get nodeId and userId from request
    var req struct {
        NodeID string `json:"nodeId"`
        UserID string `json:"userId"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := f.nodeService.FollowNode(r.Context(), req.NodeID, req.UserID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// RegisterFunctions registra todas las Cloud Functions
func (f *NodeFunctions) RegisterFunctions() {
    functions.HTTP("CreateNode", f.CreateNode)
    functions.HTTP("GetNodeFeed", f.GetNodeFeed)
    functions.HTTP("FollowNode", f.FollowNode)
}

