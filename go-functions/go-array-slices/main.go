package main

import "fmt"


func main() {
	// Arrays
	var fruitArr [2] string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Shorthand
	vegetableArr := [2]string{"Potatoes", "Carrots"}

	//Slices
	fruitSlice := []string{"Apples", "Oranges", "Grapes"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	fmt.Println(vegetableArr)
	fmt.Println(vegetableArr[1])

	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[2])

	fmt.Println(len(fruitArr))
	fmt.Println(len(fruitSlice))

	fmt.Println(fruitSlice[1:3])

}