package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func main() {
	table := make(map[int]*Node, 10)
	hash := &HashTable{
		Table: table,
		Size:  10,
	}

	fmt.Println("Number of spaces: ", hash.Size)

	for i := 0; i < 90; i++ {
		insert(hash, i)
	}
	traverse(hash)
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(hash *HashTable, value int) int {
	// calculate the hash using the key
	index := hashFunction(value, hash.Size)
	element := Node{
		Value: value,
		Next:  hash.Table[index],
	}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			// get the value which will be a Node(LinkedList)
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)

				// since Next field points to another Node
				t = t.Next
			}
			fmt.Println()

		}
	}
}
