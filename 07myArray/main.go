package main

import "fmt"

func main() {

	fmt.Println("Welcome to Array tutorial in golang")

	// declaring array

	var fruitsList [5]string
	fruitsList[0] = "Apple"
	fruitsList[1] = "Mango"
	// we skip 2nd place, lets see what happens
	fruitsList[3] = "Banana"

	fmt.Println("fruitsList : ", fruitsList)

	// o/p => fruitsList :  [Apple Mango  Banana ] there is one extra space for absence of 2nd element
	fmt.Println("fruitsList length: ", len(fruitsList))

	// declaring Array and assigning values

	var veggiesList = [3]string{"Potato", "Tomato", "Peanut"}
	fmt.Println("veggiesList : ", veggiesList)
	fmt.Println("VeggiesList length : ", len(veggiesList))
}
