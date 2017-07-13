package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// 函数变成类方法只需要参数向前提一下
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))
}
