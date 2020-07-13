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
	//When using maps in Golang, you need to have the variable, and then "make" the map available.
	hashAmt := 2056
	var hashT map[int]string
	hashT = make(map[int]string)
	hashIdx := HashSlot("Dogwood", hashAmt)
	hashT[hashIdx] = "Dogwood"
	fmt.Println(hashT[HashSlot("Dogwood", hashAmt)])

	//Can we auto-assign a map?
	hashT2 := make(map[int]string)
	hashT2[HashSlot("Dilbert", hashAmt)] = "Dilbert"
	fmt.Println(hashT2[HashSlot("Dilbert", hashAmt)])
	//Yes we can.
	//We could make this more thorough by having each map contain a map instead of a string, then doing another hash to get to that for collisions. The chances of a double hash being the same is very unlikely.
	//We can create an anonymous function in the declaration of HashTable, and then fill it in with a specific one when we create the new HashTable. We could actually do this infinitely with changing our map to contain a hashtable instead of a string.
	ht := HashTable{
		hash: make(map[int]string),
		idx:  2056,
		hSlot: func(s string, x int) int {
			tot := 1
			for i := 0; i < len(s); i++ {
				tot = int(s[i]) * tot
			}
			tot = tot % x
			return tot
		},
	}
	ht.hash[ht.hSlot("Dany", ht.idx)] = "Dany"
	fmt.Println(ht.hash[ht.hSlot("Dany", ht.idx)])
	//6th Trees

	//7th Maybe Heaps
}
