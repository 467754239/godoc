/*
	map 原地修改
*/
package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

func main() {
	m := make(map[string]*Student)
	m["zhengyscn"] = &Student{
		Id:   1,
		Name: "zhengyscn",
	}

	fmt.Printf("%= v\n", m)
	m["zhengyscn"].Id = 12
	fmt.Printf("%= v\n", m)
}
