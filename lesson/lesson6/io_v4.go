package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type ByteCounter struct {
	Sum int
}

type LineCounter struct {
	Line int
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	b.Sum += len(p)
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	//l.Line += len(strings.Split(string(p), "\n"))
	l.Line += strings.Count(string(p), "\n")
	return len(p), nil
}

func main() {
	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatal(err)
	}

	b := new(ByteCounter)
	io.Copy(b, &buf1)
	fmt.Println(b.Sum)

	l := new(LineCounter)
	io.Copy(l, &buf2)
	fmt.Println(l.Line)
}
