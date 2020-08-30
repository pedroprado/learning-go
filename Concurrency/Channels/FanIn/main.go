package main

import (
	"fmt"
	"sync"
)

//Fan out: ONE channel -> Multiple Channels
//Multiple functions that read from ONE channel, and send to another channels

//Fan in: Multiple Channels -> ONE channel
//ONE function that reads from MULTIPLE channels, and send to a single channel

func main() {

	c1 := sending(10)
	c2 := sending(20)

	fanIn := fanIn(c1, c2)

	for v := range fanIn {
		fmt.Println(v)
	}
}

func sending(n int) chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)

	}()
	return c

}

func fanIn(c1 chan int, c2 chan int) chan int {
	fanInChan := make(chan int)

	go func() {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			for v := range c1 {
				fanInChan <- v
			}
			wg.Done()
		}()

		go func() {
			for v := range c2 {
				fanInChan <- v
			}
			wg.Done()
		}()

		wg.Wait()
		close(fanInChan)
	}()

	return fanInChan
}
