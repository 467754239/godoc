/*
	recover是让当前函数退出
	panic是让程序退出
*/
package main

import "fmt"

func f1() {
	var p *int
	fmt.Println(&p) // 空指针的地址
	fmt.Println(*p) // 空指针的值	--> 终极错误
}

func f2() {
	var arr [3]int
	fmt.Println(arr[2])
}

func f3() {
	var n int
	n = 0
	fmt.Println(10 / n)
}

func main() {
	panic("---exit---")
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	//var f func()
	//f = f1
	//f()

	f3()
}
