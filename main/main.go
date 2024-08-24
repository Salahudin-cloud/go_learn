package main

import (
	"fmt"
	//"example.com/hello_world"
	"log"

	"example.com/greetings"
	
)

func main() {

	/*
		set the properties of the predefined logger, including
		the log entery prefix and aa flag to disable printing
		the time, source file, and line number. 
	*/

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names 
	names := []string{"rerina", "udin"}


	// request a greetings message
	message, err := greetings.Hellos(names)
	/*
		if an error was returnes , print it to the conosle 
		and exit the programa
	*/

	if err != nil {
		log.Fatal(err)
	}

	/*
		if no error was returnes, print the returned message 
		to the console 
	*/

	fmt.Println(message)
	
}