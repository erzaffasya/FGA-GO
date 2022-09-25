package main

import (
	"fmt"
)

func main() {
	// +nil, hemat memory
	// start ex pointer
	// var firstPerson string = "John"
	// var secondPerson *string = &firstPerson

	// fmt.Println("FirstPerson (value) :",firstPerson)
	// fmt.Println("FirstPerson (memory address) :",&firstPerson)
	// fmt.Println("SecondPerson (value) :",*secondPerson)
	// fmt.Println("SecondPerson (memory address) :",secondPerson)
	
	// fmt.Println("\n",strings.Repeat("#",30),"\n")
	
	// *secondPerson = "Doe"
	
	// fmt.Println("FirstPerson (value) :",firstPerson)
	// fmt.Println("FirstPerson (memory address) :",&firstPerson)
	// fmt.Println("SecondPerson (value) :",*secondPerson)
	// fmt.Println("SecondPerson (memory address) :",secondPerson)
	// end ex pointer

	// start ex pointer as param
	var a int = 10
	fmt.Println("Before: ",a)

	changeValue(&a)
	fmt.Println("After: ",a)
}

func changeValue(number *int)  {
	*number = 20
}