package usecase

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func validateFile(fileHeader *multipart.FileHeader) error {
	if !isAllowedImageFormat(fileHeader.Filename) {
		return errors.New("400: Image format is not allowed")
	}

	fileSize := fileHeader.Size
	if fileSize < 10240 || fileSize > 2097152 {
		return errors.New("400: Image size should be between 10KB and 2MB")
	}

	return nil
}

func isAllowedImageFormat(filename string) bool {
	allowedFormats := []string{".jpg", ".jpeg"}
	ext := filepath.Ext(filename)
	for _, format := range allowedFormats {
		if strings.EqualFold(ext, format) {
			return true
		}
	}
	return false
}

func generateUniqueFileName() string {
	uuid := uuid.New()
	path := "backend-magang/"
	return fmt.Sprintf("%s%s.jpeg", path, uuid)
}
