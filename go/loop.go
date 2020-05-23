package main

import "fmt"

func foo(x int) bool {
	if v := x / 2; v < 2 {
		return true
	}
	return false
}

func main() {
	//Three Component Loops
	//init (initialize the index), cond (condition for end of loop), and post (what to do if the loop doesn't end)
	fmt.Println("First we'll do a three component loop.")
	sum := 0
	for i := 1; i < 10; i++ {
		sum += 1
	}
	//We expect the print out to be 9.
	fmt.Println("Sum is : ", sum)
	fmt.Println("Next we'll do a for loop that is similar to a while loop.")
	//Go doesn't have a while loop, instead all loops are "for" loops.
	//This means that the init and post statements are optional.
	x := 1
	for x <= 5 {
		fmt.Println("loop #", x)
		x++
	}
	//For-each range loop. i is the index of the slice in the string array.
	//Note : Research slicing, create a document doing the different slices/approaches the same was as here.
	//In place of i you can do a _ and this will discard the results so you don't get an error about unused variables.

	testStrings := []string{"There", "Are", "Four", "Lights"}
	for i, s := range testStrings {
		fmt.Println(i, s)
	}

	//Exiting a loop, same as in C.
	var exitString string
	s1 := "a"
	s2 := "b"
	s3 := "n"

	for i := 1; i < 6; i++ {
		if i == 1 {
			exitString += s2
		}
		if i%2 != 0 {
			exitString += s1
		} else {
			exitString += s3
		}
	}
	fmt.Println(exitString)

	var sumCont int

	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		sumCont += i
	}

	fmt.Println("sumCont == ", sumCont)

	//Like with for, the if statement can have an initializing statement : See function example.
	if foo(sumCont) {
		fmt.Println("foo returned true.")
	} else {
		fmt.Println("foot returned false.")
	}
	fmt.Print("Input a number between 1 and 20 : ")
	var userInput int
	for {
		_, err := fmt.Scanln(&userInput)
		if err != nil || (userInput > 10 || userInput < 1) {
			fmt.Println("Enter a valid number")
		} else {
			fmt.Println("You entered : ", userInput)
			break
		}
	}
}
