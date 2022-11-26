package main

import (
	"fmt"
	"go-starter/greetings"
	"log"
)

func main(){
	fmt.Println("Hello world :D")

    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetingsError: ")
    log.SetFlags(0)

    // Request a greeting message.
	message, err := greetings.Hello("Gladys")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil{
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// Call other function from same 
	messagept := greetings.Ola("Gladys")

	fmt.Println(messagept)
}