package services

import (
	"context"
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/kha0sys/nodo.social/functions/domain/repositories"
)

// StorageService maneja las operaciones de almacenamiento de archivos
type StorageService struct {
	storage repositories.StorageRepository
}

// NewStorageService crea una nueva instancia de StorageService
func NewStorageService(storage repositories.StorageRepository) *StorageService {
	return &StorageService{
		storage: storage,
	}
}

// UploadFile sube un archivo al storage y retorna su URL
func (s *StorageService) UploadFile(ctx context.Context, userID string, content io.Reader, filename string, contentType string) (string, error) {
	// Sanitiza el nombre del archivo
	sanitizedFilename := sanitizeFilename(filename)
	
	// Construye la ruta del archivo: users/{userID}/files/{filename}
	storagePath := path.Join("users", userID, "files", sanitizedFilename)
	
	return s.storage.Upload(ctx, storagePath, content, contentType)
}

// UploadImage sube una imagen al storage y retorna su URL
func (s *StorageService) UploadImage(ctx context.Context, userID string, content io.Reader, filename string) (string, error) {
	// Verifica que el archivo sea una imagen
	if !isImageFile(filename) {
		return "", fmt.Errorf("el archivo debe ser una imagen (jpg, png, gif)")
	}
	
	// Construye la ruta del archivo: users/{userID}/images/{filename}
	storagePath := path.Join("users", userID, "images", sanitizeFilename(filename))
	
	contentType := getImageContentType(filename)
	return s.storage.Upload(ctx, storagePath, content, contentType)
}

// DeleteFile elimina un archivo del storage
func (s *StorageService) DeleteFile(ctx context.Context, path string) error {
	return s.storage.Delete(ctx, path)
}

// GetFileURL obtiene la URL firmada de un archivo
func (s *StorageService) GetFileURL(ctx context.Context, path string) (string, error) {
	return s.storage.GetURL(ctx, path)
}

// Funciones auxiliares

func sanitizeFilename(filename string) string {
	// Elimina caracteres especiales y espacios
	sanitized := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.' {
			return r
		}
		return '_'
	}, filename)
	
	return sanitized
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}

func getImageContentType(filename string) string {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}
