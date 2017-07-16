package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var base string = "/tmp/d/"
	tr := tar.NewReader(os.Stdin)
	for {
		hd, err := tr.Next()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		info := hd.FileInfo()
		if info.IsDir() {
			dirpath := filepath.Join(base, filepath.Dir(hd.Name))
			err = os.MkdirAll(dirpath, 0755)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fpath := filepath.Join(base, hd.Name)
			f, err := os.Create(fpath)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
		}

	}

}
