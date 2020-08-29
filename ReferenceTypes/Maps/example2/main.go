package main

import "fmt"

//Adding elements to a nil map

//To add elements to a map, it NEED TO BE STARTED

func main() {

	var myGreeting map[string]string
	myGreeting = map[string]string{}

	myGreeting["Tim"] = "Good morning"

	fmt.Println(myGreeting)
}
