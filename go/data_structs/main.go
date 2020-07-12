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
	nodeOne = Node{next: nil, prev: nil, data: "Stack2"}
	nodePtr = &nodeOne
	pushStack(nodePtr, stackPtr)
	stackString := popStack(stackPtr).data
	nodeOne = Node{next: nil, prev: nil, data: "Stack3"}
	nodePtr = &nodeOne
	pushStack(nodePtr, stackPtr)
	nodeOne = Node{next: nil, prev: nil, data: "Stack4"}
	nodePtr = &nodeOne
	pushStack(nodePtr, stackPtr)
	nodeOne = Node{next: nil, prev: nil, data: "Stack5"}
	nodePtr = &nodeOne
	pushStack(nodePtr, stackPtr)
	fmt.Println(stackString)
	fmt.Println(newStack.count)
	nodePop := popStack(&newStack).data
	fmt.Println(nodePop)
	nodePop = popStack(&newStack).data
	fmt.Println(nodePop)
	//4th Queues

	//5th Hash Tables

	//6th Trees

	//7th Maybe Heaps
}
