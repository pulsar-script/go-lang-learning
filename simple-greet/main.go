// main.go
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	// The rootCmd is the entry point of our application
	var rootCmd = &cobra.Command{
		Use:   "greet",
		Short: "greet is a simple CLI for saying hi and hello",
		Long:  `A very simple command-line application built with Cobra to demonstrate the basics.`,
	}

	// --- Define our commands ---

	// 1. The 'hello' command
	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints 'Hello, World!'",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	// A variable to hold the value of the --shout flag
	var shout bool

	// 2. The 'hi' command
	var hiCmd = &cobra.Command{
		Use:   "hi [NAME]",
		Short: "Prints a greeting to the specified name",
		Args:  cobra.ExactArgs(1), // This command requires exactly one argument
		Run: func(cmd *cobra.Command, args []string) {
			// args[0] will contain the first argument, which is NAME
			greeting := fmt.Sprintf("Hi, %s!", args[0])

			// If the --shout flag was used, convert to uppercase
			if shout {
				greeting = strings.ToUpper(greeting)
			}
			fmt.Println(greeting)
		},
	}

	// --- Wire everything up ---

	// Add the --shout flag to the 'hi' command
	// The "s" is a shorthand for the flag (-s)
	hiCmd.Flags().BoolVarP(&shout, "shout", "s", false, "Shout the greeting in uppercase")

	// Add the 'hello' and 'hi' commands to our root command
	rootCmd.AddCommand(helloCmd, hiCmd)

	// Execute the root command to start the application
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

/*

====== How to use =======

# Run the hello command
./simple-greet hello
# Expected Output: Hello, World!

# Run the hi command with a name
./simple-greet hi Gemini
# Expected Output: Hi, Gemini!

# Run the hi command with the --shout flag
./simple-greet hi "Awesome User" --shout
# Expected Output: HI, AWESOME USER!

# You can also use the shorthand flag
./simple-greet hi "Awesome User" -s
# Expected Output: HI, AWESOME USER!


*/
