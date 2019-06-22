package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	// permissions := permissions[1:]

	p1 := permissions[0:3]
	p2 := permissions[3:6]
	p3 := permissions[6:9]

	p1 = tripletToBinary(p1)
	p2 = tripletToBinary(p2)
	p3 = tripletToBinary(p3)

	p1Int, _ := strconv.ParseInt(p1, 2, 64)
	p2Int, _ := strconv.ParseInt(p2, 2, 64)
	p3Int, _ := strconv.ParseInt(p3, 2, 64)

	returnValue := p1Int*100 + p2Int*10 + p3Int
	tempReturnValue := int(returnValue)
	returnString := "0" + strconv.Itoa(tempReturnValue)

	return returnString
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Printf("usage: %s filename permission \n",
			filepath.Base(arguments[0]))
		os.Exit(1)
	}

	filename, _ := filepath.EvalSymlinks(arguments[1])
	permissions := arguments[2]
	if len(permissions) != 9 {
		fmt.Println("Permissions should be 9 characters (rwxrwxrwx): ",
		 permissions)
		 os.Exit(-1)
	}

	bin := convertToBinary(permissions)
	newPerm, _ := strconv.ParseUint(bin, 0, 32)

	// convert to the mode type for file
	newMode := os.FileMode(newPerm)
	os.Chmod(filename, newMode)

}
