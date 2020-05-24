package main

import (
	"fmt"
	"os"
)

//Use defer like try catch blocks? Try writing a file with a bad file name.

func main() {
	var userInput string
	_, err := fmt.Scanln(&userInput)

	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	f, err := os.Create(userInput)
	if err != nil {
		panic("Cannot create specified file.")
	}
	defer f.Close()

	fmt.Fprintf(f, "Hello")
}

// Yes, give it a bad input and it will return the error message.
