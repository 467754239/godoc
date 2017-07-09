package main

import "fmt"

type Item struct {
	Id       int
	Price    float64
	Quantity int
}

type SpecialItem struct {
	Item
	catalogId int
}

func (item *Item) Cost() float64 {
	return item.Price * float64(item.Quantity)
}

type LuxuryItem struct {
	Item
	Markup float64
}

func (item *LuxuryItem) Cost() float64 {
	//return item.Item.Price * float64(item.Item.Quantity) * item.Markup
	return item.Item.Cost() * item.Markup
}

func main() {

}
