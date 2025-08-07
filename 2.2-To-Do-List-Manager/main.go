package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Todolist struct

var TodoList List

func main() {

	// laoding daat from file ( if data exists)
	TodoList.loadAndStore()

	// Creating ONE scanner to use for all inputs
	scanner := bufio.NewScanner(os.Stdin)

mainLoop:

	for {

		fmt.Println("")
		fmt.Println("Welcome to To-Do List Manager")
		fmt.Print("\nMenu 1:\n1. Add Task\n2. List All Task\n3. Exit")

		fmt.Println("")
		fmt.Print("Enter Your Option : ")

		//! fmt.Scan(&userOptionMenu1)
		//! Classical Error : we did not use fmt.Scan from now
		// because when we enter 1 in CLI, it actully send '1\n', Scan take 1 and left \n (in buffer) ,
		// then we our next user input func execute e.g. reader.ReadString() it take that \n automatically and stop. not ask for user to input
		// solution => use scanner

		// Read the whole line of user input
		scanner.Scan()          // this is not fmt.Scan() , this Scan is different method
		input := scanner.Text() // Get a text from scanner

		// Convert the text input to an int
		userOptionMenu1, err := strconv.Atoi(input)
		checkErrorNil(err)

		//*Test
		// fmt.Printf("%T", userOptionMenu1)

		// switch
		switch userOptionMenu1 {
		case 1:
			// Pass the scanner to the function so it can be reused
			addTask(scanner)
			continue mainLoop
		case 2:
			//TODO listTask
			listAllTasks(scanner)
			continue mainLoop
		case 3:
			//TODO exit
			break mainLoop

		default:
			fmt.Println("Please select valid option")
			//TODO continue loop
			continue mainLoop
		}

	}

	// saving data into file
	//! saveToFile(&TodoList)  => dont make func if yourare performing operation on struct, make method

	TodoList.saveToFile()

}

//* Task struct
//! Issue, need exported fields to work with marshelling

// type Task struct {
// 	ID          int
// 	description string
// 	isComplete  bool
// }

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

// List struct inside []Task (slice of Task) to hold all Task struct
// ? why insted of defining globle slice of Task []Task
// * Because it give use methods like addTask , listTasks
type List struct {
	Tasks []Task
}

// * Add Task feature
// 1. method - addTaskMethod
// 2. func - addTask

// * add task [method] -- method is general func for any stuct made using type List stuct
func (list *List) addTaskMethod(newTask Task) {

	// add
	list.Tasks = append(list.Tasks, newTask)

}

