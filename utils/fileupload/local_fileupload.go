package fileupload

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	EMPTY_STRING = ""
)

type LocalFileUploader struct {
	validators []UploadValidator
}

func NewLocalFileUploader(
	extValidator *ExtensionValidator,
) *LocalFileUploader {
	validators := []UploadValidator{extValidator}
	return &LocalFileUploader{validators}
}

func (u *LocalFileUploader) Save(reader io.Reader, uploadDir string, filename string) (string, error) {
	for _, validator := range u.validators {
		if err := validator.Validate(filename, reader); err != nil {
			return EMPTY_STRING, err
		}
	}

	if err := createDirectory(uploadDir); err != nil {
		return EMPTY_STRING, fmt.Errorf("could not create upload directory: %v", err)
	}

	filePath := filepath.Join(uploadDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return EMPTY_STRING, fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil {
		return EMPTY_STRING, fmt.Errorf("could not copy data to file: %v", err)
	}

	return filename, nil
}

func createDirectory(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}
