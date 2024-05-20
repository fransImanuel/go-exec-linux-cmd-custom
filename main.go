package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("sudo", "mkdir", "cnth_zip")
	cmd.Stdin = bytes.NewBufferString("R@ngerHit@m007\n")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(stdout))
}

// sudo zip -r 2023-10.zip 2023-10
// R@ngerHit@m007