// * [func] for take input and add task to TODO list
// addTask now takes a scanner as an argument
func addTask(scanner *bufio.Scanner) {

addTaskLoop:
	for {

		// take user input
		fmt.Println("")
		fmt.Print("\nEnter Your Task here : ")

		//! Error
		//! reader := bufio.NewReader(os.Stdin)
		//! newTask, err := reader.ReadString('\n')

		//Using same sacnner to read
		scanner.Scan()
		description := scanner.Text()

		//Its good practice to trim any leading/trailing whitespaces
		description = strings.TrimSpace(description)

		if description == "" {
			fmt.Println("Task description cannot be empty")
			// return
			continue addTaskLoop
		}

		// create unique ID's in easy way
		//! nextID := 1
		//! This causing inconsistency in indexing

		var nextID int
		// *TEST
		fmt.Printf("Test 0.1 length is %v =>", len(TodoList.Tasks))

		if len(TodoList.Tasks) == 0 {
			// check list is empty
			nextID = 0

		} else {
			// we are adding 1 to total number of exiting task in TodoList ( length ) .Task slice element, basically next index .. 1,2,3 ... nextID = 4
			// [len(TodoList.Tasks)-1]

			// *TEST
			// fmt.Printf("Test 0.2 length is %v =>", TodoList.Tasks[len(TodoList.Tasks)].ID)

			fmt.Printf("")
			//! nextID = TodoList.Tasks[len(TodoList.Tasks)].ID + 1
			//* READ BELOW EXPLANANTION

			lastID := TodoList.Tasks[len(TodoList.Tasks)-1].ID //* talking about length it should be minimum 1 , not talking about index it is start from 0

			nextID = lastID + 1

			fmt.Printf("Test 0.3 length is %v =>", nextID)

		}

		newTaskStruct := Task{
			ID:          nextID,
			Description: description,
			IsComplete:  false,
		}

		TodoList.addTaskMethod(newTaskStruct)
		fmt.Println("Task added successfully!")
		fmt.Println("")
		fmt.Println("\nMenu 2 \n1.Add next Task \n2. Exit")
		fmt.Println("")
		fmt.Print("Enter your option here : ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		userOptionMenu2, err := strconv.Atoi(input)
		checkErrorNil(err)

		switch userOptionMenu2 {
		case 1:
			continue addTaskLoop

		case 2:
			break addTaskLoop

		default:
			fmt.Println("please enter valid option")
			continue addTaskLoop
		}

	}
}

// * List all tasks feature
// 1. method - listAllTasksMethod
// 2. func - listAllTasks

// * listAllTasksMthod [method] -- method is general func for all / any struct made using List struct
func (list *List) listAllTasksMethod() {
	for i, task := range list.Tasks {

		// completed mark
		var statusIcon string

		if task.IsComplete {
			statusIcon = "✓"
		} else {
			statusIcon = "✗"
		}

		fmt.Printf("\n %v. [%v] ID: %v Task: %v isComplete: %v\n", i, statusIcon, task.ID, task.Description, task.IsComplete)
		fmt.Println("-----------------------------------------")
	}
}

// * [func] list all tasks
func listAllTasks(scanner *bufio.Scanner) {

listTasksLoop:
	for {

		fmt.Println("")
		fmt.Println("=========================== TODO App =========================")
		TodoList.listAllTasksMethod()

		//menu 3
		fmt.Println("")
		fmt.Printf("\nMenu \n1. Delete \n2. Complete \n3. Exit\n")
		fmt.Println("")
		fmt.Print("Enter your option here : ")

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		userOptionMenu3, err := strconv.Atoi(input)
		checkErrorNil(err)

		switch userOptionMenu3 {
		case 1:
			//TODO delete task
			deleteTask(scanner)

		case 2:
			//TODO complete
			isCompleteTask(scanner)

		case 3:
			//TODO exit
			break listTasksLoop

		default:
			fmt.Println("Please enter valid option")
			continue listTasksLoop
		}

	}

}

//* Delete task feature
// 1. method
// 2. func

// // * [Method] for delete task
// func (list *List) deleteTaskMethod(id int) {

// 	// var fro index
// 	var Idx int

// 	// extract index from id
// idLoop:
// 	for i, value := range list.Tasks {

// 		if id == value.ID {
// 			Idx = i
// 			break idLoop
// 		}
// 	}

// 	// validation
// 	if Idx > len(list.Tasks) {
// 		fmt.Println("invalid Index, Task doesnot exits")
// 		return
// 	}

// 	// remove
// 	list.Tasks = append(list.Tasks[:Idx], list.Tasks[Idx+1:]...)

// }

//* [func] for delete task

func deleteTask(scanner *bufio.Scanner) {

deleteLoop:
	for {

		fmt.Println("")
		fmt.Print("\nMenu: \n1. Delete Task \n2. Exit\n")
		fmt.Print("\nEnter your option : ")

		scanner.Scan()
		optionInput := scanner.Text()
		optionInput = strings.TrimSpace(optionInput)

		userOptionMenu4, err := strconv.Atoi(optionInput)
		checkErrorNil(err)

		switch userOptionMenu4 {
		case 1:
		case 2:
			break deleteLoop
		default:
			fmt.Println("Plese select valid option")
			continue deleteLoop
		}

		fmt.Println("")
		fmt.Print("\nEnter the Id of task you want to delete here : ")
		scanner.Scan()

		input := scanner.Text()
		input = strings.TrimSpace(input)

		userDeleteId, err := strconv.Atoi(input)
		checkErrorNil(err)

		// T odoList.deleteTaskMethod(userDeleteIdx)  Not use
		// made common function to find Index by ID

		// *Test
		// fmt.Printf("Test 1 userDeleteId => %v & type => %T", userDeleteId, userDeleteId)

		givenIndex := TodoList.findIdxById(userDeleteId)
		// --- Then, in the code that CALLS this function: ---
		if givenIndex == -1 {
			fmt.Println("Invalid ID, task does not exist.")
			continue // Or handle the error appropriately
		}

		// error handling

		if givenIndex == 404 {
			fmt.Println("")
			fmt.Println("Invalid Index, Task does not found")
			continue deleteLoop
		}

		// delete code
		TodoList.Tasks = append(TodoList.Tasks[:givenIndex], TodoList.Tasks[givenIndex+1:]...)

	}

}

//* Complete task feature
// 1. method
// 2. func

// * [method]
func (task *Task) isCompleteTaskMethod(userBool bool) {

	if userBool {
		task.IsComplete = true
		statusIcon := "✓"
		fmt.Println("")
		fmt.Printf("\n %v is [%v] \n", task.Description, statusIcon)
	} else {
		statusIcon := "✗"
		task.IsComplete = false
		fmt.Println("")
		fmt.Printf("\n %v is [%v] \n", task.Description, statusIcon)
	}

}

// * [func] to complete or uncomplete task
func isCompleteTask(scanner *bufio.Scanner) {
isCompleteLoop:
	for {

		fmt.Println("")
		fmt.Print("\nMenu : \n1. Update Task \n2. Exit\n")
		fmt.Print("Enter your option here : ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		userOptionMenu5, err := strconv.Atoi(input)
		checkErrorNil(err)

		switch userOptionMenu5 {
		case 1:
		case 2:
			break isCompleteLoop

		default:
			fmt.Println("Invalid Option , Please enter valid option")
			continue isCompleteLoop
		}

		fmt.Println("")
		fmt.Print("\nEnter the Id of task that you want to edit : ")
		scanner.Scan()
		idInString := scanner.Text()
		idInString = strings.TrimSpace(idInString)
		idInInt, err := strconv.Atoi(idInString)
		checkErrorNil(err)

		givenIndex := TodoList.findIdxById(idInInt)

		// --- Then, in the code that CALLS this function: ---
		if givenIndex == -1 {
			fmt.Println("Invalid ID, task does not exist.")
			continue // Or handle the error appropriately
		}

		// error handling
		if givenIndex == 404 {
			fmt.Println("")
			fmt.Println("Invalid Index, task does not found")
			continue isCompleteLoop
		}

		// 0 / 1 to complete or notcomplete show task
		fmt.Println("")
		fmt.Println("Your Task to Update : ")
		if TodoList.Tasks[givenIndex].IsComplete {
			statusIcon := "✓"
			fmt.Println("")
			fmt.Printf("\n %v is [%v] \n", TodoList.Tasks[givenIndex].Description, statusIcon)
		} else {
			statusIcon := "✗"
			fmt.Println("")
			fmt.Printf("\n %v is [%v] \n", TodoList.Tasks[givenIndex].Description, statusIcon)
		}

		fmt.Println("")
		fmt.Print("\n #Note \n - enter '1' to complete task \n - enter '0' to uncomplete task\n")
		fmt.Println("")
		fmt.Print("\nEnter here your option : ")
		scanner.Scan()
		userBoolInputInString := scanner.Text()

		//! TrimSpace doesn't have side effects and its return value is ignored (SA4017)
		userBoolInputInString = strings.TrimSpace(userBoolInputInString)

		userBoolInputInInt, err := strconv.Atoi(userBoolInputInString)

		checkErrorNil(err)

		// bool var for user input

		var userBoolInputInBool bool

		// conversting int to bool
		if userBoolInputInInt == 1 {
			userBoolInputInBool = true
		} else {
			userBoolInputInBool = false
		}

		TodoList.Tasks[givenIndex].isCompleteTaskMethod(userBoolInputInBool)

	}

}

// * [method] Find Index by ID method for List struct type
func (list *List) findIdxById(id int) int {

	// *TEST
	// fmt.Printf("\nTest 2 value %v & type => %T \n", id, id)
	// *TEST
	// fmt.Printf("\nTest 2.2 value %v & type => %T \n", id-1, id-1)

	//! var Idx int
	//! Classical Error : The Problem: The "Zero Value" Trap
	//! Error = even for out of range index it , never show err
	//! => beacuse Idx initialize with 0
	// ( read more in down )

	//* Solution sentinel value
	var Idx int = -1

	// match ID with database
idLoop:
	for i, task := range list.Tasks {

		// *TEST
		// fmt.Printf("\nTest 3 value %v & type => %T \n", task.ID, task.ID)

		//! our id is out range so this if never become true, loop never break , it naturally end

		if task.ID == id {
			Idx = i
			// *TEST
			// fmt.Printf("\nTest 4 value %v & type => %T \n", Idx, Idx)
			break idLoop
		}
	}

	// *TEST
	// fmt.Printf("\nTest 5 value %v & type => %T \n", Idx, Idx)

	//! loop never break and Idx don't got index value, so its value is Idx = 0
	//! it never satisfy our if
	// validation for our index does that index is within rang
	// if Idx > len(list.Tasks) {
	// 	fmt.Println("Invalid Index, task does not found ")
	// 	return 404
	// }

	//! here Idx = 0 pass , we continously change 0th index task always
	return Idx

}

// error
func checkErrorNil(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
	}
}

// * [method] create & save to file
func (list *List) saveToFile() error {

	// convert it into JSON frmat so that can be written in bytes (disk only support bytes) and go []slice cannot written on disk directly
	jsonData, err := json.Marshal(list.Tasks)
	checkErrorNil(err)

	// *TEST
	// fmt.Printf("Test 6 listInJson => %v ", jsonData)

	// writing into file
	err = os.WriteFile("./todoList.json", jsonData, 0644)
	checkErrorNil(err)

	return nil

}

// * [method] to load data from file and store into go slice

func (list *List) loadAndStore() error {

	// reading file
	jsonData, err := os.ReadFile("todoList.json")
	checkErrorNil(err)

	// error handling for file not exits
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	// for other errors
	checkErrorNil(err)

	// check if file was empty
	if len(jsonData) == 0 {
		return nil
	}

	//* unmarshalling
	err = json.Unmarshal(jsonData, &list.Tasks)

	checkErrorNil(err)

	return nil
}

/*
================================================================================
* Go Slice Quick Reference: Understanding the 'index out of range' Panic //
================================================================================

This note explains one of the most common runtime errors in Go, especially
when working with collections of data like slices.

--------------------------------------------------------------------------------
* 1. THE ERROR MESSAGE
--------------------------------------------------------------------------------

panic: runtime error: index out of range [X] with length Y

- This means: "You tried to access an element at index 'X', but the slice
  only has a length of 'Y', so that index doesn't exist."

- Example from our project: `panic: runtime error: index out of range [0] with length 0`
  This meant we tried to access index 0 in a slice that was empty (length 0).


--------------------------------------------------------------------------------
* 2. THE CORE CONCEPT: Length vs. Index
--------------------------------------------------------------------------------

In Go, slices are "zero-indexed". This is the key to understanding the error.

- Length (len): The TOTAL number of elements in the slice. It's a simple count.
- Index: The "address" or "position" of an element. It ALWAYS starts at 0.

The last valid index in any slice is ALWAYS `len(slice) - 1`.

A VISUAL EXAMPLE:
Let's imagine this slice: `mySlice := []string{"Apple", "Banana", "Cherry"}`

- What is the LENGTH?
  `len(mySlice)` is 3. (There are three strings in it).

- What are the INDICES?
  The "addresses" of the elements are 0, 1, and 2.

   +----------+----------+----------+
   | "Apple"  | "Banana" | "Cherry" |  <-- The Elements (Length = 3)
   +----------+----------+----------+
       0          1          2         <-- The Valid Indices

- To access "Apple": `mySlice[0]`
- To access "Cherry" (the last element): `mySlice[2]`, which is `mySlice[len(mySlice) - 1]`


--------------------------------------------------------------------------------
* 3. THE MISTAKE THAT CAUSES THE PANIC
--------------------------------------------------------------------------------

The panic is almost always caused by trying to access `slice[len(slice)]`.
This index NEVER exists.

- Using our example: `mySlice[len(mySlice)]` would be `mySlice[3]`.
  As you can see from the diagram, there is no box at index 3. This causes the panic.

- The bug in our project was this line:
  `nextID = TodoList.Tasks[len(TodoList.Tasks)].ID`

  When `TodoList.Tasks` was empty, `len` was 0. The code tried to access `Tasks[0]`,
  which didn't exist, causing the `index out of range [0] with length 0` panic.


--------------------------------------------------------------------------------
* 4. THE SOLUTION: CORRECT PATTERNS
--------------------------------------------------------------------------------

Before accessing an element, especially the first or last, ALWAYS consider the empty case.

* Pattern 1: Safely accessing the LAST element of a slice.
func getLastElement(slice []string) {
    if len(slice) == 0 {
        fmt.Println("The slice is empty, cannot get the last element.")
        return
    }
    * If we get here, we know the slice is not empty.
    lastElement := slice[len(slice)-1]
    fmt.Println("The last element is:", lastElement)
}

* Pattern 2: Generating a new sequential ID (our project's specific problem).
func getNextID() int {
    var nextID int
    * ALWAYS handle the empty case first!
    if len(TodoList.Tasks) == 0 {
        * If there are no tasks, the first ID should be 1.
        nextID = 1
    } else {
        * If the slice is NOT empty, we can safely get the ID of the LAST task.
        lastTaskID := TodoList.Tasks[len(TodoList.Tasks)-1].ID
        nextID = lastTaskID + 1
    }
    return nextID
}


* FINAL RULE OF THUMB:
* Before you write `mySlice[i]`, mentally ask yourself:
* "Am I 100% sure this slice is not empty, and am I 100% sure 'i' is
* between 0 and len(mySlice)-1?"




/*
===================================================
The Problem: The "Zero Value" Trap
===================================================

* The bug is in these two lines:
var Idx int
return Idx

In Go, when you declare a variable without giving it an explicit value, it is automatically initialized to its "zero value". For an int, the zero value is 0.

* Let's trace what happens when you enter an ID that does not exist, like 99.

var Idx int runs. The variable Idx is created and its value is set to 0.

The for loop starts. It iterates through all your tasks.

The condition if task.ID == id (e.g., if task.ID == 99) will never be true.

Because the if condition is never met, the line Idx = i is never executed. The break is never hit.

The loop finishes naturally after checking every task.

What is the value of Idx now? It was never changed. It is still 0.

Your validation check runs: if Idx > len(list.Tasks). This becomes if 0 > len(list.Tasks), which is always false.

The function skips the if block and executes return Idx. It returns 0.

So, your function incorrectly reports that the task was found at index 0, which is misleading and wrong.

====================================================================
Option 1: The "Sentinel Value" (-1)
===================================================================

A common pattern in programming is to use a "sentinel value"—a special value that could never be a real result—to indicate failure.
For array indices, the standard sentinel value is -1, because a valid index can never be negative.

*/
