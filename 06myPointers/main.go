package main

import "fmt"

func main() {

	fmt.Println("Welcome to Pointers tutorial in go")

	// how to create / declare pointer but not assign address < nil > (empty pointer)
	var ptr1 *int
	fmt.Println("ptr1 : ", ptr1)

	// how to create pointer and assign address
	myNumber := 23

	var ptr2 = &myNumber
	fmt.Println("ptr2 store address : ", ptr2)       // access address
	fmt.Println("ptr2 store address data : ", *ptr2) // access data

	// some operation on data accessing through ptr2
	*ptr2 = *ptr2 + 2
	fmt.Println("updated value of variable : ", myNumber)

}
