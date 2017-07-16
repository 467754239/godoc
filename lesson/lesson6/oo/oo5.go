/*
	不同的对象可以有相同的方法名
*/
package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p Point) ScaleByV1(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) ScaleByV2(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func ScaleByV1(p Point, factor float64) {
	p.X *= factor
	p.Y *= factor
}

func ScaleByV2(p *Point, factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	var factor float64
	p := Point{1, 2}
	factor = 2
	ScaleByV1(p, factor)
	fmt.Println(p)

	p = Point{1, 2}
	ScaleByV2(&p, factor)
	fmt.Println(p)

	p = Point{1, 2}
	p.ScaleByV1(factor)
	fmt.Println(p)

	p = Point{1, 2}
	p.ScaleByV2(factor)
	fmt.Println(p)
}
