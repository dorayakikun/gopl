package unarchive

import (
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
	return nil, errors.New("unimplemented")
}
