/*
	defer
		1. 关闭文件、关闭coker、关闭锁、关闭channel。
		2. 函数返回前执行.
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func v1() {
	defer func() {
		fmt.Println("defer lambda func")
	}()
	fmt.Println("hello world!")
}

func v2() {
	fd, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	for {
		buf := make([]byte, 1024)
		_, err = fd.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file error %v", err)
			return
		}
		fmt.Print(string(buf))
	}

}

func read(f *os.File) (string, error) {
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func read_v2(f *os.File) (string, error) {
	var total []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		total = append(total, buf[:n]...)
	}
	return string(total), nil
}

func v3() {
	f, err := os.Open("/etc/profile")
	if err != nil {
		log.Fatal(err)
	}
	var (
		content string
		retries int
	)
	retries = 3
	for i := 1; i <= retries; i++ {
		content, err = read(f)
		if err == nil {
			break
		}
		time.Sleep(time.Second << 1)
	}
	fmt.Println(content)
}

func main() {
	v1()
	v2()
	v3()
}
