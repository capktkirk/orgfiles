package main

import (
	"fmt"
	//"os"
)

func main() {
	//1st Arrays
	//Arrays are fixed size in Golang. We can set them up by individually setting them.
	fmt.Println("This is a test of the array systems.")
	arrayTester()
	fmt.Println("--------------------------------------------------------\n")

	//2nd Linked Lists
	fmt.Println("Node testing.")
	newNode := Node{next: nil, prev: nil, data: "First"}
	addNode("second", &newNode)
	addNode("third", &newNode)
	printList(newNode)
	printListBack(newNode)
	addNode("Fourth", &newNode)
	printList(newNode)

	//3rd Stacks
	fmt.Println("--------------------------------------------------------\n")
	fmt.Println("Stack Testing.")
	nodeOne := Node{next: nil, prev: nil, data: "Stack1"}
	var nodePtr *Node
	nodePtr = &nodeOne
	newStack := Stack{top: nodePtr}
	pushStack("Stack1", &newStack)
	pushStack("Stack2", &newStack)
	pushStack("Stack3", &newStack)
	//fmt.Println(popStack(&newStack).data)
	//fmt.Println(popStack(&newStack).data)
	printStack(&newStack)

	//4th Queues
	//Stack is the same as Queue essentially.

	//5th Hash Tables

	//6th Trees

	//7th Maybe Heaps
}
