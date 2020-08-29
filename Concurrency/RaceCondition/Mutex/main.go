package main

import (
	"fmt"
	"sync"
)

//Data Race Condition happens when there are multiple threads (goroutines) accessing a SAHRED MEMORY SPACE
//In this case, multiple threads making computation over the same variable will lead to inconsistent data
//This happens due the different execution speed of those threads

//Solving Race Condition : mutex, atomic operations, channels

//Using Mutex

func main() {

	var counter int = 0

	var wg sync.WaitGroup

	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			//locks the shared memory space => only the thread that gets the mutex is accessing that memory space
			mutex.Lock()
			temp := counter
			temp++
			counter = temp
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
