package triggers

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/kha0sys/nodo.social/functions/services"
	"github.com/kha0sys/nodo.social/functions/internal/firebase"
	"github.com/kha0sys/nodo.social/functions/internal/imageprocessor"
)

type StorageTriggers struct {
	storageClient   *storage.Client
	imageProcessor  *imageprocessor.ImageProcessor
	nodeService     *services.NodeService
	productService  *services.ProductService
}

func NewStorageTriggers(
	storageClient *storage.Client,
	imageProcessor *imageprocessor.ImageProcessor,
	nodeService *services.NodeService,
	productService *services.ProductService,
) *StorageTriggers {
	return &StorageTriggers{
		storageClient:  storageClient,
		imageProcessor: imageProcessor,
		nodeService:    nodeService,
		productService: productService,
	}
}

// OnImageUploaded se ejecuta cuando se sube una imagen al Storage
func (t *StorageTriggers) OnImageUploaded(ctx context.Context, e firebase.StorageEvent) error {
	// Solo procesar imágenes
	if !isImage(e.Name) {
		return nil
	}

	// Generar thumbnails y optimizar imagen
	bucket := t.storageClient.Bucket(e.Bucket)
	obj := bucket.Object(e.Name)

	// Leer la imagen original
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("error reading image: %v", err)
	}
	defer reader.Close()

	// Procesar imagen y generar thumbnails
	thumbnails, err := t.imageProcessor.ProcessImage(reader)
	if err != nil {
		return fmt.Errorf("error processing image: %v", err)
	}

	// Guardar thumbnails
	for size, thumbnail := range thumbnails {
		thumbPath := generateThumbnailPath(e.Name, size)
		thumbObj := bucket.Object(thumbPath)
		writer := thumbObj.NewWriter(ctx)
		
		if _, err := writer.Write(thumbnail); err != nil {
			return fmt.Errorf("error writing thumbnail: %v", err)
		}
		
		if err := writer.Close(); err != nil {
			return fmt.Errorf("error closing writer: %v", err)
		}
	}

	// Actualizar metadata del objeto original
	objectAttrs := storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{
			"thumbnailsGenerated": "true",
			"processedAt":        "timestamp",
		},
	}
	
	if _, err := obj.Update(ctx, objectAttrs); err != nil {
		return fmt.Errorf("error updating metadata: %v", err)
	}

	return nil
}

// OnImageDeleted se ejecuta cuando se elimina una imagen del Storage
func (t *StorageTriggers) OnImageDeleted(ctx context.Context, e firebase.StorageEvent) error {
	if !isImage(e.Name) {
		return nil
	}

	// Eliminar thumbnails asociados
	bucket := t.storageClient.Bucket(e.Bucket)
	sizes := []string{"small", "medium", "large"}

	for _, size := range sizes {
		thumbPath := generateThumbnailPath(e.Name, size)
		if err := bucket.Object(thumbPath).Delete(ctx); err != nil {
			// Ignorar errores si el thumbnail no existe
			if err != storage.ErrObjectNotExist {
				return fmt.Errorf("error deleting thumbnail: %v", err)
			}
		}
	}

	return nil
}

func isImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp"
}

func generateThumbnailPath(originalPath, size string) string {
	dir := filepath.Dir(originalPath)
	filename := filepath.Base(originalPath)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	
	return filepath.Join(dir, "thumbnails", fmt.Sprintf("%s_%s%s", name, size, ext))
}
