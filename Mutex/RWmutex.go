package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (ba *BankAccount) addBalance() {
	ba.RWMutex.Lock()
	ba.Balance++
	ba.RWMutex.Unlock()
}

func (ba *BankAccount) getBalance() int {
	ba.RWMutex.RLock()
	account := ba.Balance
	ba.RWMutex.RUnlock()
	return account

}

func main() {

	account := BankAccount{}

	for i := 0; i < 100; i++ {

		go func() {
			for i := 0; i < 100; i++ {
				account.addBalance()
				fmt.Println(account.getBalance())
			}
		}()
	}

	time.Sleep(4 * time.Second)
}
