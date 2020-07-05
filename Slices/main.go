package main

import "fmt"

//SLICES ARE REFERENCE TYPES!

//When we create a slice, GO runtime creates:
//1.The slice, which has: Pointer (Ptr), capacity and length
//2.An array which contains the VALUES(and which is pointed by Ptr)

func main() {

	//Go is a PASS BY VALUE language: it means, when we pass an argument, this argument is a
	//copy of the original value
	//Thus, we cannot modify an argument value without passing its reference

	//BUT:
	//Slices are different
	mySlice := []string{"Hi", "from", "Go", "lang"}

	fmt.Println(mySlice)

	//We would expect that "mySlice" would not be modified,
	//But: when whe pass the slice to a function, the "slice" is copied
	//but the pointer (ptr) to the underlyng array (which contains the values) is not!
	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Hello"
}
