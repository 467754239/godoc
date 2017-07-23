/*
	序列化的坑
		结构体的成员变量首字母是否大写

	要序列化的对象如果实现MarshalJSON方法，会优先调用MarshalJSON方法
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name string
	id   int
}

func (s *Student) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.id)
}

func main() {
	s := &Student{
		Name: "zhengyscn",
		id:   1,
	}
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}
