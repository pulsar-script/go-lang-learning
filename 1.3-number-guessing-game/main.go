package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for {
		fmt.Println("========================================================================")
		fmt.Println("Welcome to Number guessing game ")

		// generate random number
		// Create a new random generator with a time-based seed
		r := rand.New(rand.NewSource(time.Now().UnixMicro()))
		SecretNumber := r.Intn(10) + 1

		// loop 5 chances
		for i := 1; i <= 5; i++ {

			// take input within range
			userNum := getNumberInRange(1, 10)
			// call checking func
			var isWin bool = checkGuessing(SecretNumber, userNum, 5-i)

			if isWin {
				break
			}
		}

		// play again ?
		var playAgain string
		fmt.Println("")
		fmt.Println("want to play again ? Y/n")
		fmt.Scan(&playAgain)

		// if playAgain == "Y" {
		//!  main()   <=  stack over flow error
		// } else {
		// 	fmt.Println("")
		// 	fmt.Println("Thank you for playing")
		// }

		if playAgain != "y" && playAgain != "Y" {
			fmt.Println("\nThank you for playing")
			break // Break out of the infinite 'for' loop //* solution => for stack over flow error
		}
	}

}

// check guessing
func checkGuessing(sysNumber int, userNumber int, idx int) bool {

	if sysNumber == userNumber {
		fmt.Println("---------------------------------------------------------------------")
		fmt.Printf("Congratulations! you win ,Number is %v", sysNumber)
		return true

	} else {
		fmt.Println("---------------------------------------------------------------------")

		if idx == 0 {
			fmt.Println("Sorry ! you loss the game ! Better Luck Next Time")
			return false
		}

		fmt.Printf("Oops! guess number is not match, Try again! you have left %v chances", idx)
		return false

	}

}

// func to set userinput rang
func getNumberInRange(v1 int, v2 int) int {

	// take user input
	fmt.Println("")
	fmt.Println("Enter Guessed number bewteen 1 - 5 : ")
	var userNum int
	fmt.Scan(&userNum)

	// check user input is it in rang
	if v1 <= userNum && userNum <= v2 {
		return userNum
	} else {
		fmt.Println("")
		fmt.Println("Error : Please enter number with in Range")
		//! getNumberInRange(v1, v2)  <= it should be in return , because calling func return int when it correct
		return getNumberInRange(v1, v2)

	}

	//! return userNum  <= no need to seperatly return , if-else return and func return is same one

}
