package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//making our custom type that implements the Writer Interface
type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// bs := make([]byte, 99999)

	//resp.Body is an interface that has to implement the Reader interfacer
	//The Reader interface says "how to read the data fetched from the server"
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//using the Copy function: it takes a Writer and a Reader=: outputs the data read to the terminal
	// io.Copy(os.Stdout, resp.Body)

	//using custom Writer
	lw := logWriter{}
	n, _ := io.Copy(lw, resp.Body)

	fmt.Println(n)

}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
