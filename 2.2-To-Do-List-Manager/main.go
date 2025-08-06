package main

import (
	"bufio"
	"fmt"
	"os"
)

// Todolist struct

var TodoList List

func main() {
	// for {

	fmt.Println("")
	fmt.Println("Welcome to To-Do List Manager")
	fmt.Print("\nMenu 1:\n1. Add Task\n2. List All Task\n3. Exit")

	fmt.Println("")
	fmt.Print("Enter Your Option : ")
	var userOptionMenu1 int
	fmt.Scan(&userOptionMenu1)

	//*Test
	// fmt.Printf("%T", userOptionMenu1)

	// switch
	switch userOptionMenu1 {
	case 1:
		addTask()
	case 2:
		//TODO listTask
	case 3:
		//TODO exit
	}

	fmt.Println("")
	fmt.Println("Tasks", TodoList)

	// }
}

// Task struct
type Task struct {
	ID          int
	description string
	isComplete  bool
}

// List struct inside []Task (slice of Task) to hold all Task struct
// ? why insted of defining globle slice of Task []Task
// * Because it give use methods like addTask , listTasks
type List struct {
	Tasks []Task
}

// add task method

func (list *List) addTaskMethod(newTask Task) {

	// add
	list.Tasks = append(list.Tasks, newTask)

}

// func for take input and add task to TODO list

func addTask() {

	// take user input
	fmt.Println("")
	fmt.Print("\nEnter Your Task here : ")

	reader := bufio.NewReader(os.Stdin)
	newTask, err := reader.ReadString('\n')
	checkErrorNil(err)

	//TODO generate ID
	newTaskStruct := Task{
		ID:          1,
		description: newTask,
		isComplete:  false,
	}

	TodoList.addTaskMethod(newTaskStruct)
}

// error
func checkErrorNil(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
	}
}
