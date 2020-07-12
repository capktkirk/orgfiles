package main

import (
	"fmt"
	//"os"
)

type Node struct {
	next *Node
	prev *Node
	data string
}

type Stack struct {
	top   *Node
	count int
}

func pushStack(n *Node, s *Stack) {
	n.next = s.top
	s.top = n
	s.count++
}

func popStack(s *Stack) *Node {
	temp := s.top
	s.top = temp.next
	s.count--
	return temp
}

func addNode(d string, n *Node) {
	temp := n
	addString := Node{next: nil, prev: nil, data: d}
	for {
		if temp.next != nil {
			temp = temp.next
		} else {
			temp.next = &addString
			temp.next.prev = temp
			return
		}
	}
}

func printList(ll Node) {
	temp := &ll
	for {
		if temp.next != nil {
			fmt.Print(temp.data)
			fmt.Print(" ")
		} else {
			fmt.Print(temp.data)
			fmt.Print("\n")
			return
		}
		temp = temp.next
	}
}

func printListBack(ll Node) {
	temp := &ll
	for {
		if temp.next != nil {
			temp = temp.next
		} else {
			for {
				if temp.prev != nil {
					fmt.Print(temp.data, " ")
					temp = temp.prev
				} else {
					fmt.Print(temp.data, " \n")
					return
				}
			}
		}
		temp = temp.next
	}
}

func arrayTester() {
	var arr [4]string
	arr[0] = "This"
	arr[1] = "Is"
	arr[2] = "An"
	arr[3] = "Array"
	fmt.Println(arr[0], arr[3])
	fmt.Println(arr)
	//Arrays can be set all at once as well.
	fib := [5]int{1, 1, 2, 3, 5}
	fmt.Println(fib[0], fib[1], fib[2], fib[3])
	fmt.Println("String array slicing.")
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	st := "Double"
	//
	fmt.Println("string:", st)
	fmt.Println("st_ln:", len(st))
	//Prints out a single character from a string.
	fmt.Println(string(st[1]))
	//prints out a slice.
	fmt.Println(string(st[2:4]))
	//appending to a slice
	var sl []string
	sl = append(sl, "do")
	sl = append(sl, st[2:4])
	sl = append(sl, "ly")
	fmt.Println(sl)
	//fmt.Println(st[2])
}
