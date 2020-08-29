package main

import "fmt"

//MAPS ARE REFERENCE TYPE
//This means that for updating it using a function, we do not need to worry about pointers

//Keys are indexed: we can iterate over a Map

//We do not have to know (define) all the keys ate compile time

//Used for storing closely related values

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)

	update(colors, "red", "A")

	fmt.Println(colors)

}

func print(colors map[string]string) {
	// fmt.Println(&colors)
	for key, value := range colors {
		fmt.Println(key, value)
	}
}

//updating: do not need to worry about pointers
func update(colors map[string]string, key string, value string) {
	colors[key] = value
}
