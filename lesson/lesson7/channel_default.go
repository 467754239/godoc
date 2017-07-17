/*
	当select中所有的channel都不通的时候走default
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Microsecond) //time = time.NewTicker(1000*time.Microsecond).c
	boom := time.After(5000 * time.Microsecond)

	for {
		select {
		case <-tick:
			fmt.Println("滴答...")
		case <-boom:
			fmt.Println("嘣!!!")
			return
		default:
			fmt.Println("吃一口面")
			time.Sleep(300 * time.Microsecond)
		}
	}
}
