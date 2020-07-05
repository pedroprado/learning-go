package main

import (
	"errors"
	"fmt"
)

type myErr struct {
	text string
}

func (err myErr) Error() string {
	return err.text
}

func main() {

	myErr := myErr{
		text: "Meu err",
	}

	result := myErr.Error()

	fmt.Println(result)
	fmt.Printf("%T \n", myErr)

	commonErr := errors.New("Errors String")
	fmt.Println(commonErr)
	fmt.Printf("%T \n", commonErr)

	v := 4
	formattedErr := fmt.Errorf("Error number %v", v)
	fmt.Println(formattedErr)
	fmt.Printf("%T \n", formattedErr)
}
