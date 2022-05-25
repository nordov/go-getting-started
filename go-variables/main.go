package main

import "fmt"

// Global scope variables
var fname string = "David" 
var lname = "Odio" 			// It can infer it is a string
var llname = "Vargas"		// Shorthand for variable declaration

func main() {

	// Shorthand, can only be done inside function, no globally
	fullName := fname + " " + lname + " " + llname

	var age = 37

	const isCool= true 

	fmt.Println(fname, lname)
	fmt.Println("the full name is ", fullName)
	fmt.Printf("The variable age of value %v is of type %T\n", age, age)
	fmt.Printf("The variable age of value %v is of type %T\n", isCool, isCool)

}