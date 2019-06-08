package main

import (
	"path/filepath"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument !")
		os.Exit(1)
	}

	filename := arguments[1]

	fileinfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileinfo.Mode() & os.ModeSymlink != 0 {
		fmt.Println(filename, " is a symbolic link")

		// process the symblinks to know the file it points to
		realPath, err := filepath.EvalSymlinks(filename)
		if err != nil {
			fmt.Println("Path in symlink can not be resolved")
			os.Exit(1)
		}
		fmt.Println("Path :", realPath)
	}
}
