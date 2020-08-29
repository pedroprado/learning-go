package main

import (
	"fmt"
	"net/http"
	"runtime"
)

//Goroutines: use multiple cpu cores to execute code in paralel

//Func main is the "Main Goroutine"
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

	for _, link := range links {
		go checkLink(link, c)
	}

	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)

	}

}

//Function that checks if the link is responding
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- "Error"
		return
	}
	fmt.Println(link, " is up!")
	c <- "Done"
}
