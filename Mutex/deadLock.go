package main

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock() // Ini di kunci oleh goroutine1

	fmt.Println("Lock user1 : ", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock() // ini di kunci oleh goroutine 2
	fmt.Println("Lock user2 : ", user2.Name)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func main() {

	user1 := UserBalance{
		Name:    "Mikhael",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Deni",
		Balance: 1000000,
	}

	go transfer(&user1, &user2, 100000) // user1 900000 user 2 1100000
	go transfer(&user2, &user1, 400000) // user1 1300000 user 1  700000

	time.Sleep(3 * time.Second)

	fmt.Println("Saldo user 1 : ", user1.Balance)
	fmt.Println("Saldo user 2 ", user2.Balance)

}
