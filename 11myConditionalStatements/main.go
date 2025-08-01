package main

import "fmt"

func main() {

	fmt.Println("Welcome to conditional statments tutorials in go-lang")

	loginCount := 20
	var result string

	// syntax 1
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = " Exact 10"
	}

	fmt.Println(result)

	// syntax 2
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// syntax 3
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is more than 10")
	}

}
