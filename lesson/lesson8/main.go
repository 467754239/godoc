package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	s := "http://51reboot.com/course/go/"
	dir := filepath.Dir(s)
	name := filepath.Base(s)
	fullname := filepath.Join(dir, name)
	fmt.Println(dir)
	fmt.Println(name)
	fmt.Println(fullname)

}
