package main

import "fmt"

func main() {

	fmt.Println("Welcome to structures tutorial in go-lang")

	// using User
	hitesh := User{"Hitesh", "hit@go.dev", true, 42}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details %+v \n", hitesh) //%+v give key - value
	fmt.Printf("Name : %v \n", hitesh.Name)
	fmt.Printf("Email : %v \n", hitesh.Email)

}

// in go-lang we dont have class insted we have struct
// no-inheritance, no super or parent keywords

// struct declearation

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
