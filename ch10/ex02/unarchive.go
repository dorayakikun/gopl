package unarchive

import (
	"compressor/zip"
	"errors"
	"os"
)

type File interface {
	FileInfo() os.FileInfo
	Open()(*os.File, error)
}

type ReadCloser interface {
	Next()(File , error)
}

func OpenReader(name string) (ReadCloser, error) {
	uz := &zip.UnarchiveZip{}

	return uz, errors.New("unimplmented")
}
