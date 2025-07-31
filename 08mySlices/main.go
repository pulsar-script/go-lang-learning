package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Skices tutorial in golang")

	// slices are basiclly array with advance features

	//* 1) syntax for slice
	// - if you provide length [3], then it become normal array, dont provide length
	// - need to initialized immediately for this synatx

	var fruitLists = []string{"Apple", "Mango"}
	fmt.Printf("Type of fruitLists slice is %T \n", fruitLists)

	// how to add element in slices
	//! fruitLists[0] = "apple" we dont do this
	fruitLists = append(fruitLists, "Peach", "Banana")
	fmt.Println(fruitLists)

	//* Operations of slice

	// slice
	var newList1 = append(fruitLists[1:]) // skip 0th, start from 1st and rest of till the end
	fmt.Println(newList1)

	var newList2 = append((fruitLists[1:3])) // start from 1st and upto 2nd, end range is exculsive
	fmt.Println(newList2)

	var newList3 = append((fruitLists[:3])) // take all from start upto exculsive to end rang
	fmt.Println(newList3)

	//* 2) syntax for slice

	highscores := make([]int, 4) // this synatx declare array with 4 default length
	fmt.Println(highscores)
	fmt.Printf("Type of inital slice made by make() syntax is %T \n", highscores)

	highscores[0] = 456
	highscores[1] = 875
	highscores[2] = 121
	highscores[3] = 432
	// highscores[4] = 333
	//!panic: runtime error: index out of range [4] with length 4
	// Because this is array

	fmt.Println(highscores)

	// when add element
	// - make() re-allocate all memeory for all elements
	// - saves lots of memory
	// - save time
	// - help in performance optimization
	highscores = append(highscores, 765, 999, 111)
	fmt.Println(highscores)
	fmt.Printf("The type of variable is %T \n", highscores)

	//* Sorting
	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
}
