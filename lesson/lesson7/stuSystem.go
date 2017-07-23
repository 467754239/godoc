package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var classrooms map[string]*ClassRoom
var currentClassRoom *ClassRoom

type Student struct {
	Id   int
	Name string
}

type ClassRoom struct {
	//teacger  string
	students map[string]*Student // 成员变量
}

func (c *ClassRoom) MarshalJSON() ([]byte, error) {
	//m := make(map[string]interface{})
	//m["teacher"] = c.teacger
	//m["student"] = c.students
	//return json.Marshal(m)
	return json.Marshal(c.students)
}

func (c *ClassRoom) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &c.students)
}

func (c *ClassRoom) Add(name string, id int) error {
	if _, ok := c.students[name]; !ok { //key不存在
		c.students[name] = &Student{Id: id, Name: name}
	} else {
		fmt.Printf("student name:%s already exists.", name)
	}
	return nil
}

func (c *ClassRoom) Update(name string, id int) error {
	if stu, ok := c.students[name]; ok { //ok, key不存在
		//c.students[name].Id = id
		stu.Id = id
	} else {
		fmt.Printf("student name:%s not found.", name)
	}
	return nil
}

func (c *ClassRoom) List() error {
	for _, stuinfo := range c.students {
		fmt.Println(stuinfo.Id, stuinfo.Name)
	}
	return nil
}

func Save() error {
	buf, err := json.Marshal(classrooms)
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	return nil
}

func Load() error {
	return nil
}

func choose(args []string) error {
	name := args[0]

	if classroom, ok := classrooms[name]; ok {
		currentClassRoom = classroom
	} else {
		//
	}
	return nil
}

func add(args []string) error {
	name := ""
	id := 0
	currentClassRoom.Add(name, id)
	return nil
}

func main() {

	c1 := &ClassRoom{
		students: make(map[string]*Student), //一定要加逗号
	}
	c1.Add("zhengyscn", 1)
	c1.Add("mama", 2)
	c1.List()

	c2 := &ClassRoom{
		students: make(map[string]*Student), //一定要加逗号
	}
	c2.Add("zhengyscn", 1)
	c2.Add("mama", 2)
	c2.Update("zhengyscn", 100)
	c2.List()

	classrooms = make(map[string]*ClassRoom)
	classrooms["cl01"] = c1
	classrooms["cl02"] = c2

	if err := Save(); err != nil {
		log.Fatal(err)
	}

}
