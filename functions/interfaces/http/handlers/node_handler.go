package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/kha0sys/nodo.social/domain/dto"
    "github.com/kha0sys/nodo.social/domain/models"
    "github.com/kha0sys/nodo.social/services"
    "github.com/kha0sys/nodo.social/interfaces/http/utils"
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

// RegisterRoutes registra las rutas del handler en el router
func (h *NodeHandler) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/nodes", h.CreateNode).Methods("POST")
    r.HandleFunc("/nodes/{id}", h.GetNode).Methods("GET")
    r.HandleFunc("/nodes/{id}", h.UpdateNode).Methods("PUT")
    r.HandleFunc("/nodes/{id}", h.DeleteNode).Methods("DELETE")
    r.HandleFunc("/nodes/{id}/followers", h.AddFollower).Methods("POST")
    r.HandleFunc("/nodes/{id}/followers/{userID}", h.RemoveFollower).Methods("DELETE")
    r.HandleFunc("/nodes/{id}/followers", h.GetFollowers).Methods("GET")
    r.HandleFunc("/nodes/{id}/products", h.GetLinkedProducts).Methods("GET")
    r.HandleFunc("/nodes/feed", h.GetFeed).Methods("GET")
}

// CreateNode maneja la creación de un nuevo nodo
func (h *NodeHandler) CreateNode(w http.ResponseWriter, r *http.Request) {
    var nodeDTO dto.NodeDTO
    if err := json.NewDecoder(r.Body).Decode(&nodeDTO); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, utils.ErrInvalidInput, "Invalid request body")
        return
    }

    createdNode, err := h.nodeService.CreateNode(r.Context(), nodeDTO)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusCreated, createdNode)
}

// GetNode maneja la obtención de un nodo por ID
func (h *NodeHandler) GetNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    node, err := h.nodeService.GetNode(r.Context(), id)
    if err != nil {
        utils.RespondWithError(w, http.StatusNotFound, utils.ErrNotFound, "Node not found")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, node)
}

// UpdateNode maneja la actualización de un nodo
func (h *NodeHandler) UpdateNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var nodeDTO dto.NodeDTO
    if err := json.NewDecoder(r.Body).Decode(&nodeDTO); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, utils.ErrInvalidInput, "Invalid request body")
        return
    }

    node := nodeDTO.ToModel()
    node.ID = id

    if err := h.nodeService.UpdateNode(r.Context(), node); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, node)
}

// DeleteNode maneja la eliminación de un nodo
func (h *NodeHandler) DeleteNode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    if err := h.nodeService.DeleteNode(r.Context(), id); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Node deleted successfully"})
}

// AddFollower maneja la adición de un seguidor a un nodo
func (h *NodeHandler) AddFollower(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]
    var followerData struct {
        UserID string `json:"userID"`
    }

    if err := json.NewDecoder(r.Body).Decode(&followerData); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, utils.ErrInvalidInput, "Invalid request body")
        return
    }

    if err := h.nodeService.AddFollower(r.Context(), nodeID, followerData.UserID); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Follower added successfully"})
}

// RemoveFollower maneja la eliminación de un seguidor de un nodo
func (h *NodeHandler) RemoveFollower(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]
    userID := vars["userID"]

    if err := h.nodeService.RemoveFollower(r.Context(), nodeID, userID); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Follower removed successfully"})
}

// GetFollowers maneja la obtención de seguidores de un nodo
func (h *NodeHandler) GetFollowers(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]

    followers, err := h.nodeService.GetFollowers(r.Context(), nodeID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, followers)
}

// GetLinkedProducts maneja la obtención de productos vinculados a un nodo
func (h *NodeHandler) GetLinkedProducts(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    nodeID := vars["id"]

    products, err := h.nodeService.GetLinkedProducts(r.Context(), nodeID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, products)
}

// GetFeed maneja la obtención del feed de nodos
func (h *NodeHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

    if page < 1 {
        page = 1
    }
    if limit < 1 || limit > 100 {
        limit = 10
    }

    filters := models.FeedFilters{
        Page:  page,
        Limit: limit,
    }

    nodes, err := h.nodeService.GetFeed(r.Context(), filters)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, nodes)
}
