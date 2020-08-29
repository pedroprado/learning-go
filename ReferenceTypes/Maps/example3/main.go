package main

import (
	"fmt"
)

//Checking if a element exists in map : "comma ok idiom"

func main() {

	var myGreetings map[int]string

	myGreetings = map[int]string{}

	myGreetings[1] = "Hello!"
	myGreetings[2] = "Good morning!"

	val, exists := myGreetings[2]
	fmt.Println(`it should return "Good morning" and true`)
	fmt.Println("Got val: ", val)
	fmt.Println("Got exists: ", exists)

	val, exists = myGreetings[3]
	fmt.Println(`it should return "empty string" and false`)
	fmt.Println("Got val: ", val)
	fmt.Println("Got exists: ", exists)

}
