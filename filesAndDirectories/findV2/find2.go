package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	minusS := flag.Bool("s", false, "Sockets")
	minusP := flag.Bool("p", false, "Pipes")
	minusSL := flag.Bool("sl", false, "Symbolic links")
	minusF := flag.Bool("f", false, "Files")
	minusD := flag.Bool("d", false, "Directories")
	minusX := flag.String("x", "", "Files")

	flag.Parse()
	flags := flag.Args()

	printAll := false
	if *minusS && *minusP && *minusSL && *minusF && *minusD {
		printAll = true
	}

	if !(*minusS || *minusP || *minusSL || *minusF || *minusD) {
		printAll = true
	}

	if len(flags) == 0 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	path := flags[0]

	walkFunction := func(path string, info os.FileInfo, err error) error {
		if excludeName(path, *minusX) {
			return nil
		}

		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if printAll {
			fmt.Println(path)
			return nil
		}

		// get the mode of the path or file
		mode := fileInfo.Mode()
		if mode.IsRegular() && *minusF {
			fmt.Println(path)
			return nil
		}

		if mode.IsDir() && *minusD {
			fmt.Println(path)
			return nil
		}

		fileInfo, _ = os.Lstat(path)
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if *minusSL {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeNamedPipe != 0 {
			if *minusP {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeSocket != 0 {
			if *minusS {
				fmt.Println(path)
				return nil
			}
		}

		return nil

	}

	// walk through the file system of the given path
	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func excludeName(name, exclude string) bool {
	if exclude == "" {
		return false
	}

	if filepath.Base(name) == exclude {
		return true
	}

	return false
}
