package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("mkdir", "cnth_zip")
	// cmd.Stdin = bytes.NewBufferString("R@ngerHit@m007\n")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("1. ", err)
	}
	fmt.Println("1.", string(stdout))

	cmd = exec.Command("zip", "-r", "cnth_zipped.zip", "cnth_zip")
	stdout, err = cmd.Output()
	if err != nil {
		fmt.Println("2. ", err)
	}

	fmt.Println("2.", string(stdout))
}

// sudo zip -r 2023-10.zip 2023-10
// R@ngerHit@m007
