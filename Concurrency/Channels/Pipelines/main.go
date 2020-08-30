package main

import (
	"fmt"
)

//Pipelines: is a form of concurrency pattern

//It is a series of stages connected by channels

//Each stage is a group of gouroutines (in the same function - scope)

//The first stage = should have ONLY outbound channel

//Last stage = should only inbound channel

func main() {
	in := gen(1, 2, 3, 4, 5)
	out := sq(in)

	for v := range out {

		fmt.Println(v)
	}

}

func gen(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()

	return out
}

func sq(in chan int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			elem := v
			out <- elem * elem
		}
		close(out)
	}()

	return out
}
