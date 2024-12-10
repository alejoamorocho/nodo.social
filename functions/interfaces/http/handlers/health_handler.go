package handlers

import (
    "encoding/json"
    "net/http"
)

// HealthCheck maneja la ruta de health check
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{
        "status": "ok",
        "message": "Service is healthy",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
