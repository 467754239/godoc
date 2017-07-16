/*
	空接口可以接受任何类型
*/
package main

import "fmt"

type Point struct {
	X, Y float64
}

type Write interface {
	Write(p []byte) (int, error)
}

// 没有任何要求
// 完成的I
//type I interface {
//}

func main() {
	var I interface{}

	var n int
	I = n
	var s string
	I = s
	var p Point
	I = p
	fmt.Println(I)

}
