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
	c := fanOut(20)

	var wg sync.WaitGroup
	wg.Add(2)

	go read1(c, &wg)

	go read2(c, &wg)

	wg.Wait()

}

func fanOut(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func read1(c chan int, wg *sync.WaitGroup) {

	for v := range c {
		fmt.Println("Reading 1: ", v)
	}

	wg.Done()
}

func read2(c chan int, wg *sync.WaitGroup) {

	for v := range c {
		fmt.Println("Reading 2: ", v)
	}

	wg.Done()
}
