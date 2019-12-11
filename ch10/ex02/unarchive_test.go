package unarchive

import (
	"io"
	"log"
	"os"
)

func ExampleOpenReader() {
	unarchive("sample.tar")
	// Unordered output:
	// 吾輩は猫である。名前はまだ無い。
	// どこで生れたかとんと見当がつかぬ。
	// 何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。
}

func ExampleOpenReader2() {
	unarchive("sample.zip")
	// Unordered output:
	// 吾輩は猫である。名前はまだ無い。
	// どこで生れたかとんと見当がつかぬ。
	// 何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。
}

func unarchive(name string) {
	r, err := OpenReader(name)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		f, err := r.Next()
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			}
			break
		}
		if !f.FileInfo().IsDir() {
			fl, err := f.Open()
			if err != nil {
				log.Fatalln(err)
			}
			io.Copy(os.Stdout, fl)
		}
	}
}
