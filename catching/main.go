package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// createFile()

	readFile()
}

//catching error from print
func printing() {
	n, err := fmt.Println("GO")
	if err != nil {
		fmt.Println("There is no error")
	}
	fmt.Println(n)
}

//catching error from file
func createFile() {
	f, err := os.Create("name.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	//close file
	defer f.Close()

	line := strings.NewReader("James Bond")

	io.Copy(f, line)
}

//catching erros from read file
func readFile() {
	file, err := os.Open("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	line, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(line))
}
