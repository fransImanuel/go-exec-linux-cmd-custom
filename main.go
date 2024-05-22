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
	fmt.Println("1. Create Dir Success : ", string(stdout))

	// // 1. get list directory to find the oldest with YYYY-MM format
	// entries, err := os.ReadDir("./")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if len(entries) == 0 {
	// 	panic("No directory detected")
	// }

	// var oldestFolder string
	// var oldestTime time.Time
	// for _, e := range entries {
	// 	strs := strings.Split(e.Name(), "-")
	// 	if len(strs) != 2 {
	// 		continue
	// 	}
	// 	year, err := strconv.Atoi(strs[0])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		panic(1)
	// 	}
	// 	month, err := strconv.Atoi(strs[1])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		panic(1)
	// 	}

	// 	currentTime := time.Date(year, time.Month(month), 1, 1, 1, 1, 0, time.UTC)
	// 	if oldestTime.IsZero() || currentTime.Before(oldestTime) {
	// 		oldestTime = currentTime
	// 		oldestFolder = fmt.Sprintf("%d-%s", year, strs[1])
	// 	}

	// }

	// fmt.Println("1. Successfuly get the oldest folder : ", oldestFolder)

	// Zip The folder
	// zipFolder := oldestFolder + ".zip"
	// cmd = exec.Command("zip", "-r", zipFolder, oldestFolder)

	zipFolder := "cnth_zip" + ".zip"
	cmd = exec.Command("zip", "-r", zipFolder, "cnth_zip")
	stdout, err = cmd.Output()
	if err != nil {
		fmt.Println("2. ", err)
	}
	fmt.Println("2. Zip Sucess : ", string(stdout))

	//scp -P 43210 2023-10.zip sysadmin@10.254.212.4:/var/www/html/public/photo/survey/
	// cmd = exec.Command("scp", "-P", "43210", zipFolder, targetMachine)
	// cmd.Stdin = bytes.NewBufferString("R@ngerHi7au*\n")
	// stdout, err = cmd.Output()
	// if err != nil {
	// 	fmt.Println("3. ", err)
	// }

	targetMachine := "sysadmin@10.254.212.4:/var/www/html/public/photo/survey/"
	cmd = exec.Command("scp", "-P", "43210", zipFolder, targetMachine)
	cmdWriter, err := cmd.StdinPipe()

	err = cmd.Start()
	if err != nil {
		fmt.Println("3-1. ", err)
	}
	n, err := cmdWriter.Write([]byte("R@ngerHi7au*"))
	if err != nil {
		fmt.Println("3-2. ", err)
	}
	fmt.Println("3-3. ", n)
	err = cmd.Wait()
	if err != nil {
		fmt.Println("3-4. ", err)
	}
	fmt.Println("3. SCP success : ", string(stdout))
}

// sudo zip -r 2023-10.zip 2023-10
// R@ngerHit@m007
