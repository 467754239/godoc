package main

import (
	"flag"
	"io"
	"log"
	"net"
	"sync"
)

var (
	target = flag.String("target", "www.baidu.com:80", "target host")
)

func handleConn(conn net.Conn) {
	remote_conn, err := net.Dial("tcp", *target)
	log.Printf("target:%s", *target)
	if err != nil {
		log.Print(err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(remote_conn, conn) // Connection closed or EOF
		defer remote_conn.Close()
	}()
	go func() {
		defer wg.Done()
		io.Copy(conn, remote_conn)
		defer conn.Close()
	}()
	wg.Wait()

}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}
