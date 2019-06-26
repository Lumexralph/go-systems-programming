package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// to make sure that you are not going to
	// accidentally execute another binary file
	PS, err := exec.LookPath("ps")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(PS)

	command := []string{"ps", "-a", "-x"}
	env := os.Environ()
	err = syscall.Exec(PS, command, env)
	if err != nil {
		fmt.Println(err)
	}

}
