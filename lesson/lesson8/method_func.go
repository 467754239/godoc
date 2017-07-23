package main

import "fmt"

type Student struct {
	Name string
	Id   int
}

func (s *Student) Update(id int) {
	s.Id = id
}

func main() {
	// 静态绑定
	var f func(int)
	s := Student{Name: "zhengyscn"}
	f = s.Update
	f(2)
	fmt.Println(s)

	// 动态绑定/延迟绑定
	var f1 func(s *Student, id int)
	f1 = (*Student).Update
	f1(&s, 3)
	fmt.Println(s)

	// 延迟绑定
	s1 := Student{Name: "jack"}
	f1(&s1, 4)
	fmt.Println(s1)
}
