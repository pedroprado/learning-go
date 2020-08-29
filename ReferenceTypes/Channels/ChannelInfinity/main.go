package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

//Function Literal: is it like Anonymous Functions in Javascript

//Now we want to loop through the links, pinging them till there is an error
func main() {

	//fmt.Println(runtime.NumCPU())
	var numCPU = runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	//this is equivalent to "for{}"
	for l := range c {
		//we are not creatin infinite routines: the receiver operation is still blocking

		//Sleep is not good in the main routine: it is a good pratice for the main routine
		//always to be awake and receiving messages
		//--->>> time.Sleep(time.Second)

		//The solution is to use Function Literal
		go func(link string, c chan string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l, c)

	}

}

//Function that checks if the link is responding
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is up!")
	c <- link
}
