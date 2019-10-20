package main

import (
	"archive/tar"
	_ "archive/tar"
	"archive/zip"
	_ "archive/zip"
	"bytes"
	"flag"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var format string
	flag.StringVar(&format, "format", "zip", "compression format")
	flag.Parse()

	if len(flag.Args()) < 2 {
		log.Fatalf("invalid argument length want 2 more actual: %d\n", len(flag.Args()))
	}

	filename := flag.Args()[0]
	files := flag.Args()[1:]
	switch format {
	case "zip":
		b, err := compressZip(files)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = f.Write(b.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	case "tar":
		b, err := compressTar(files)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = f.Write(b.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unsuported format : %s\n", format)
	}
}

func compressZip(files []string) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)
	w := zip.NewWriter(b)
	defer w.Close()

	for _, file := range files {
		info, _ := os.Stat(file)

		fh, err := zip.FileInfoHeader(info)
		if err != nil {
			return nil, errors.Wrap(err, "failed get header")
		}
		fh.Name = "files/" + file

		f, err := w.CreateHeader(fh)
		if err != nil {
			return nil, errors.Wrap(err, "failed create header")
		}

		body, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, errors.Wrap(err, "failed read file")
		}
		f.Write(body)
	}

	return b, nil
}

func compressTar(files []string) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)
	w := tar.NewWriter(b)
	defer w.Close()

	for _, file := range files {
		info, _ := os.Stat(file)
		hdr := &tar.Header{

			Name: file,

			Mode: 0600,

			Size: info.Size(),
		}
		if err := w.WriteHeader(hdr); err != nil {
			return nil, errors.Wrap(err, "failed write header")
		}

		body, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, errors.Wrap(err, "failed read file")
		}
		_, err = w.Write(body)
		if err != nil {
			return nil, errors.Wrap(err, "failed write file")
		}
	}

	return b, nil
}
