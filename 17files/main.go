package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files manipulation tutorial in go-lang")

	content := "Your privacy is important to us. This privacy statement explains the personal data Microsoft processes, how Microsoft processes it, and for what purposes."

	// create file
	file, err := os.Create("./myDemoFile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)

	// close file
	defer file.Close()

	//file reading
	fileReader("./myDemoFile.txt")

}

// file reading function

func fileReader(fileName string) {

	// data come in bytes format
	databytes, err := os.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Content in byte slice: ", databytes)
	fmt.Println("Content in string: ", string(databytes))

}

// common error handling practies
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
