package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	var n time.Duration
	n = time.Hour*2 + time.Minute*30 + time.Second*1
	fmt.Println(int64(n))
	fmt.Println(n.String())

	t := time.Now()
	t1 := t.Add(time.Hour)
	fmt.Println(t1.Sub(t))
	fmt.Println(t)
	fmt.Println(t.Add(time.Hour))
	fmt.Println(t.Add(-time.Hour))

	// 时间长度、区间 Duration
	time.Sleep(time.Second * 3)
	time.Sleep(time.Minute * 3)
	time.Sleep(time.Hour * 3)
	fmt.Println("End")
	fmt.Println(time.Now())
	fmt.Println(time.Now())
}
