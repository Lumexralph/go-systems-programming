/*main - A doublylinkedList is a finite set of elements that occupy
at least three memory locations, one for the value, the previous and the
next is a pointer that points or links the current element
to the previous or next element/node in the sequence.

Head - The first node or element in the list
Tail - The last node or element in the list
*/
package main

import "fmt"

type Node struct {
	Previous *Node
	Value    int
	Next     *Node
}

// create a Node with its zero value
// i.e empty Node and get a pointer to it
var root = new(Node)
var lastNode *Node

func main() {
	fmt.Println(root)
	root = nil
	fmt.Println(root)
	// traverse(root)
	// reverseTraverse(root)
	addNode(root, 1)
	addNode(root, 4)
	// traverse(root)
	addNode(root, 10)
	
	addNode(root, 100)
	traverse(root)
	// addNode(root, 9)
	// addNode(root, 8)
	// addNode(root, 8)
	// traverse(root)
}

func addNode(t *Node, v int) int {

	if root == nil {
		// the first element/node in the sequence
		t = &Node{nil, v, nil}
		root = t
		lastNode = root
		return 0
	}

	if v == t.Value {

		fmt.Println("Node or value already exists: ", v)
		return -1
	}

	if t.Next == nil {
		// means to add the last Node in the sequence
		t.Next = &Node{lastNode, v, nil}
		lastNode = t
		return -2
	}

	if t.Previous == nil {
		t.Previous = &Node{nil, v, t}
		return -3
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

func reverseTraverse(t *Node)  {
	if t == nil {
		fmt.Println("-> Empty List!")
		return
	}

	for t != nil {

		fmt.Printf(" <- %d", t.Value)
		t = t.Previous

	}
	fmt.Println()
}