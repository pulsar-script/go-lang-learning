package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to Switch and cases tutorial in go-lang")

	// Ludo game
	// generate randome numbers using rand in math
	// we have rand in crypto also

	// Create a new random generator with a time-based seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random integer between 0 to math.MaxInt
	randomInt := r.Int()
	fmt.Println("Random integer : ", randomInt)

	// Generate a random integer between 0 to 99
	randomIntn := r.Intn(100)
	fmt.Println("Random integer between 0-99 : ", randomIntn)

	//Generate a random float between 0.0 to 1.0
	randomFloat := r.Float64()
	fmt.Println("Random float 0.0-1.0 : ", randomFloat)

	//Generate a random number in a specific range (e.g., 10-12)
	min := 10
	max := 20
	randomRange := r.Intn(max-min+1) + min
	fmt.Println("random number between 10-20 : ", randomRange)

	// game
	fmt.Println("Welcome to Ludo Game")

	diceNumber := r.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("No is 1, you can open box")
	case 2:
		fmt.Println("No is 2, sorry try again")
	case 3:
		fmt.Println("No is 3, sorry try again")
	case 4:
		fmt.Println("No is 4, sorry try again")
		fallthrough
	case 5:
		fmt.Println("No is 5, sorry try again")
	case 6:
		fmt.Println("No is 6, sorry try again")
	default:
		fmt.Println("Something Wrong")
	}

	// fallthrough
	// if that case match then its, executing part execute, then case below it also execute

}
