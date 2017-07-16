package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Distance2Point(q *Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) Distance() float64 {
	return math.Hypot(p.X, p.Y)
}

type Path []*Point

func (p Path) Distance() float64 {
	var sum float64
	for i := 0; i < len(p)-1; i++ {
		sum += p[i].Distance2Point(p[i+1])
	}
	return sum
}

type IDistance interface {
	Distance() float64
}

func main() {
	var path Path
	var i IDistance

	p1 := &Point{1, 2}
	i = p1
	fmt.Println(i.Distance())

	p2 := &Point{3, 4}
	i = p2
	fmt.Println(i.Distance())

	p3 := &Point{5, 6}
	i = p3
	fmt.Println(i.Distance())

	path = make([]*Point, 0)
	path = append(path, p1, p2, p3)
	i = path
	fmt.Println(i.Distance())

}
