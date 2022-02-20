package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	Counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				Counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", Counter)
}

func TestRaceConditionWithMutex(t *testing.T) {
	Counter := 0

	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {

			mutex.Lock()
			for j := 0; j < 100; j++ {

				Counter++

			}

			mutex.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Counter : ", Counter)
}
