/*
	自定义error
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	err = errors.New("an error")
	fmt.Println(err)
	fmt.Println(err.Error())

	var cmd string
	cmd = "dfs"
	err = fmt.Errorf("bad command:%s\n", cmd)
	fmt.Println(err)
}
