package main

import (
	"fmt"
	"time"
)

func v1() {
	timer := time.NewTicker(time.Second)
	cnt := 0
	for _ = range timer.C {
		cnt++
		if cnt > 10 {
			timer.Stop()
			return
		}
		fmt.Println("crontab")
	}
}

func v2() {
	timer := time.NewTicker(time.Second)
	go func() {
		time.Sleep(time.Second * 5)
		timer.Stop()
		return
	}()
	for _ = range timer.C {
		fmt.Println("crontab")
	}
}

func main() {
	v1()
	v2()
}
