/*
	slice、struct排序
*/
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Id   int
	Name string
}

func v1() {
	var s = []int{3, 5, 2, 1, 9, 6}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
}

func v2() {
	var s = []int{3, 5, 2, 1, 9, 6}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(s)
}

func v3() {
	ss := []Student{}
	ss = append(ss, Student{
		Id:   2,
		Name: "aa",
	})
	ss = append(ss, Student{
		Id:   1,
		Name: "cc",
	})
	ss = append(ss, Student{
		Id:   3,
		Name: "bb",
	})
	fmt.Println(ss)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Id < ss[j].Id
	})
	fmt.Println(ss)
}

func main() {
	v1() // slice正序
	v2() // slice倒叙
	v3() // struct排序
}
