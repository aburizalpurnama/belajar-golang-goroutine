package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	// lock
	// account.RWMutex.Lock()

	// write data
	account.Balance += amount

	// unlock
	// account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	// lock
	// account.RWMutex.RLock()

	// read data
	balance := account.Balance

	// lock
	// account.RWMutex.RUnlock()

	return balance
}

func TestWRMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", account.GetBalance())
}
