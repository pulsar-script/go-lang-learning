package main

// Definition
// Those functions define in struct know as methods

// code from struct dir
import "fmt"

func main() {

	fmt.Println("Welcome to structures tutorial in go-lang")

	// using User
	hitesh := User{"Hitesh", "hit@go.dev", true, 42}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details %+v \n", hitesh) //%+v give key - value
	fmt.Printf("Name : %v \n", hitesh.Name)
	fmt.Printf("Email : %v \n", hitesh.Email)

	// accessing method
	hitesh.greeter()

	fmt.Printf("Email Before calling emailChanger method : %v \n", hitesh.Email)

	hitesh.emailChanger("knox@bin.go")

	fmt.Printf("Email after calling emailChanger method : %v \n", hitesh.Email) //? why it not change ?
	// because these methods use copy of struct
	// thats why for changing actual struct value, use use *pointers , thats why pointers concept exist

	// we use further

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

// method
//1
func (u User) greeter() {
	fmt.Println("Hello i am from greeter method")
}

//2
func (u User) emailChanger(newEmail string) {
	// try to change email field value in struct
	u.Email = newEmail
	fmt.Println("New Email is : ", u.Email)
}
