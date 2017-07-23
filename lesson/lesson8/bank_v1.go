package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	lock  sync.Mutex
	money int
}

func (a *Account) DoPrepare() {
	time.Sleep(time.Microsecond)
}

func (a *Account) GetGongZi(n int) {
	a.money += n
}

func (a *Account) GiveWife(n int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
}

func (a *Account) Buy(n int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
	}
}

func (a *Account) Left() int {
	return a.money
}

func main() {
	var account Account
	c1 := make(chan int)

	account.GetGongZi(10)
	go func() {
		account.Buy(5)
	}()
	go func() {
		account.GiveWife(6)
	}()
	time.Sleep(200 * time.Microsecond)
	fmt.Println(account.Left())

}
