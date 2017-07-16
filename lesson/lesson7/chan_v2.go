package main

import "fmt"

func sum(s []string, c chan string) {
	var sum string
	for _, v := range s {
		sum += v + " "
	}
	c <- sum // send sum to c
}

func main() {
	s := []string{"hello", "golang", "c++", "python"}
	c := make(chan string)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x + y)
}
