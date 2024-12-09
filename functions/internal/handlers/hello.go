package handlers

import (
	"encoding/json"
	"net/http"
)

// HelloWorld is an HTTP Cloud Function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "¡Hola desde Nodo Social!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
