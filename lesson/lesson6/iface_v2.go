package main

import "fmt"

type Point struct {
	X, Y float64
}

type IIinstance interface {
	Instance() float64
}

type Path []Point

func (p *Path) Instance() float64 {
	return float64(1)
}

func main() {
	var i IIinstance //接口变量
	p := &Path{{1, 2}, {3, 4}}
	i = p
	fmt.Println(i)
}
