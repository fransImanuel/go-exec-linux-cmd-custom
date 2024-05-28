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
	fmt.Println("---------------------------Program Started-----------------------------")
	fmt.Printf("----------------------------%v---------------------------\n", time.Now())
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

	// // 4. Remove oldest  File : ", oldestFolder, " - ", string(stdout))
	// cmd = exec.Command("rm", zipFolder)
	// stdout, err = cmd.Output()
	// if err != nil {
	// 	fmt.Println("4-1. ", err)
	// }
	// cmd = exec.Command("rm", "-rf", oldestFolder)
	// stdout, err = cmd.Output()
	// if err != nil {
	// 	fmt.Println("4-2. ", err)
	// }
	// fmt.Println("4. Remove File ", oldestFolder, " Success")
	fmt.Println("4. Remove File ", oldestFolder, " SKIPPED( on comment )")

	// 5. SSH into the target machine, unzip the file, and remove the zip file
	remoteZipFile := "/var/www/html/public/photo/survey/" + zipFolder
	remoteUnzipDir := "/var/www/html/public/photo/survey/"
	sshCommand := fmt.Sprintf("unzip -o %s -d %s && rm %s", remoteZipFile, remoteUnzipDir, remoteZipFile)
	cmd = exec.Command("sshpass", "-p", password, "ssh", "-p", "43210", "sysadmin@10.254.212.4", sshCommand)
	stdout, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("5. ", err)
	}
	fmt.Println("5. Unzip and remove zip file on remote success : ", string(stdout))

	// 6. SSH into the target machine to list files
	sshCommand = fmt.Sprintf("ls %s", remoteUnzipDir)
	cmd = exec.Command("sshpass", "-p", password, "ssh", "-p", "43210", "sysadmin@10.254.212.4", sshCommand)
	stdout, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("6. ", err)
	}
	fmt.Println("6. List files on remote success : ", string(stdout))

	// Process the list of files on the remote machine
	fileList := string(stdout)
	fmt.Println("6. List of files on remote machine: ", fileList)

	// Process the file list
	remoteEntries := strings.Split(fileList, "\n")

	// 7. get list directory to find the oldest with YYYY-MM format in the 2nd VM
	var oldestFolder2ndVM string
	var oldestTime2ndVM time.Time
	for _, e := range remoteEntries {
		if e == "" {
			continue
		}
		strs := strings.Split(e, "-")
		if len(strs) != 2 {
			continue
		}
		year, err := strconv.Atoi(strs[0])
		if err != nil {
			fmt.Println("7-1. Error : ", err)
		}
		month, err := strconv.Atoi(strs[1])
		if err != nil {
			fmt.Println("7-2. Error : ", err)
		}

		currentTime := time.Date(year, time.Month(month), 1, 1, 1, 1, 0, time.UTC)
		if oldestTime2ndVM.IsZero() || currentTime.Before(oldestTime2ndVM) {
			oldestTime2ndVM = currentTime
			oldestFolder2ndVM = fmt.Sprintf("%d-%s", year, strs[1])
		}
	}
	fmt.Println("7. Successfuly get the oldest folder in 2nd VM : ", oldestFolder2ndVM)

	// 8. Remove oldest  File in 2nd VM

	fmt.Println("---------------------------Program Finished-----------------------------")
	fmt.Printf("----------------------------%v---------------------------\n", time.Now())
	// })

	// c.Start()

	select {}

}

// sudo zip -r 2023-10.zip 2023-10
// R@ngerHit@m007
