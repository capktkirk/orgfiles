package main

import (
	"fmt"
	//"math"
	//"io/ioutil"
)

type Matrix [][]bool

func printString(p string, B []int, C []int, n int) {
	if B[n] != -1 {
		printString(p, B, C, B[n])
		end := n - B[n]
		fmt.Println(p[B[n]:end])
	} else {
		fmt.Println(p[0:n])
	}
}

func buildMatrix(s string) Matrix {
	var ret Matrix
	testString := s
	size := len(testString)
	fmt.Println("Size == ", size)
	row := make([]bool, size)
	for i := 0; i < size; i++ {
		row = append(row, false)
	}
	for i := 0; i < size; i++ {
		ret = append(ret, row)
	}
	return ret
}

func main() {
	//var P [size][size]bool
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("There was an error. ", err)
	}
	palindrome := input
	P := buildMatrix(palindrome)
	size := len(palindrome)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Println("P[", i, "][", j, "] == ", P[i][j])
		}
	}

	for j := 0; j < size; j++ {
		for i := 0; i < size-j; i++ {
			if i < (i + j - 1) {
				var neighbor bool
				neighbor = P[i+1][i+j-1]
				P[i][i+j] = (palindrome[i] == palindrome[i+j]) && neighbor
			} else if i == i+j {
				P[i][i+j] = true
			} else if i == i+j-1 {
				P[i][i+j] = (palindrome[i] == palindrome[i+j])
			}
		}
	}

	var cuts []int
	var backtrack []int
	cuts[0] = 0
	backtrack[0] = -1
	printString(palindrome, backtrack[:], cuts[:], size)
}
