package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// slice
// [X] var ContactsList = make([]struct{}, 5)
var ContactsList []Contact //* Define slice out of any function to make it globally accessible

// reader
var Reader = bufio.NewReader(os.Stdin) // Globally accessible

func main() {

mainLoop: //* Giving label to for , we can give it to others also so we can identify and perfrom operation
	for {

		fmt.Println("")
		fmt.Println("Welcome to In-Memory Contact List")
		fmt.Println("1 Add Contact \n2 List All Contact \n3 Exit")

		fmt.Print("Enter Idx to perfrom that operation : ")

		userOptionString, err := Reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		//! strconv.Atoi: parsing "1\r\n": invalid syntax =>
		//the line ending is actually two characters: a carriage return (\r) followed by a newline (\n)
		//So, when you type 1 and press enter, the userOptionString variable doesn't just contain "1". It contains "1\r\n".

		userOptionString = strings.TrimSpace(userOptionString)

		userOptionInt, err := strconv.Atoi(userOptionString)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number (1, 2, or 3).")
			continue // This will skip the rest of the loop and start from the top
		}

		// fmt.Printf("%T", userOptionInt)

		switch userOptionInt {
		case 1:
			contactAdder()

		case 2:
			printAllList()

		case 3:
			fmt.Println("Exiting the program. Goodbye!")
			break mainLoop //* This 'break' statement will exit the for loop. otherwise it show err ,break won't work

		default:
			fmt.Println("Exiting the program. Goodbye!")
		}

	}

}

// struct
type Contact struct {
	name    string
	phoneNo int
}

// function to add new contact
func contactAdder() {

	fmt.Println("")
	fmt.Println("Enter Name : ")
	name, err := Reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter Phone Number : ")
	phoneNo, err := Reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// convert phone number into Int
	phoneNo = strings.TrimSpace(phoneNo)
	phoneNoInt, err := strconv.Atoi(phoneNo)
	if err != nil {
		panic(err)
	}

	//? Stuck here how to give new name to struct we make ?
	// => We dont need to assign new name (key) each time to struct
	// we simply use one common name (key)
	// then does it look like

	// { newContact :
	// newContact :
	// newContact :
	// newContact :
	// newContact : }

	//* No because, even we need to give name (key) while creating struct, but while appending in slice , slice datatype wont take name (key) , and this name (key) vanish after adding into slice (temporarly store)

	newContact := Contact{name, phoneNoInt}

	// add into struct
	ContactsList = append(ContactsList, newContact)

	fmt.Println("New Contact is Successfully added")

}

// func for print all lists
func printAllList() {
	for i := range ContactsList {
		fmt.Println("")
		fmt.Printf("%v : %v", ContactsList[i].name, ContactsList[i].phoneNo)
		fmt.Println("")
	}
}

//* Why os.Exit(0) is a good alternative

// This is precisely why using os.Exit(0) is often the preferred method for exiting a command-line program. It's more direct and avoids the need for a labeled break, making the code simpler and less prone to this kind of logical error. It explicitly says "I am terminating the entire program now."
// For your project, you now have two correct and idiomatic ways to achieve your goal:
// Labeled break: A more general-purpose way to exit an outer loop.
// os.Exit(0): A specific way to terminate the entire program.
