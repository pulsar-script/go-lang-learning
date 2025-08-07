package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome to user input"
	fmt.Println(welcome)

	// input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza")

	// comma ok || comma err syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating", input)
	fmt.Printf("Type of the input is %T", input)

	// This is like try and catch
	// variation
	// input, _
	// input, err
	// _, err

}
