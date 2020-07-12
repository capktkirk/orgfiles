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
	nodeOne := Node{next: nil, prev: nil, data: "Stack1"}
	var nodePtr *Node
	nodePtr = &nodeOne
	newStack := Stack{top: nodePtr}
	stackPtr := &newStack
	pushStack(nodePtr, stackPtr)
	//newStack.top = &nodeOne
	//pushStack(&nodeOne, newStack)
	//nodeTwo := Node{data: "Stack2"}
	//pushStack(&nodeTwo, newStack)

	//4th Queues

	//5th Hash Tables

	//6th Trees

	//7th Maybe Heaps
}
