//Recursive Palindrome Cutting

package main

import (
	"fmt"
	"math"
	//"io/ioutil"
)

type Matrix [][]bool

type pal_struct struct {
	mtx       Matrix
	cuts      []int
	backtrack []int
	size      int
}

func printString(palindrome string, p_data pal_struct, n int) {
	if p_data.backtrack[n] != -1 {
		idx := p_data.backtrack[n]
		printString(palindrome, p_data, idx)
		fmt.Println(palindrome[idx:(p_data.size-idx)], "\n")
	} else {
		fmt.Println(palindrome[0:n])
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

func makeTables(P Matrix, cuts []int, backtrack []int, size int) pal_struct {
	var pal_data pal_struct
	for i := 1; i <= size; i++ {
		min := math.MaxInt64
		choice := 0
		for j := 0; j < i; j++ {
			bool_right := cuts[j] < min
			if P[j][i-1] && bool_right {
				min = cuts[j]
				choice = j
			}
		}
		cuts[i] = min + 1
		backtrack[i] = choice
	}
	fmt.Println(cuts)
	fmt.Println(backtrack)
	pal_data.mtx = P
	pal_data.cuts = cuts
	pal_data.backtrack = backtrack
	pal_data.size = size
	return pal_data
}

func main() {
	//var P [size][size]bool
	/*
		var input string
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println("There was an error. ", err)
		}
		palindrome := input
	*/
	palindrome := "ababbbabbababa"
	P := buildMatrix(palindrome)
	size := len(palindrome)
	for i := 0; i < size; i++ {
		P[i][i] = true
	}
	for j := 0; j < size; j++ {
		for i := 0; i < size-j; i++ {
			if i < i+j-1 {
				P[i][i+j] = (palindrome[i] == palindrome[i+j]) && P[i+1][i+j-1]
			} else if i == i+j {
				P[i][i+j] = true
			} else if i == i+j-1 {
				P[i][i+j] = (palindrome[i] == palindrome[i+j])
			}
		}
	}
	var cuts [15]int
	var backtrack [15]int
	cuts[0] = 0
	backtrack[0] = -1
	fmt.Println(cuts)
	fmt.Println(backtrack)
	p_data := makeTables(P, cuts[:], backtrack[:], size)
	printString(palindrome, p_data, size)
}
