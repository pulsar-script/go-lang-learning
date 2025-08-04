package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("")
	fmt.Println("Welcome to Word & Character Counter")

	fmt.Println("")
	fmt.Println("Enter your content")

	//input
	reader := bufio.NewReader(os.Stdin)

	// comma || err
	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	// character count
	charCount := len(input)

	// word count
	stringSlice := strings.Fields(input)
	wordCount := len(stringSlice)

	fmt.Println("")
	fmt.Println("Your Content = ", input)
	fmt.Println("Character Count = ", charCount)
	fmt.Println("Word Count = ", wordCount)

}
