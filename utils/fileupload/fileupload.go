package fileupload

import (
	"io"
)

type FileUploader interface {
	Save(reader io.Reader, filename string) (string, error)
}
