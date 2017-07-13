package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	// 直接指针
	p := &Point{1, 2}
	p.ScaleBy(2)
	fmt.Println(p)

	// 声明结构体后再指针指向
	p1 := Point{1, 2}
	p2 := &p1
	p2.ScaleBy(2)
	fmt.Println(p)

	// 使用结构体调用，隐式取地址
	p3 := Point{1, 2}
	p3.ScaleBy(2) // &p3.ScaleBy(2)
	fmt.Println(p)
}
