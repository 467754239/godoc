/*
	不同的对象可以有相同的方法名
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// type关键字用来定义类型的
type Path []Point

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (path Path) Distance() float64 {
	var fsum float64
	for i := 0; i < len(path); i++ {
		fsum += path[0].Distance(path[1])
	}
	return fsum
}

func main() {
	path := Path{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(path.Distance())

}
