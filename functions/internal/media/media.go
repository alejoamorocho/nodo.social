// Package media proporciona tipos y utilidades compartidas para el manejo de recursos multimedia
package media

import "fmt"

// MediaURL representa una URL de un recurso multimedia junto con su tipo y miniatura.
// Se utiliza para almacenar imágenes, videos y otros recursos multimedia asociados a nodos,
// productos y actualizaciones.
type MediaURL struct {
    // URL es la dirección del recurso multimedia
    URL       string `json:"url" firestore:"url"`
    // Type indica el tipo de medio (image, video, document, etc.)
    Type      string `json:"type" firestore:"type"`
    // ThumbnailURL es la URL de la miniatura del recurso (opcional)
    ThumbnailURL string `json:"thumbnailUrl,omitempty" firestore:"thumbnailUrl,omitempty"`
}

// Validate verifica que los campos obligatorios del recurso multimedia estén presentes y sean válidos.
// Retorna un error si la URL está vacía o si el tipo no es válido.
func (m *MediaURL) Validate() error {
    if m.URL == "" {
        return fmt.Errorf("la URL del recurso multimedia es obligatoria")
    }
    if m.Type == "" {
        return fmt.Errorf("el tipo de recurso multimedia es obligatorio")
    }
    // Validar que el tipo sea uno de los permitidos
    validTypes := map[string]bool{
        "image":    true,
        "video":    true,
        "document": true,
        "audio":    true,
    }
    if !validTypes[m.Type] {
        return fmt.Errorf("tipo de recurso multimedia inválido: %s", m.Type)
    }
    return nil
}

