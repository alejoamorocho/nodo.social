package repositories

import (
	"context"
	"io"
)

// StorageRepository define la interfaz para operaciones con almacenamiento de archivos
type StorageRepository interface {
	// Upload sube un archivo al storage y retorna su URL pública
	Upload(ctx context.Context, path string, content io.Reader, contentType string) (string, error)

	// Download descarga un archivo del storage
	Download(ctx context.Context, path string) (io.ReadCloser, error)

	// Delete elimina un archivo del storage
	Delete(ctx context.Context, path string) error

	// GetURL obtiene la URL pública de un archivo
	GetURL(ctx context.Context, path string) (string, error)
}
