package main

import (
	"fmt"
)

//Maps are built over HASHTABLES
//Hashtables are a key/value datastructure
//The keys are calculated using HASH FUNCTIONS (that needs to be fast!)
//Indexes are not ordered (like array)

func main() {

	var colors = []string{"Red", "Blue", "Yellow", "Pink", "Purple", "Brown", "Green"}

	var buckets = 7

	var hashMap = make([]string, buckets)

	for _, color := range colors {
		hash := hashFunction(color, buckets)
		hashMap[hash] = color
	}
	fmt.Println(hashMap[3])

	fmt.Println(hashMap)
}

func hashFunction(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	fmt.Println("Bucket number :", sum%buckets, "For ", word)
	return sum % buckets
}
