package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Data Race Condition happens when there are multiple threads (goroutines) accessing a SAHRED MEMORY SPACE
//In this case, multiple threads making computation over the same variable will lead to inconsistent data
//This happens due the different execution speed of those threads

//Solving Race Condition : mutex, atomic operations, channels

//Using Atomicity: makes an operation in a single instruction over a memory space

func main() {

	var counter int64 = 0

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {

			//this will make and atomic operation over the given memory space
			atomic.AddInt64(&counter, 1)

			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
