package imageprocessor

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/disintegration/imaging"
)

type ImageProcessor struct {
	sizes map[string]int
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{
		sizes: map[string]int{
			"small":  150,
			"medium": 300,
			"large":  600,
		},
	}
}

// ProcessImage genera thumbnails de diferentes tamaños para una imagen
func (p *ImageProcessor) ProcessImage(reader io.Reader) (map[string][]byte, error) {
	// Decodificar imagen original
	img, format, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("error decoding image: %v", err)
	}

	thumbnails := make(map[string][]byte)

	// Generar thumbnails para cada tamaño
	for size, width := range p.sizes {
		// Redimensionar imagen manteniendo aspecto
		thumb := imaging.Resize(img, width, 0, imaging.Lanczos)

		// Codificar thumbnail
		var buf bytes.Buffer
		switch format {
		case "jpeg":
			err = jpeg.Encode(&buf, thumb, &jpeg.Options{Quality: 85})
		case "png":
			err = png.Encode(&buf, thumb)
		default:
			err = fmt.Errorf("unsupported image format: %s", format)
		}

		if err != nil {
			return nil, fmt.Errorf("error encoding thumbnail: %v", err)
		}

		thumbnails[size] = buf.Bytes()
	}

	return thumbnails, nil
}
