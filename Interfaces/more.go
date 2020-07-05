package main

import (
	"errors"
	"fmt"
)

type user struct {
	name string
	age  int
}

//Now, for a value to be of type "betterBot", this value has to implement all the functions
//of the interface "betterBot"
type betterBot interface {
	greeting(string, int) (string, error)
	getBotVersion() float64
	respondToUser(user) string
}

//Implementi all the 3 functions of the "betterBot" interface
//Now my englishBot is also of type "betterBot"
//Now, any function that receives the "betterBot" interface as an argument can
//also receive a value of type englishBot as an argument
func (eb englishBot) greeting(s string, i int) (string, error) {

	len := len(s)
	err := errors.New("")
	if len <= 0 {
		err = errors.New("Invalid string")

	}
	return s, err
}

func (eb englishBot) getBotVersion() float64 {
	return 4.14
}

func (eb englishBot) respondToUser(u user) string {
	return "Sent"
}

func printBetterBot(bb betterBot) {
	fmt.Println("This is of type betterBot!")
}
