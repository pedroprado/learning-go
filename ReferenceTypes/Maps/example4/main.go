package main

import (
	"fmt"
)

//Ranging through a map

func main() {

	var myGreetings map[int]string
	myGreetings = map[int]string{
		1: "Hello!",
		2: "Olá!",
		3: "Good morning!",
	}

	fmt.Println("ranging using keys only")
	for key := range myGreetings {
		fmt.Println(key, myGreetings[key])
	}

	fmt.Println("\nranging using keys and values")
	for key, value := range myGreetings {
		fmt.Println(key, value)
	}
}
