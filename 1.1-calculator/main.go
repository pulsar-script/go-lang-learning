package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to simple terminal base calculator")

	fmt.Println("Enter two numbers")
	fmt.Println("Enter Number One")
	var num1 float64
	fmt.Scan(&num1)
	fmt.Println("Enter Number Two")
	var num2 float64
	fmt.Scan(&num2)

	//* --- Note on fmt.Scan and Pointers ---
	//
	//* Why we use `&` with `fmt.Scan`:
	//
	// 1. fmt.Scan needs to modify the original variable. Go is a "pass by value" language,
	//    meaning a function receives a copy of a variable's value, not the original.
	//
	// 2. The `&` symbol is the "address-of" operator. It gives us the memory address of a variable.
	//    By passing the address (`&variable`), we are giving fmt.Scan the exact location
	//    in memory where our variable lives.
	//
	// 3. This allows fmt.Scan to write the user's input directly into the original variable,
	//    overwriting its previous value, instead of just modifying a temporary copy.
	//
	// In short: Use `&` to pass the memory address so `fmt.Scan` can successfully store
	// the user's input in your variable.
	//
	// Example:
	//
	// var myNumber int
	// fmt.Scan(&myNumber) // Correct: Passes the address to store the value.
	//
	// fmt.Scan(myNumber)  // Incorrect: Passes a copy of the value, which fmt.Scan cannot modify.

	fmt.Printf("num1 = %v and num2 = %v", num1, num2)

	fmt.Print(" Menu \n Press Idx to perform that operation \n 1 Addition \n 2 Substration \n 3 Multiplication \n 4 Division \n")
	var selectedOption uint
	fmt.Scan(&selectedOption)

	var Result float64

	switch selectedOption {
	case 1:
		Result = adder(num1, num2)

	case 2:
		Result = suber(num1, num2)

	case 3:
		Result = muler(num1, num2)

	case 4:
		Result = divider(num1, num2)

	default:
		fmt.Println("Invalid Operation")
	}

	fmt.Printf("Result = %v", Result)

}

// func for add
func adder(v1 float64, v2 float64) float64 {
	return v1 + v2
}

// func for sub
func suber(v1 float64, v2 float64) float64 {
	return v1 - v2
}

// func for mult
func muler(v1 float64, v2 float64) float64 {
	return v1 * v2
}

// func for divider
func divider(v1 float64, v2 float64) float64 {
	// this is need to added
	if v2 == 0 {
		panic("Error: Cannot divide by zero")
	}
	return v1 / v2
}
