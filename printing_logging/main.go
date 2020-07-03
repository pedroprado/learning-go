package main

import (
	"log"
	"os"
)

func main() {

	logs := openLogFile()
	defer logs.Close()
	log.SetOutput(logs)

	file := normalErr()
	defer file.Close()

	// file := fatalErr()
	// defer file.Close()

	// file := panicErr()
	// defer file.Close()

}

func normalErr() *os.File {
	file, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("Error: ", err)
	}
	return file
}

func fatalErr() *os.File {
	file, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("Fatal: ", err)
	}
	return file
}

func panicErr() *os.File {
	file, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln("Panic: ", err)
	}
	return file
}

func openLogFile() *os.File {
	//Open/Create a file for logs
	logs, err := os.Create("logs.txt")
	if err != nil {
		log.Println("Error: ", err)
	}
	return logs
}
