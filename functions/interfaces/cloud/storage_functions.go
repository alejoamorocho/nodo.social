package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/kha0sys/nodo.social/functions/services"
)

var (
	storageService *services.StorageService
)

func init() {
	functions.HTTP("UploadFile", handleFileUpload)
	functions.HTTP("GetSignedURL", handleGetSignedURL)
}

type uploadResponse struct {
	URL string `json:"url"`
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	// Verifica la autenticación
	userID, err := authenticateRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Procesa el archivo
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Determina si es una imagen
	isImage := r.FormValue("type") == "image"

	var url string
	if isImage {
		url, err = storageService.UploadImage(ctx, userID, file, header.Filename)
	} else {
		url, err = storageService.UploadFile(ctx, userID, file, header.Filename, header.Header.Get("Content-Type"))
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading file: %v", err), http.StatusInternalServerError)
		return
	}

	// Retorna la URL del archivo
	response := uploadResponse{URL: url}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleGetSignedURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	// Verifica la autenticación
	userID, err := authenticateRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Obtiene el path del archivo
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		http.Error(w, "Path parameter is required", http.StatusBadRequest)
		return
	}

	// Verifica que el path pertenezca al usuario
	if !isUserFile(userID, filePath) {
		http.Error(w, "Unauthorized access to file", http.StatusForbidden)
		return
	}

	// Obtiene la URL firmada
	url, err := storageService.GetFileURL(ctx, filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting signed URL: %v", err), http.StatusInternalServerError)
		return
	}

	// Retorna la URL firmada
	response := uploadResponse{URL: url}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Función auxiliar para verificar que el archivo pertenece al usuario
func isUserFile(userID string, filePath string) bool {
	// El path debe comenzar con "users/{userID}/"
	expectedPrefix := path.Join("users", userID)
	return path.Dir(filePath) == expectedPrefix
}

// authenticateRequest verifica el token de autenticación y retorna el ID del usuario
func authenticateRequest(r *http.Request) (string, error) {
	// Obtener el token del header de autorización
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header")
	}

	// El header debe ser "Bearer <token>"
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return "", fmt.Errorf("invalid authorization header format")
	}

	token := authHeader[7:]

	// Verificar el token con Firebase Auth
	ctx := context.Background()
	app, err := InitFirebase(ctx)
	if err != nil {
		return "", fmt.Errorf("error initializing firebase: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting auth client: %v", err)
	}

	// Verificar y decodificar el token
	decodedToken, err := authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return "", fmt.Errorf("error verifying ID token: %v", err)
	}

	return decodedToken.UID, nil
}
