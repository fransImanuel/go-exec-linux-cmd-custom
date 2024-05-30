package main

import (
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
	fmt.Println("************Program Starting************", time.Now())
	c := cron.New()
	// // setial bulan jam 2 pagi tanggal 5
	textResult := ""
	startTime := time.Now()
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

		fmt.Println("1. Successfuly get the oldest folder in 1st VM: ", oldestFolder)
		textResult += fmt.Sprintf("1. Successfuly get the oldest folder in 1st VM: %s<br>\n", oldestFolder)

		// 2. Zip The folder
		zipFolder := oldestFolder + ".zip"
		cmd := exec.Command("zip", "-r", zipFolder, oldestFolder)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println("2. ", err)
		}
		// fmt.Println("2. Zip Sucess : ", string(stdout))
		fmt.Println("2. Zip Sucess in 1st VM")
		textResult += fmt.Sprintf("2. Zip Sucess in 1st VM<br>\n")

		// 3. Send File using scp
		host := "sysadmin@10.254.212.4"
		targetMachine := host + ":/var/www/html/public/photo/survey/"
		password := "R@ngerHi7au*"
		cmd = exec.Command("sshpass", "-p", password, "scp", "-P", "43210", zipFolder, targetMachine)
		stdout, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("3. ", err)
		}
		fmt.Println("3. SCP ZIP success from 1st VM to 2nd VM: ", string(stdout))
		textResult += fmt.Sprintf("3. SCP ZIP success from 1st VM to 2nd VM: %s<br>\n", string(stdout))

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
		// fmt.Println("4. Remove File ", oldestFolder, " and zipped ", zipFolder, " in current machine  Success")
		fmt.Println("4. Remove File ", oldestFolder, " and zipped ", zipFolder, " in 1st VM SKIPPED( on comment )")
		textResult += fmt.Sprintf("4. Remove File %s and zipped %s in 1st VM SKIPPED( on comment )<br>\n", oldestFolder, zipFolder)

		// 5. SSH into the target machine, unzip the file, and remove the zip file
		remoteZipFile := "/var/www/html/public/photo/survey/" + zipFolder
		remoteUnzipDir := "/var/www/html/public/photo/survey/"
		sshCommand := fmt.Sprintf("unzip -o %s -d %s && rm %s", remoteZipFile, remoteUnzipDir, remoteZipFile)
		cmd = exec.Command("sshpass", "-p", password, "ssh", "-p", "43210", "sysadmin@10.254.212.4", sshCommand)
		stdout, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("5. ", err)
		}
		fmt.Println("5. Unzip and remove zip file on 2nd VM ")
		textResult += fmt.Sprintf("5. Unzip and remove zip file on 2nd VM <br>\n")
		// fmt.Println("5. Unzip and remove zip file on remote success : ", string(stdout))

		// 6. SSH into the target machine to list files
		sshCommand = fmt.Sprintf("ls %s", remoteUnzipDir)
		cmd = exec.Command("sshpass", "-p", password, "ssh", "-p", "43210", "sysadmin@10.254.212.4", sshCommand)
		stdout, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("6. ", err)
		}
		fmt.Println("6. List files on 2nd VM : ", string(stdout))
		textResult += fmt.Sprintf("6. List files on 2nd VM : %s\n<br>", string(stdout))

		// 7. get list directory to find the oldest with YYYY-MM format in the 2nd VM
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
		textResult += fmt.Sprintf("7. Successfuly get the oldest folder in 2nd VM : %s<br>\n", oldestFolder2ndVM)

		// 8. Remove oldest folder in 2nd VM
		remoteOldestFolder := remoteUnzipDir + oldestFolder2ndVM
		sshCommand = fmt.Sprintf("rm -rf %s", remoteOldestFolder)
		fmt.Println("8. command that will be used ", sshCommand)
		textResult += sshCommand + "<br>"

		// cmd = exec.Command("sshpass", "-p", password, "ssh", "-p", "43210", "sysadmin@10.254.212.4", sshCommand)
		// stdout, err = cmd.CombinedOutput()
		// if err != nil {
		// 	fmt.Println("8. ", err)
		// }
		// fmt.Println("8. Remove oldest folder in 2nd VM success : ", oldestFolder2ndVM, string(stdout))
		fmt.Println("8. Remove oldest folder in 2nd VM SKIPPED ( on comment ) : ", oldestFolder2ndVM)
		textResult += fmt.Sprintf("8. Remove oldest folder in 2nd VM SKIPPED ( on comment ) : %s<br>\n", oldestFolder2ndVM)

		finishedTime := time.Now()
		fmt.Printf("\n--------Program Finished at %v--------\n", finishedTime)
		textResult += fmt.Sprintf("\n--------Program Finished at %v--------<br>\n", finishedTime)

		// // 9. Send Email
		smtpConfig := env.GetSMTPConfig()
		smtpClient := mail.InitEmail(smtpConfig)
		Email := []string{"frans.imanuel@visionet.co.id", "lishera.prihatni@visionet.co.id", "ari.darmawan@visionet.co.id", "azky.muhtarom@visionet.co.id"}
		if err := smtpClient.Send(Email, nil, nil, "MetaForce Auto Backup", "text/html", textResult, []string{"program_log.txt"}); err != nil {
			fmt.Println("9. Send Email Error: ", err)
		}
		fmt.Println("9. Send Email Success")

	})

	c.Start()

	select {}

}

// scp -P 43210 go-metaforce-auto-backup-media sysadmin@10.254.212.5:/home/sysadmin/project/metaforce-api/public/photo/survey
// R@ngerKun1ng&

// ./go-metaforce-auto-backup-media > program_log.txt 2>&1 &
