package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	cmd := exec.Command("mkdir", "cnth_zip")
	// cmd.Stdin = bytes.NewBufferString("R@ngerHit@m007\n")
	// stdout, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("1. ", err)
	// }
	// fmt.Println("1. Create Dir Success : ", string(stdout))

	// 1. get list directory to find the oldest with YYYY-MM format
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	if len(entries) == 0 {
		panic("No directory detected")
	}

	var SmallestYear, HighestMonth int
	for _, e := range entries {
		fmt.Println(e.Name())
		// tahun paling kecil, bulan paling gede
		strs := strings.Split(e.Name(), "-")
		year, err := strconv.Atoi(strs[0])
		if err != nil {
			fmt.Println(err)
			panic(1)
		}
		month, err := strconv.Atoi(strs[1])
		if err != nil {
			fmt.Println(err)
			panic(1)
		}

		if SmallestYear == 0 || SmallestYear > year {
			SmallestYear = year
		}
		if HighestMonth == 0 || HighestMonth < month {
			HighestMonth = month
		}
	}

	oldestFolder := fmt.Sprintf("%d-%d", SmallestYear, HighestMonth)
	fmt.Println("1. Successfuly get the oldest folder : ", oldestFolder)

	// Zip The folder
	cmd = exec.Command("zip", "-r", oldestFolder+".zip", oldestFolder)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("2. ", err)
	}
	fmt.Println("2. Zip Sucess : ", string(stdout))

}

// sudo zip -r 2023-10.zip 2023-10
// R@ngerHit@m007
