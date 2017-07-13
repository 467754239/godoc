package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// 函数变成类方法只需要参数向前提一下
func Distance(path []Point) float64 {
	var fsum float64
	for i := 0; i < len(path); i++ {
		fsum += path[0].Distance(path[1])
	}
	return fsum
}

func main() {
	path := []Point{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(Distance(path))
}
