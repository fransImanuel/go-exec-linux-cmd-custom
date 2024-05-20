package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("sudo", "mkdir", "cnth_zip")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(stdout))
}

// sudo zip -r 2023-10.zip 2023-10
// R@ngerHit@m007
