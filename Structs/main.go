package main

import "fmt"

//STRUCTS ARE VALUE TYPES

//This means that for updating it using a function we have to pass its address

//Keys are not indexed: we cannot iterate over a Struct

//We have to know (define) all the keys ate compile time

//Used to store date about "an Object"

type address struct {
	street string
	city   string
}

type person struct {
	name     string
	lastName string
	address
}

func main() {

	pedro := person{
		name:     "Pedro",
		lastName: "Prado",
		address: address{
			street: "José Maria de Souza",
			city:   "São Pedro",
		},
	}

	fmt.Printf("%+v \n", pedro)

	//Three ways of updating a field

	//1.
	//pedro.name = "Lucas"

	//2.
	//pedro.update("Lucas")

	//3.
	//(&pedro).update("PAULO")

	fmt.Printf("%+v \n", pedro)

}

//Passing address of the struct to be updated
func (p *person) update(newName string) {
	fmt.Printf("%+v\n", *p)
	p.name = newName
}
