package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i <= 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}
