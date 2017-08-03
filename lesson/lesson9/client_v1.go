package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr) //Dial 拨号
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())

	// 第一个\r\n表示请求头结束
	// 第二个\r\n表示body结束
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("write size", n)

	buf := make([]byte, 4096)
	n, err = conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	// buf[:n]???
	// buf申请了4096字节，但是只读取n个字节，后面均为空那么就切片忽略。
	fmt.Println(string(buf[:n]))
}
