package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
	message, err := greetings.Hello("Thomas")
	if err != nil {
		// If an error was returned, print it to the console and exit.
		log.Fatal(err)
	}
	// If no error was returned, print the message to the console.
	fmt.Println(message)
}
