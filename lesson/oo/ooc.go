package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

type ClassRoom struct {
	students []Student
}

var classrooms = make(map[string]ClassRoom, 1)

func Select(args []string) error {
	var c ClassRoom
	var classname string
	classname = args[0]
	_, ok := classrooms[classname]
	if !ok {
		c.students = []Student{}
		classrooms[args[0]] = c
	}
	return nil
}

func (c *ClassRoom) Add(id int, name string) {
	c.students
}

//
//func (c *ClassRoom) list() {
//	c.students
//}

/*
	{"c1" : [{1, 'binggan'}, {2, 'binggan2'}]}
	{"c2" : [{1, 'binggan'}, {2, 'binggan2'}]}
*/

func main() {
	var className string
	var funcHander func(string) error
	actionmap := map[string]func([]string) error{
		"select": Select,
		"add":    funchandler.Add,
	}

	f := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		// 去除两端的空格和换行
		line = strings.TrimSpace(line)
		// 按空格分割字符串得到字符串列表
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		// 获取命令和参数列表
		cmd := args[0]
		args = args[1:]

		// 获取命令函数
		actionfunc := actionmap[cmd]
		if actionfunc == nil {
			fmt.Println("bad cmd ", cmd)
			continue
		}
		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}

	//var c ClassRoom
	//c = ClassRoom{}
	//c.Select("a1")

	//cr.students = []Student{{1, "abc"}, {2, "abc"}}
	//classrooms = make(map[string]ClassRoom, 1)
	//classrooms["aa"] = cr
	//fmt.Println(classrooms)

}

//func main() {
//	//var c ClassRoom
//	//c = ClassRoom{}
//	//c.Add()
//
//	/*
//		select reboot
//		add binggan1 1
//		select qinghua
//		add binggan1 1
//	*/
//	actionmap := map[string]func([]string) error{
//		"select":
//	}
//
//	f := bufio.NewReader(os.Stdin)
//
//	for {
//		fmt.Print("> ")
//		line, _ := f.ReadString('\n')
//		// 去除两端的空格和换行
//		line = strings.TrimSpace(line)
//		// 按空格分割字符串得到字符串列表
//		args := strings.Fields(line)
//		if len(args) == 0 {
//			continue
//		}
//		// 获取命令和参数列表
//		cmd := args[0]
//		args = args[1:]
//
//		// 获取命令函数
//		actionfunc := actionmap[cmd]
//		if actionfunc == nil {
//			fmt.Println("bad cmd ", cmd)
//			continue
//		}
//		err := actionfunc(args)
//		if err != nil {
//			fmt.Printf("execute action %s error:%s\n", cmd, err)
//			continue
//		}
//	}
//
//}
