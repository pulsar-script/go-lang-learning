package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func main() {

	// initializing map
	contacts := make(map[string]Contact)

	// scanner
	scanner := bufio.NewScanner(os.Stdin)

	// // load data from file
	// err := loadAndStore(&contacts)
	// fmt.Println("Error : ", err)

	// Consider this:
	if err := loadAndStore(&contacts); err != nil {
		fmt.Printf("Fatal error loading contacts: %v\n", err)
		return // Exit if we can't load the file
	}

mainLoop:

	for {

		// list all contacts
		fmt.Println("")

		if len(contacts) == 0 {
			fmt.Println("No Contact Exits, Add Plz")
		} else {
			fmt.Println("===================================================")
			for key, value := range contacts {
				fmt.Println("")
				fmt.Printf("Name : %v\n", key)
				fmt.Print("Contact Details : ")
				fmt.Printf("%v %v %v\n", value.Name, value.Email, value.Phone)

			}
			fmt.Println("===================================================")

		}

		// add contact
		fmt.Println("")
		fmt.Print("Enter the Name of contact : ")
		scanner.Scan()
		nameInput := scanner.Text()

		fmt.Println("")
		fmt.Print("Enter the Email of contact : ")
		scanner.Scan()
		emailInput := scanner.Text()

		fmt.Println("")
		fmt.Print("Enter the Phone Number : ")
		scanner.Scan()
		phoneNoInput := scanner.Text()

		// duplicate name validation
		existedContact, ok := contacts[nameInput]

		if ok {
			fmt.Printf(" %v name is already exits", existedContact.Name) // if found existed contact it return in existedContact , ok is bool
			continue mainLoop
		}

		contacts[nameInput] = Contact{
			Name:  nameInput,
			Email: emailInput,
			Phone: phoneNoInput,
		}

		break mainLoop

	}

	err := saveToFile(&contacts)
	fmt.Println("Error : ", err)

}

// struct for storing contact
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// save data to file

func saveToFile(contactMap *map[string]Contact) error {

	jsonData, err := json.MarshalIndent(contactMap, " ", "  ")
	/*
		jsonData, _ := json.MarshalIndent(contacts, "", "  ")
		The first argument (contacts) is the data to marshal.
		The second argument ("") is a prefix to add to the beginning of every line. We usually leave this empty.
		The third argument ("  ") is the string to use for each level of indentation. Two spaces is a common choice.
	*/

	// other erros
	if err != nil {
		return err
	}

	// write file
	err = os.WriteFile("./contactList.json", jsonData, 0644)

	if err != nil {
		return err
	}

	return nil

}

// load data from file
func loadAndStore(contacts *map[string]Contact) error {

	// reader
	jsonData, err := os.ReadFile("./contactList.json")

	// end of file
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err != nil {
		return err
	}

	// check file is empty
	if len(jsonData) == 0 {
		return nil
	}

	// stroing it
	err = json.Unmarshal(jsonData, contacts)

	if err != nil {
		return err
	}

	return nil

}
