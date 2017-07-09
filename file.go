/*
	读取文件的方式
		1. 按行读取
		2. 按块读取
		3. 全部读取

	文件读取的多种方式
		1. file.Read
		2. ioutil.ReadFile
		3. bufio.Scanner
		4. bufio.Reader
		5. io.Copy


*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func v1() {
	var f *os.File
	f, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Printf("open file error:%v", err)
	}
	defer f.Close()
}

func v2() {
	f, err := os.Open("defer.go")
	if err != nil {
		return
	}
	read(f)
}

func read(f *os.File) (string, error) {
	var total []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			return "", err
		}
		if err == io.EOF {
			break
		}
		total = append(total, buf[:n]...)
	}
	return string(total), nil
}

func readTerminal() {
	fd, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	/*
		-- 流式读取
		裸读取、很少使用
		按块的方式读取
		速冻慢、效率低
	*/
	buf := make([]byte, 4096)
	n, _ := fd.Read(buf)
	fmt.Println(buf[:n])

	/*
		-- 流式读取
		加上buffer的读取，很高效。
	*/
	r := bufio.NewReader(fd)
	n, err = r.Read(buf)
	fmt.Println(buf[:n])

	/*
		按行读取、按分隔符读取
		内置buffer
	*/
	r1 := bufio.NewScanner(f)

	/*
		小文件一次性读取
	*/
	b, err = ioutil.ReadFile("filename")
	b, err = ioutil.ReadAll(fd)

}

func main() {
	v1()
	v2()
}
