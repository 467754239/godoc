/*
	闭包坑
*/
package main

import "fmt"

func v1() {
	var (
		i      int
		fslice []func()
	)

	for i := 0; i < 3; i++ {
		fslice = append(fslice, func() {
			fmt.Println(i)
		})
	}
	fmt.Println(&i)
	for _, f := range fslice {
		f()
	}
}

func v2() {
	var (
		i      int
		fslice []func()
	)

	for i := 0; i < 3; i++ {
		i := i
		fslice = append(fslice, func() {
			fmt.Println(i)
		})
	}
	fmt.Println(&i)
	for _, f := range fslice {
		f()
	}
}

func main() {
	v1() // 闭包坑
	fmt.Println("-------")
	v2() // 绕过闭包坑
}
