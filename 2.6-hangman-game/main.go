package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// reader for input
var Reader = bufio.NewReader(os.Stdin)

func main() {
	//TODO string of words
	wordsSlice := []string{"apple", "mango", "banana", "watermelon", "coconute", "pineapple"}

	//* TEST
	// fmt.Println("wordSlice =>", wordsSlice)

	// Generate random index to choose secret word
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	randomUniqueIndex := r.Intn(len(wordsSlice))

	//* TEST
	// fmt.Println("randomUniqueIndex =>", randomUniqueIndex)

	// secret word
	secretWord := wordsSlice[randomUniqueIndex]

	//* TEST
	// fmt.Println("secretWord =>", secretWord)

	// game score variables

	var GuessesLeft int = 10

	var discoveredWord []rune                // slice of rune
	var guessedLetters = make(map[rune]bool) // all user input letters => if correct guess then true | if incorrect guess then false

	// prepare discovered word
	for range secretWord {

		// discoverdWord = append(discoverdWord, "-")
		//! cannot use "-" (untyped string constant) as rune value in argument to append
		// In Go, anything in double quotes (" ") is a string.
		// Anything in single quotes (' ') is a rune.

		discoveredWord = append(discoveredWord, '_')
	}

	//game start
	fmt.Println("")
	fmt.Println("--- New Game ---")

	//game engine
gameEngine:
	for {

		//TODO Show states, word
		fmt.Println("")
		fmt.Println("Word : ", string(discoveredWord)) // -------
		fmt.Println("Guessed Left : ", GuessesLeft)    // 10
		// fmt.Println("Guessed Letters : ", guessedLetters)

		fmt.Println("Incorrect guess letters : ")
		for key, value := range guessedLetters {
			// fmt.Printf("%v = %v\n", string(key), value) // empty
			if !value {
				fmt.Printf(" %c ", key)
			}
		}

		//TODO take input
		fmt.Println("")
		fmt.Print("\nEnter your guess letter here : ")
		inputLetter, err := Reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		//trim
		inputLetter = strings.TrimSpace(inputLetter)

		// check if input is single character or not
		if len(inputLetter) > 1 {
			fmt.Println("")
			fmt.Println("Please enter only single character")
			continue gameEngine

		}

		fmt.Println(" guess letter in string = ", inputLetter)
		guessLetterInRune := []rune(inputLetter)[0] // converting string into slice of rune
		// fmt.Println("guess letter in rune = ", guessLetterInRune)

		// Use the format verb %c to print it as a character.
		// Use the format verb %q to print it as a single-quoted character.

		// checking if that letter is already guessed or not
		_, ok := guessedLetters[guessLetterInRune]

		// if yes used
		if ok {
			fmt.Println("")
			fmt.Printf("This letter %s is already used", string(guessLetterInRune)) // no double letters word .. try to keeping game simple
			continue gameEngine
		}

		// letter is not used
		guessedLetters[guessLetterInRune] = false // false means letter is not used into secret / discoverd word | input is correct but guess is not correct

		// lets check guess letter is correct or not ?

		for i, value := range secretWord {

			// fmt.Println("")
			// fmt.Printf(" %v = %c ", i, value)

			// check guess letter is correct or not
			if value == guessLetterInRune {
				// yes guess letter match in secret word
				discoveredWord[i] = guessLetterInRune // which showes our secret word, replace _ with correct guess letter
				guessedLetters[value] = true          // letter is used, so true
				// break guessLetterCheckerLoop // => now double letters words also work e.g. apple
			}
		}

		// lets check if guess letter is correct or not , if inccorrect then chances --

		isUsed := guessedLetters[guessLetterInRune]

		// * TEST
		// fmt.Println(" isUsed => ", isUsed)

		if !isUsed {
			// false mean it not used
			GuessesLeft--
		}

		// check for winning condition each time
		if secretWord == string(discoveredWord) {
			fmt.Println("")
			fmt.Println("Word : ", string(discoveredWord)) // -------
			fmt.Println("Guessed Left : ", GuessesLeft)    // 10
			fmt.Println("congratulations ! you win")
			fmt.Println("")
			return
		}

		// check loosing condition
		if GuessesLeft == 0 {
			fmt.Println("")
			fmt.Println("Word : ", string(discoveredWord)) // -------
			fmt.Println("Guessed Left : ", GuessesLeft)    // 10
			fmt.Println("You loose , out of guess")
			return
		}

	}

}
