package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	// c := cron.New()
	// c.AddFunc("0 0 1 * *", func() {

	// 1. get list directory to find the oldest with YYYY-MM format
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	if len(entries) == 0 {
		panic("No directory detected")
	}

	var oldestFolder string
	var oldestTime time.Time
	for _, e := range entries {
		strs := strings.Split(e.Name(), "-")
		if len(strs) != 2 {
			continue
		}
		year, err := strconv.Atoi(strs[0])
		if err != nil {
			fmt.Println("1-1. Error : ", err)
			// panic(1)
		}
		month, err := strconv.Atoi(strs[1])
		if err != nil {
			fmt.Println("1-2. Error : ", err)
			// panic(1)
		}

		currentTime := time.Date(year, time.Month(month), 1, 1, 1, 1, 0, time.UTC)
		if oldestTime.IsZero() || currentTime.Before(oldestTime) {
			oldestTime = currentTime
			oldestFolder = fmt.Sprintf("%d-%s", year, strs[1])
		}

	}

	fmt.Println("1. Successfuly get the oldest folder : ", oldestFolder)

	// 2. Zip The folder
	zipFolder := oldestFolder + ".zip"
	cmd := exec.Command("zip", "-r", zipFolder, oldestFolder)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("2. ", err)
	}
	fmt.Println("2. Zip Sucess : ", string(stdout))

	// 3. Send File using scp
	host := "sysadmin@10.254.212.4"
	targetMachine := host + ":/var/www/html/public/photo/survey/"
	password := "R@ngerHi7au*"
	cmd = exec.Command("sshpass", "-p", password, "scp", "-P", "43210", zipFolder, targetMachine)
	stdout, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("3. ", err)
	}
	fmt.Println("3. SCP success : ", string(stdout))

	// 4. Remove oldest  File : ", oldestFolder, " - ", string(stdout))
	cmd = exec.Command("rm", zipFolder)
	stdout, err = cmd.Output()
	if err != nil {
		fmt.Println("4-1. ", err)
	}
	cmd = exec.Command("rm", "-rf", oldestFolder)
	stdout, err = cmd.Output()
	if err != nil {
		fmt.Println("4-2. ", err)
	}
	fmt.Println("4. Remove File ", oldestFolder, " Success")

	// 5. SSH into the target machine, unzip the file, remove the zip file, and list files
	remoteZipFile := "/var/www/html/public/photo/survey/" + zipFolder
	remoteUnzipDir := "/var/www/html/public/photo/survey/"
	sshCommand := fmt.Sprintf("unzip -o %s -d %s && rm %s && ls %s", remoteZipFile, remoteUnzipDir, remoteZipFile, remoteUnzipDir)
	cmd = exec.Command("sshpass", "-p", password, "ssh", "-p", "43210", "sysadmin@10.254.212.4", sshCommand)
	stdout, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("5. ", err)
	}
	fmt.Println("5. Unzip, remove zip file, and list files on remote success : ", string(stdout))

	// })

	// c.Start()

	select {}

}

// sudo zip -r 2023-10.zip 2023-10
// R@ngerHit@m007
