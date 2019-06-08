package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	dir, err := os.Getwd()
	if err == nil {
		fmt.Println(dir)
	} else {
		fmt.Println("Error: ", err)
	}

	if len(arguments) == 1 {
		return
	}
	if arguments[1] != "-P" {
		return
	}

	// Get the file info
	fileinfo, err := os.Lstat(dir)
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		// get the symlink bit and get the realpath
		// from the symlink
		realPath, err := filepath.EvalSymlinks(dir)
		if err != nil {
			fmt.Println("Path in symlink can not be resolved")
			os.Exit(1)
		}
		fmt.Println("Path :", realPath)
	}
}
