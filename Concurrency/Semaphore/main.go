package main

import (
	"fmt"
)

//Semaphore: it is a flag (variable) used to control access to some resource (other variables)

//Semaphores can be used to syncronize multiple threads execution (like the WaiGroup)

func main() {

	c := make(chan int)

	semaphore := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		semaphore <- true

	}()

	go func() {
		for i := 10; i < 20; i++ {
			c <- i
		}
		semaphore <- true

	}()

	go func() {
		<-semaphore
		<-semaphore
		close(c)

	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finished")
}
