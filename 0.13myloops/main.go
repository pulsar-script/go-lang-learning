package main

import "fmt"

func main() {

	fmt.Println("Welcome to Loops tutorial in go-lang")

	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "sunday"}
	fmt.Println(days)

	// syntax 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// syntax 2
	for i := range days {
		fmt.Println(days[i])
	}

	// syntax 3
	for index, value := range days {
		fmt.Printf("For Index %v value is %v \n", index, value)
	}

	// other operations

	// this synatx is kind a while loop
	rougueValue := 1
	for rougueValue < 10 {

		// if conditions with continue
		if rougueValue == 5 {
			rougueValue++ // this will skip 5
			continue
		}

		//if condition with break
		if rougueValue == 6 {
			fmt.Println("We are at index 6, now we are breaking loop")
			break
		}

		// if conditions with goto statement
		if rougueValue == 8 {
			goto sayHello
		}

		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

	// goto statement
	// label : operation
sayHello:
	fmt.Println("Hello !!")

}
