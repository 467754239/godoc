package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	// 内存buffer
	buf := new(bytes.Buffer)
	buf.WriteString(`
	hello gopher
	123456
	main new
	`)
	io.Copy(os.Stdout, buf)
}
