package main

import "fmt"

//STRINGS ARE VALUE TYPES

//This means: to modify a value string using a function, we have the address of the string!

func main() {

	name := "Pedro"
	println(&name)
	update(&name)

	fmt.Println(name)
}

func update(name *string) {
	//Everything in Go is "passed by value"
	//Even the memory address (passed as a pointer) is passed and stored in another memory space
	println(&name)
	//Change the value stored in that address
	*name = "Lucas"
}
