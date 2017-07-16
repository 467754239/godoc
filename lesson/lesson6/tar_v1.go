package main

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	tr := tar.NewReader(os.Stdin)
	for {
		hd, err := tr.Next()
		if err == io.EOF {
			return
		}

		if err != nil {
			return
		}
		fmt.Println(hd.Name)
		io.Copy(ioutil.Discard, tr) // > /dev/null
	}

}
