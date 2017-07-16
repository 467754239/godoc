package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type LineCounter struct {
	Line int
}

func (l *LineCounter) Write(p []byte) (int, error) {
	l.Line += len(strings.Split(string(p), "\n"))
	return len(p), nil
}

func main() {
	l := new(LineCounter)
	io.Copy(l, os.Stdin)
	fmt.Println(l.Line)
}
