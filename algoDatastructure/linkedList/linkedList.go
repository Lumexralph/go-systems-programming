/*main - A linkedList is a finite set of elements that occupy
at least two memory locations, one for the value and the
next is a pointer that points or links the current element
to the next element/node in the sequence.

Head - The first node or element in the list
Tail - The last node or element in the list
*/
package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
}


// create an empty Node, allocate memory for it and
// return an address/pointer for it
var root = new(Node)


func main()  {
	fmt.Println(root)
	root = nil
	fmt.Println(root)
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 100)
	addNode(root, 9)
	addNode(root, 8)
	addNode(root, 8)
	traverse(root)
}

func addNode(t *Node, v int) int {

	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	if v == t.Value {

		fmt.Println("Node or value already exists: ", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}


func traverse(t *Node)  {
	if t == nil {
		fmt.Println("-> Empty List!")
		return
	}

	for t != nil {

		fmt.Printf("%d -> ", t.Value)
		t = t.Next

	}
	fmt.Println()
}