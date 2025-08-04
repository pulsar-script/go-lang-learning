package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(" ")
	fmt.Println("Welcome to BMI calculator")

	var height float64
	var weight float64
	fmt.Println("")
	fmt.Print("Enter Height in (meter) : ")
	fmt.Scan(&height)
	fmt.Println("")
	fmt.Print("Enter Weight in (kg) : ")
	fmt.Scan(&weight)

	calculatedBMI := calculatBMI(weight, height)
	fmt.Println("")
	fmt.Printf("Your BMI is %.2f", calculatedBMI)

}

func calculatBMI(v1 float64, v2 float64) float64 {

	if v2 == 0 {
		panic("Error : Zero cannot be divided")
	}

	result := v1 / math.Pow(v2, 2)

	return result
}

//
//* === Go `fmt.Scan` Functions ===
//
// These functions are used for reading input from the standard input (the console).
//
// A crucial point for all of them: you must pass the *address* of the variable
// where the input will be stored, using the ampersand `&` operator (e.g., `&myVariable`).
//
//
//* --- 1. `fmt.Scan` ---
//
// Description:
//   Reads space-separated values. It will read across multiple lines to find all the
//   required values. It stops reading when it encounters whitespace.
//
//* Syntax:
//   func Scan(a ...any) (n int, err error)
//
// Example:
//   var firstName, lastName string
//   fmt.Print("Enter your first and last name: ")
//   fmt.Scan(&firstName, &lastName) // Will correctly read "John" and "Doe" even if on separate lines.
//
//
//* --- 2. `fmt.Scanf` ---
//
// Description:
//   Scans input based on a specific format string. The input must match the format
//   exactly.
//
//* Syntax:
//   func Scanf(format string, a ...any) (n int, err error)
//
// Important Format Verbs:
//   - `%s`: string, stops at whitespace.
//   - `%d`: integer.
//   - `%f`: floating-point number.
//   - `%c`: single character (special case, see below).
//
// Example:
//   var name string
//   var age int
//   fmt.Print("Enter name and age (e.g., 'Name: John Age: 30'): ")
//   fmt.Scanf("Name: %s Age: %d", &name, &age)
//
//
//* --- 3. `fmt.Scanln` ---
//
// Description:
//   Reads space-separated values from a *single line*. It stops scanning immediately
//   when it encounters a newline character. If the line has more data than expected,
//   it will return an error.
//
//* Syntax:
//   func Scanln(a ...any) (n int, err error)
//
// Example:
//   var num1, num2 int
//   fmt.Print("Enter two numbers on one line: ")
//   fmt.Scanln(&num1, &num2) // Will succeed for "10 20", but fail for "10 20 30".
//
//
//* --- The `%c` Verb: A Special Case ---
//
// Important Point:
//   Unlike all other format verbs (`%s`, `%d`, etc.) which automatically skip
//   leading whitespace (spaces, tabs, newlines), `%c` does NOT. It reads the very
//   next character it finds in the input stream, whatever it may be.
//
// Example:
//   var char1, char2, char3 rune
//   fmt.Print("Enter three characters (e.g., 'a b'): ")
//   fmt.Scanf("%c%c%c", &char1, &char2, &char3)
//!   If the user inputs "a b", char1 gets 'a', char2 gets ' ' (a space),
//!   and char3 gets 'b'. The space is read as a character.
//
//
