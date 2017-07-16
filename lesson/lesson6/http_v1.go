package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", handle)
	log.Println("start http...")
	log.Fatal(http.ListenAndServe(":8001", nil))

}
