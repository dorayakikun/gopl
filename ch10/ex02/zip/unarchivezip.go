package zip

import (
	unarchive "compressor"
	"os"
)

type ZipFile struct {

}

func (z *ZipFile) FileInfo() os.FileInfo {
	return nil
}

func (z *ZipFile) Open()(*os.File, error) {
	return nil, nil
}

type UnarchiveZip struct {
	files []*ZipFile
}

func (uz *UnarchiveZip) Next() (unarchive.File, error) {
	return nil, nil
}
