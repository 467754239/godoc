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

	// 定义waitgroup
	wg := new(sync.WaitGroup)
	wg.Add(2)

	account.GetGongZi(10)
	go func() {
		account.Buy(5)
		wg.Done()
	}()
	go func() {
		account.GiveWife(6)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(account.Left())

}
