package firebase

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/kha0sys/nodo.social/functions/domain/repositories"
)

// FirebaseStorageRepository implementa StorageRepository usando Firebase Storage
type FirebaseStorageRepository struct {
	bucket  *storage.BucketHandle
	app     *firebase.App
	baseURL string
}

// NewFirebaseStorageRepository crea una nueva instancia de FirebaseStorageRepository
func NewFirebaseStorageRepository(ctx context.Context, app *firebase.App, bucketName string) (repositories.StorageRepository, error) {
	client, err := app.Storage(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting storage client: %v", err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, fmt.Errorf("error getting bucket: %v", err)
	}

	baseURL := fmt.Sprintf("https://storage.googleapis.com/%s", bucketName)

	return &FirebaseStorageRepository{
		bucket:  bucket,
		app:     app,
		baseURL: baseURL,
	}, nil
}

// Upload implementa StorageRepository.Upload
func (r *FirebaseStorageRepository) Upload(ctx context.Context, path string, content io.Reader, contentType string) (string, error) {
	obj := r.bucket.Object(path)
	writer := obj.NewWriter(ctx)
	writer.ContentType = contentType

	if _, err := io.Copy(writer, content); err != nil {
		return "", fmt.Errorf("error copying content: %v", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("error closing writer: %v", err)
	}

	return fmt.Sprintf("%s/%s", r.baseURL, path), nil
}

// Download implementa StorageRepository.Download
func (r *FirebaseStorageRepository) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	obj := r.bucket.Object(path)
	return obj.NewReader(ctx)
}

// Delete implementa StorageRepository.Delete
func (r *FirebaseStorageRepository) Delete(ctx context.Context, path string) error {
	obj := r.bucket.Object(path)
	return obj.Delete(ctx)
}

// GetURL implementa StorageRepository.GetURL
func (r *FirebaseStorageRepository) GetURL(ctx context.Context, path string) (string, error) {
	// Genera una URL firmada que expira en 1 hora
	opts := &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(1 * time.Hour),
	}
	
	return r.bucket.SignedURL(path, opts)
}
