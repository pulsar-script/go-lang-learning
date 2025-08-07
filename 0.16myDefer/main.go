package main

import "fmt"

func main() {
	// defer use to skip that line from line by line execution and execute it at very end

	//1
	defer fmt.Println("World")
	fmt.Println("Hello")

	//2
	// order of multiple defer statements
	// LIFO : Last in first out
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Four") // non-defer statment do not follow LEFO
	// Four
	// Three
	// Two
	// One

	//3
	deferTestFunc()

}

func deferTestFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // temporary stack store in memory, 0, 1, 2, 3, 4
		// then print 4, 3, 2, 1, 0

	}
}
