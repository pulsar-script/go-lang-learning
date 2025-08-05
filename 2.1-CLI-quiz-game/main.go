package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// globle level final question, options, answer slice of struct , []struct
// ! var FinalQuestionsSlice = []struct{}  [X]
var FinalQuestionsSlice []Question // correct syntax

// globel var for score
var Score int

func main() {
	fmt.Println("")
	fmt.Println("Welcome to CLI Quiz game")

	// opening file for reading
	file, err := os.Open("./quiz.csv")
	checkNilErr(err)
	defer file.Close()

	// creatng reader to read file
	reader := csv.NewReader(file)

	// reading record by record (row)
	for {
		record, err := reader.Read()
		// fmt.Println("", record) // Printing record that are in slice of string []string

		// The io.EOF error signals that we have reached the end of the file | Loop termination
		if err == io.EOF {
			break
		}

		// random err hadmling
		checkNilErr(err)

		// [test] print 1 record
		// fmt.Println("Record one : ", record[0])

		//! for _, record := range record {
		// unncessory loop, we are not extracting all question, [ [record], [record], [...], [...] ] , means [][]string ( slice of slice of string)
		//* we are extracting each row at time record, [ question, option1, option2, option3, correctAnswerString]

		// triming /r/n
		correctAnswerString := strings.TrimSpace(string(record[4]))
		correctAnswerInt, err := strconv.Atoi(correctAnswerString) // converting into int
		checkNilErr(err)

		optionsSlice := []string{string(record[1]), string(record[2]), string(record[3])}

		//! newQuestion := Question{record[0], record[1], record[2], record[3], record[4]}
		// getting erros we case - all record values are in bytes
		// - you are not following Questions struct structure
		// - not assigning new proccessed values , assinging old values
		// e.g. correctAnswerInt, err = strconv.Atoi(record[4]) // converting into int

		newQuestion := Question{string(record[0]), optionsSlice, correctAnswerInt}

		//! FinalQuestionsSlice.append(FinalQuestionsSlice, newQuestion)
		FinalQuestionsSlice = append(FinalQuestionsSlice, newQuestion)

		//! }
	}

	// Display questions
	showQuestions(FinalQuestionsSlice)

}

// Question struct
type Question struct {
	question      string
	options       []string
	correctAnswer int
}

// error handling
func checkNilErr(err error) {
	if err != nil {
		fmt.Println("")
		fmt.Println("Error : ", err)
	}
}

// funct to display questions one by one in proper format

func showQuestions(questionSlice []Question) {

	for idx, question := range questionSlice {

		fmt.Println("------------------------------------------------------------------")
		fmt.Printf("%v. %v\n", idx, question.question)
		fmt.Println("")
		fmt.Println("Options :")
		for i, option := range question.options {
			fmt.Printf("%v) %v\n", i, option)
		}
		fmt.Println("")
		getUserInputAndCheckAns(question.correctAnswer)

	}
	fmt.Println("")
	fmt.Println("Your Total Score is = ", Score)
}

// func for take input and check answer and + - score

func getUserInputAndCheckAns(correctAns int) {

optionLoop:
	for {

		// take user
		fmt.Println("")
		fmt.Print("Enter Your Index of answer here : ")
		var userInput int
		fmt.Scan(&userInput)

		// fmt.Printf("\nuserInput is %v and its type is %T\n", userInput, userInput)

		// validation
		if 0 <= userInput && userInput <= 2 {
			// checking correct answer
			if userInput == correctAns {
				fmt.Println("")
				fmt.Println("Congratulations, Your Answer is Correct, You score 1 mark")
				Score++
				break optionLoop
			} else {
				fmt.Println("")
				fmt.Println("oops! Wrong answer, you loose 1 mark")
				Score--
				break optionLoop
			}

		} else {
			fmt.Println("")
			fmt.Println("Invalid Option, please select option between 0, 1 ,2")
			continue optionLoop
		}

	}

}
