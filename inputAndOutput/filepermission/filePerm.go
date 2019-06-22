package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func tripletToBinary(triplet string) string {
	switch triplet {
	case "rwx":
		return "111"
	case "-wx":
		return "011"
	case "--x":
		return "001"
	case "---":
		return "000"
	case "r-x":
		return "101"
	case "r--":
		return "100"
	case "rw-":
		return "110"
	case "-w-":
		return "010"
	default:
		return "unknown"
	}
}

func convertToBinary(permissions string) string {
	binaryPermission := permissions[1:]

	p1 := binaryPermission[0:3]
	p2 := binaryPermission[3:6]
	p3 := binaryPermission[6:9]

	return tripletToBinary(p1) + tripletToBinary(p2) + tripletToBinary(p3)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("usage: %s filename \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	filename := arguments[1]
	info, _ := os.Stat(filename)
	mode := info.Mode()

	fmt.Println(filename, " mode is ", mode)
	fmt.Println("As string is ", mode.String()[1:10])
	fmt.Println("As binary is ", convertToBinary(mode.String()))
}
