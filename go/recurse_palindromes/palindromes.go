package main

import (
	"fmt"
	"vector"
)

func printString(p string, B []int, C []int, n int) {
	if B[n] != -1 {
		printString(p, B, C, B[n])
		end := n - B[n]
		fmt.Println(p[B[n]:end])
	} else {
		fmt.Println(p[0:n])
	}
}

func main() {
	testString := "tacocat"
	size := len(testString)
	fmt.Println("Size == ", size)

}
