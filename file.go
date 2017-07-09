/*
	读取文件的方式
	1. 按行读取
	2. 按块读取
	3. 全部读取
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func v1() {
	var f *os.File
	f, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Printf("open file error:%v", err)
	}
	defer f.Close()
}

func v2() {
	f, err := os.Open("defer.go")
	if err != nil {
		return
	}
	read(f)
}

func read(f *os.File) (string, error) {
	var total []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			return "", err
		}
		if err == io.EOF {
			break
		}
		total = append(total, buf[:n]...)
	}
	return string(total), nil
}

func main() {
	v1()
	v2()
}
