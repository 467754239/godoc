package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	// "http://www.vipkid.com.cn/path/1.jpg?a=1&c=123#untar"
	// ftp://zhengyansheng:123456@www.vipkid.com.cn/
	s := os.Args[1]
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Scheme", u.Scheme)
	fmt.Println("Host", u.Host)
	fmt.Println("Path", u.Path)
	fmt.Println("RawQuery", u.RawQuery)
	fmt.Println("user", u.User)
	fmt.Println("xx", u.Fragment)
}
