package main

import (
	"bufio"
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

	buff := bufio.NewReader(conn) //读入缓存
	for {
		line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}
