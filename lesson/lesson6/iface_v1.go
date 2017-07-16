package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// w -> socket
	fmt.Fprintf(w, "hello world %v\n", "gopher")
}

func main() {
	var (
		f   *os.File
		err error
	)

	f, err = os.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// f -> *os.File
	_, err = fmt.Fprintf(f, "hello world %v\n", "golang")
	if err != nil {
		log.Printf("return error %v\n", err)
	}
}
