package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Functions in go-lang")

	greeter()

	result := adder(10, 5)
	fmt.Println(result)

	proResult := proAdder(2, 4, 7, 10, 6)
	fmt.Println(proResult)

	myMsg, randomNo := proGreeter("king")
	fmt.Printf("My message is %v and random number is %v \n", myMsg, randomNo)

	_, randNumber := proGreeter("My name is king") // _ use when you want to skip some returned value
	fmt.Printf("Random number is %v \n", randNumber)

}

// 1
func greeter() {
	fmt.Println("Hello")
}

// 2
func adder(v1 int, v2 int) int { // type safety on parameters and return value in functions know as function signatures
	return v1 + v2
}

// 3 when dont know number on parameters
// then we take it as slice
func proAdder(values ...int) int {
	total := 0
	for _, val := range values { // _ when we dont need key
		total += val
	}

	return total
}

// 4 when return multiple values
func proGreeter(msg string) (string, int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	randInt := r.Intn(10)
	return msg, randInt
}
