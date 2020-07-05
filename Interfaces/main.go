package main

import "fmt"

//Any other types that implement "getGreeting()" function, is also gonna be of type "bot"
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	//eb is of type betterBot
	printBetterBot(eb)
	//sb is not of type betterBot (it does not implements the interface)
	//printBetterBot(sb)
}

func printGreeting(b bot) {
	fmt.Printf("%T \n", b)
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}
