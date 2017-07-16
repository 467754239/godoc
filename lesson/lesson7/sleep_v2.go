/*
	搞笑sleep排序
	主函数一旦退出 所有的协程都死掉了。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{2, 7, 1, 6, 4}
	for _, n := range s {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n)
		}(n)
	}
	time.Sleep(10 * time.Second)
}
