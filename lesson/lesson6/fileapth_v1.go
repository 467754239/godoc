/*
	遍历文件树
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/usr", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
}
