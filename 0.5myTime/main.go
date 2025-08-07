package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to Time package of golang")

	// current time
	presentTime := time.Now()

	// time with full syntax
	fmt.Println(presentTime)

	// time with Format package
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// creating time with own data
	createdTime := time.Date(2024, time.January, 02, 15, 45, 10, 0, time.UTC)
	fmt.Println("created time : ", createdTime)

	fmt.Println(createdTime.Format("01-02-2006 Monday"))

}

// command for help
//* go help

// commands to build application

// system details & properties for build
//* go env

//* go build ( This automatically identify your os)

// build for other os
//* GOOS="windows" go build
//* GOOS="linux" go build
//* GOOS="darwin" go build
