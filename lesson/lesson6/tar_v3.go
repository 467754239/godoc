package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	uncompress, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	tr := tar.NewReader(uncompress)
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
