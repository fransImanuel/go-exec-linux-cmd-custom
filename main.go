package main

import (
	"flag"
	"fmt"
	"go-exec-linux-cmd-custom/util/env"
	"go-exec-linux-cmd-custom/util/mail"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	host := flag.String("host@ip", "", "example: sysadmin@192.168.1.1")
	password := flag.String("password", "", "password")
	// Parse the flags
	flag.Parse()
	if *host == "" {
		panic("need to define host@ip flag")
	}
	if *password == "" {
		panic("need to define password flag")
	}
	fmt.Println("Registered host is "+*host+" and password is ", *password)
	// panic(1)
	fmt.Println("************Program Starting************", time.Now())

	// // setial bulan jam 2 pagi tanggal 5
	textResult := ""
	startTime := time.Now()
	c := cron.New()
	c.AddFunc("0 2 5 * *", func() {

		// 1. get list directory to find the oldest with YYYY-MM format
		fmt.Printf("\n--------Program Started at %v--------\n", startTime)
		textResult += fmt.Sprintf("\n--------Program Started at %v--------<br>\n", startTime)
		entries, err := os.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}
		if len(entries) == 0 {
			panic("No directory detected")
		}

		var oldestFolder string
		var oldestTime time.Time
		fmt.Printf("1. List Folder in 1st VM : ")
		textResult += fmt.Sprintf("1. List Folder in 1st VM : ")
		for _, e := range entries {
			fmt.Printf("%s ", e.Name())
			textResult += fmt.Sprintf(e.Name(), " ")
			strs := strings.Split(e.Name(), "-")
			if len(strs) != 2 {
				continue
			}
			year, err := strconv.Atoi(strs[0])
			if err != nil {
				fmt.Println("1-1. Error : ", err)
				panic(1)
			}
			month, err := strconv.Atoi(strs[1])
			if err != nil {
				fmt.Println("1-2. Error : ", err)
				panic(1)
			}

			currentTime := time.Date(year, time.Month(month), 1, 1, 1, 1, 0, time.UTC)
			if oldestTime.IsZero() || currentTime.Before(oldestTime) {
				oldestTime = currentTime
				oldestFolder = fmt.Sprintf("%d-%s", year, strs[1])
			}
		}
		textResult += "<br>"
		fmt.Println()

		fmt.Println("2. Successfuly get the oldest folder in 1st VM: ", oldestFolder)
		textResult += fmt.Sprintf("2. Successfuly get the oldest folder in 1st VM: %s<br>\n", oldestFolder)

		// 3. Zip The folder
		zipFolder := oldestFolder + ".zip"
		cmd := exec.Command("zip", "-r", zipFolder, oldestFolder)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println("3. ", err)
			panic(1)
		}
		// fmt.Println("3. Zip Sucess : ", string(stdout))
		fmt.Println("3. Zip Sucess in 1st VM")
		textResult += fmt.Sprintf("3. Zip Sucess in 1st VM<br>\n")

		// 4. Send File using scp
		targetMachine := *host + ":/var/www/html/public/photo/survey/"
		cmd = exec.Command("sshpass", "-p", *password, "scp", "-P", "43210", zipFolder, targetMachine)
		stdout, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("4. ", err)
			panic(1)
		}
		fmt.Println("4. SCP ZIP success from 1st VM to 2nd VM: ", string(stdout))
		textResult += fmt.Sprintf("4. SCP ZIP success from 1st VM to 2nd VM: %s<br>\n", string(stdout))

		// // 5. Remove oldest  File : ", oldestFolder, " - ", string(stdout))
		// cmd = exec.Command("rm", zipFolder)
		// stdout, err = cmd.Output()
		// if err != nil {
		// 	fmt.Println("5-1. ", err)
		// }
		// cmd = exec.Command("rm", "-rf", oldestFolder)
		// stdout, err = cmd.Output()
		// if err != nil {
		// 	fmt.Println("5-2. ", err)
		// }
		// fmt.Println("5. Remove File ", oldestFolder, " and zipped ", zipFolder, " in current machine  Success")
		fmt.Println("5. Remove File ", oldestFolder, " and zipped ", zipFolder, " in 1st VM SKIPPED( on comment )")
		textResult += fmt.Sprintf("5. Remove File %s and zipped %s in 1st VM SKIPPED( on comment )<br>\n", oldestFolder, zipFolder)

		// 6. SSH into the target machine, unzip the file, and remove the zip file
		remoteZipFile := "/var/www/html/public/photo/survey/" + zipFolder
		remoteUnzipDir := "/var/www/html/public/photo/survey/"
		sshCommand := fmt.Sprintf("unzip -o %s -d %s && rm %s", remoteZipFile, remoteUnzipDir, remoteZipFile)
		cmd = exec.Command("sshpass", "-p", *password, "ssh", "-p", "43210", *host, sshCommand)
		stdout, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("6. ", err)
			panic(1)
		}
		fmt.Println("6. Unzip and remove zip file on 2nd VM ")
		textResult += fmt.Sprintf("6. Unzip and remove zip file on 2nd VM <br>\n")
		// fmt.Println("6. Unzip and remove zip file on remote success : ", string(stdout))

		// 7. SSH into the target machine to list files
		sshCommand = fmt.Sprintf("ls %s", remoteUnzipDir)
		cmd = exec.Command("sshpass", "-p", *password, "ssh", "-p", "43210", *host, sshCommand)
		stdout, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("7. ", err)
			panic(1)
		}
		fmt.Println("7. List files on 2nd VM : ", string(stdout))
		textResult += fmt.Sprintf("7. List files on 2nd VM : %s\n<br>", string(stdout))

		// 8. get list directory to find the oldest with YYYY-MM format in the 2nd VM
		// Process the list of files on the remote machine
		fileList := string(stdout)
		// Process the file list
		remoteEntries := strings.Split(fileList, "\n")
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
				fmt.Println("8-1. Error : ", err)
				panic(1)
			}
			month, err := strconv.Atoi(strs[1])
			if err != nil {
				fmt.Println("8-2. Error : ", err)
				panic(1)
			}

			currentTime := time.Date(year, time.Month(month), 1, 1, 1, 1, 0, time.UTC)
			if oldestTime2ndVM.IsZero() || currentTime.Before(oldestTime2ndVM) {
				oldestTime2ndVM = currentTime
				oldestFolder2ndVM = fmt.Sprintf("%d-%s", year, strs[1])
			}
		}
		fmt.Println("8. Successfuly get the oldest folder in 2nd VM : ", oldestFolder2ndVM)
		textResult += fmt.Sprintf("8. Successfuly get the oldest folder in 2nd VM : %s<br>\n", oldestFolder2ndVM)

		// 9. Remove oldest folder in 2nd VM
		remoteOldestFolder := remoteUnzipDir + oldestFolder2ndVM
		sshCommand = fmt.Sprintf("rm -rf %s", remoteOldestFolder)
		fmt.Println("9. command that will be used ", sshCommand)
		textResult += sshCommand + "<br>"

		// cmd = exec.Command("sshpass", "-p", password, "ssh", "-p", "43210", "sysadmin@10.254.212.4", sshCommand)
		// stdout, err = cmd.CombinedOutput()
		// if err != nil {
		// 	fmt.Println("9. ", err)
		// }
		// fmt.Println("9. Remove oldest folder in 2nd VM success : ", oldestFolder2ndVM, string(stdout))
		fmt.Println("9. Remove oldest folder in 2nd VM SKIPPED ( on comment ) : ", oldestFolder2ndVM)
		textResult += fmt.Sprintf("9. Remove oldest folder in 2nd VM SKIPPED ( on comment ) : %s<br>\n", oldestFolder2ndVM)

		finishedTime := time.Now()
		fmt.Printf("\n--------Program Finished at %v--------\n", finishedTime)
		textResult += fmt.Sprintf("\n--------Program Finished at %v--------<br>\n", finishedTime)

		// // 9. Send Email
		smtpConfig := env.GetSMTPConfig()
		smtpClient := mail.InitEmail(smtpConfig)
		Email := []string{"frans.imanuel@visionet.co.id", "lishera.prihatni@visionet.co.id", "ari.darmawan@visionet.co.id", "azky.muhtarom@visionet.co.id"}
		if err := smtpClient.Send(Email, nil, nil, "MetaForce Auto Backup", "text/html", textResult, []string{"program_log.txt"}); err != nil {
			fmt.Println("9. Send Email Error: ", err)
			panic(1)
		}
		fmt.Println("9. Send Email Success")

	})
	c.Start()
	select {}

}
