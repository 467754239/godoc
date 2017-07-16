/*
	主动暴露当前http状态
	当前接受多少个http请求，成功多少、失败多少、请求多少。
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Monitor struct {
	counter int
}

func (m *Monitor) Run() {
	for {
		time.Sleep(time.Second)
		m.counter++
	}
}

func (m *Monitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "count:%d\n", m.counter)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	var m Monitor
	http.HandleFunc("/", handle)
	http.Handle("/monitor", &m)
	go m.Run()
	log.Println("start http...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
