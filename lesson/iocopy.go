package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func v1() {
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err == io.EOF {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func v1_1() {
	var fd *os.File
	var err error
	var n int
	fmt.Println(len(os.Args))
	if len(os.Args) > 1 {
		fd, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
	} else {
		fd = os.Stdin
	}

	buf := make([]byte, 1024)
	for {
		n, err = fd.Read(buf)
		if err != nil {
			return
		}
		os.Stdout.Write(buf[:n])
	}
}

func v2() {
	for {
		io.Copy(os.Stdout, os.Stdin)
	}
}

func v3() {
	var fd *os.File
	var err error
	if len(os.Args) > 1 {
		fd, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
	} else {
		fd = os.Stdin
	}
	// io.Copy会把右边所有的拷贝到左边直到文件结束，遇到EOF。
	io.Copy(os.Stdout, fd)

}

func main() {
	v3()
}
