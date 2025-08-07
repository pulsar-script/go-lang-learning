package main

import "fmt"

const LogginToken string = "qwertuudifjdjd" // public const , first capital letter
// means all files in this dir/ can access it

func main() {

	// variable
	var username string = "Asim"
	fmt.Println(username)
	fmt.Printf("The type of variable is %T \n", username) // Type Placeholder %T

	// bool
	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("The type of variable is %T \n", isLoggedin)

	// small int
	var smallNum uint8 = 255 // 0 to 255
	fmt.Println(smallNum)
	fmt.Printf("The type of variable is %T \n", smallNum)

	// small float
	var smallFloat float32 = 1200.125554333288 // o/p => 1200.1256
	fmt.Println(smallFloat)
	fmt.Printf("The type of variable is %T \n", smallFloat)

	// big float
	var bigFloat float64 = 1200.125554333288 // o/p => 1200.125554333288
	fmt.Println(bigFloat)
	fmt.Printf("The type of variable is %T \n", bigFloat)

	// default value and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("The type of variable is %T \n", anotherVariable)

	// implicit type
	var website = "www.animal&low.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	// public const
	fmt.Println(LogginToken)
	fmt.Printf("The type of variable is %T \n", LogginToken)

}
