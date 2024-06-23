package fileupload

import (
	"errors"
	"io"
	"path/filepath"
	"strings"
)

type FileType string

const (
	FileTypeVideo    FileType = "video"
	FileTypeImage    FileType = "image"
	FileTypeDocument FileType = "document"
)

type UploadValidator interface {
	Validate(filename string, reader io.Reader) error
}

type ExtensionValidator struct {
	allowedExtensions map[FileType][]string
}

func NewExtensionValidator() *ExtensionValidator {
	return &ExtensionValidator{
		allowedExtensions: map[FileType][]string{
			FileTypeVideo:    {".mp4", ".mpeg", ".avi", ".mov"},
			FileTypeImage:    {".jpg", ".jpeg", ".png", ".gif"},
			FileTypeDocument: {".xls", ".xlsx", ".csv", ".pdf"},
		},
	}
}

func (v *ExtensionValidator) Validate(filename string, _ io.Reader) error {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, allowed := range v.allowedExtensions {
		if contains(allowed, ext) {
			return nil
		}
	}
	return errors.New("invalid file extension")
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
