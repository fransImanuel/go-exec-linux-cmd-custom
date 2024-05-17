package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "text")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(stdout))
}
