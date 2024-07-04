package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// fmt.Println("************Program Starting************", time.Now())

	// //Check SSH credentials
	// checkCmd := exec.Command("sshpass", "-p", *password, "ssh", "-p", "43210", "-o", "StrictHostKeyChecking=no", *host, "echo", "SSH connection successful")
	// checkOutput, err := checkCmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("SSH authentication failed: ", err)
	// 	fmt.Println(string(checkOutput))
	// 	panic(1)
	// }
	// fmt.Println("SSH authentication successful: ", string(checkOutput))

	// // // setial bulan jam 2 pagi tanggal 5
	textResult := ""
	// startTime := time.Now()
	// c := cron.New()
	// c.AddFunc("0 2 5 * *", func() {

	// 	// 1. get list directory to find the oldest with YYYY-MM format
	// 	fmt.Printf("\n--------Program Started at %v--------\n", startTime)
	// 	textResult += fmt.Sprintf("\n--------Program Started at %v--------<br>\n", startTime)
	// 	entries, err := os.ReadDir("./")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if len(entries) == 0 {
	// 		panic("No directory detected")
	// 	}

	// 	var oldestFolder string
	// 	var oldestTime time.Time
	// 	fmt.Printf("1. List Folder in 1st VM : ")
	// 	textResult += fmt.Sprintf("1. List Folder in 1st VM : ")
	// 	for _, e := range entries {
	// 		fmt.Printf("%s ", e.Name())
	// 		textResult += fmt.Sprintf("%s ", e.Name())
	// 		strs := strings.Split(e.Name(), "-")
	// 		if len(strs) != 2 {
	// 			continue
	// 		}
	// 		year, err := strconv.Atoi(strs[0])
	// 		if err != nil {
	// 			fmt.Println("1-1. Error : ", err)
	// 			panic(1)
	// 		}
	// 		month, err := strconv.Atoi(strs[1])
	// 		if err != nil {
	// 			fmt.Println("1-2. Error : ", err)
	// 			panic(1)
	// 		}

	// 		currentTime := time.Date(year, time.Month(month), 1, 1, 1, 1, 0, time.UTC)
	// 		if oldestTime.IsZero() || currentTime.Before(oldestTime) {
	// 			oldestTime = currentTime
	// 			oldestFolder = fmt.Sprintf("%d-%s", year, strs[1])
	// 		}
	// 	}
	// 	textResult += "<br>"
	// 	fmt.Println()

	// 	fmt.Println("2. Successfuly get the oldest folder in 1st VM: ", oldestFolder)
	// 	textResult += fmt.Sprintf("2. Successfuly get the oldest folder in 1st VM: %s<br>\n", oldestFolder)

	// 	// 3. Zip The folder
	// 	zipFolder := oldestFolder + ".zip"
	// 	cmd := exec.Command("zip", "-r", zipFolder, oldestFolder)
	// 	stdout, err := cmd.Output()
	// 	if err != nil {
	// 		fmt.Println("3. ", err)
	// 		panic(1)
	// 	}
	// 	// fmt.Println("3. Zip Sucess : ", string(stdout))
	// 	fmt.Println("3. Zip Sucess in 1st VM")
	// 	textResult += fmt.Sprintf("3. Zip Sucess in 1st VM<br>\n")

	// 	// 4. Send File using scp
	// 	targetMachine := *host + ":/var/www/html/public/photo/survey/"
	// 	cmd = exec.Command("sshpass", "-p", *password, "scp", "-P", "43210", zipFolder, targetMachine)
	// 	stdout, err = cmd.CombinedOutput()
	// 	if err != nil {
	// 		fmt.Println("4. ", err)
	// 		panic(1)
	// 	}
	// 	fmt.Println("4. SCP ZIP success from 1st VM to 2nd VM: ", string(stdout))
	// 	textResult += fmt.Sprintf("4. SCP ZIP success from 1st VM to 2nd VM: %s<br>\n", string(stdout))

	// 	// 5. Remove oldest  File : ", oldestFolder, " - ", string(stdout))
	// 	cmd = exec.Command("rm", zipFolder)
	// 	stdout, err = cmd.Output()
	// 	if err != nil {
	// 		fmt.Println("5-1. ", err)
	// 	}
	// 	cmd = exec.Command("rm", "-rf", oldestFolder)
	// 	stdout, err = cmd.Output()
	// 	if err != nil {
	// 		fmt.Println("5-2. ", err)
	// 	}
	// 	fmt.Println("5. Remove File ", oldestFolder, " and zipped ", zipFolder, " in current machine  Success")
	// 	// fmt.Println("5. Remove File ", oldestFolder, " and zipped ", zipFolder, " in 1st VM SKIPPED( on comment )")
	// 	// textResult += fmt.Sprintf("5. Remove File %s and zipped %s in 1st VM SKIPPED( on comment )<br>\n", oldestFolder, zipFolder)
	// 	textResult += fmt.Sprintf("5. Remove File %s and zipped %s in 1st VM <br>\n", oldestFolder, zipFolder)

	// 	// 6. SSH into the target machine, unzip the file, and remove the zip file
	// 	remoteZipFile := "/var/www/html/public/photo/survey/" + zipFolder
	// 	remoteUnzipDir := "/var/www/html/public/photo/survey/"
	// 	sshCommand := fmt.Sprintf("unzip -o %s -d %s && rm %s", remoteZipFile, remoteUnzipDir, remoteZipFile)
	// 	cmd = exec.Command("sshpass", "-p", *password, "ssh", "-p", "43210", *host, sshCommand)
	// 	stdout, err = cmd.CombinedOutput()
	// 	if err != nil {
	// 		fmt.Println("6. ", err)
	// 		panic(1)
	// 	}
	// 	fmt.Println("6. Unzip and remove zip file on 2nd VM ")
	// 	textResult += fmt.Sprintf("6. Unzip and remove zip file on 2nd VM <br>\n")
	// 	// fmt.Println("6. Unzip and remove zip file on remote success : ", string(stdout))

	// 	// 7. SSH into the target machine to list files
	// 	sshCommand = fmt.Sprintf("ls %s", remoteUnzipDir)
	// 	cmd = exec.Command("sshpass", "-p", *password, "ssh", "-p", "43210", *host, sshCommand)
	// 	stdout, err = cmd.CombinedOutput()
	// 	if err != nil {
	// 		fmt.Println("7. ", err)
	// 		panic(1)
	// 	}
	// 	fmt.Println("7. List files on 2nd VM : ", string(stdout))
	// 	textResult += fmt.Sprintf("7. List files on 2nd VM : %s\n<br>", string(stdout))

	// 	// 8. get list directory to find the oldest with YYYY-MM format in the 2nd VM
	// 	// Process the list of files on the remote machine
	// 	fileList := string(stdout)
	// 	// Process the file list
	// 	remoteEntries := strings.Split(fileList, "\n")
	// 	var oldestFolder2ndVM string
	// 	var oldestTime2ndVM time.Time
	// 	for _, e := range remoteEntries {
	// 		if e == "" {
	// 			continue
	// 		}
	// 		strs := strings.Split(e, "-")
	// 		if len(strs) != 2 {
	// 			continue
	// 		}
	// 		year, err := strconv.Atoi(strs[0])
	// 		if err != nil {
	// 			fmt.Println("8-1. Error : ", err)
	// 			panic(1)
	// 		}
	// 		month, err := strconv.Atoi(strs[1])
	// 		if err != nil {
	// 			fmt.Println("8-2. Error : ", err)
	// 			panic(1)
	// 		}

	// 		currentTime := time.Date(year, time.Month(month), 1, 1, 1, 1, 0, time.UTC)
	// 		if oldestTime2ndVM.IsZero() || currentTime.Before(oldestTime2ndVM) {
	// 			oldestTime2ndVM = currentTime
	// 			oldestFolder2ndVM = fmt.Sprintf("%d-%s", year, strs[1])
	// 		}
	// 	}
	// 	fmt.Println("8. Successfuly get the oldest folder in 2nd VM : ", oldestFolder2ndVM)
	// 	textResult += fmt.Sprintf("8. Successfuly get the oldest folder in 2nd VM : %s<br>\n", oldestFolder2ndVM)

	// 	// 9. Remove oldest folder in 2nd VM
	// 	remoteOldestFolder := remoteUnzipDir + oldestFolder2ndVM
	// 	sshCommand = fmt.Sprintf("rm -rf %s", remoteOldestFolder)
	// 	fmt.Println("9. command that will be used ", sshCommand)
	// 	textResult += sshCommand + "<br>"

	// 	cmd = exec.Command("sshpass", "-p", *password, "ssh", "-p", "43210", *host, sshCommand)
	// 	stdout, err = cmd.CombinedOutput()
	// 	if err != nil {
	// 		fmt.Println("9. ", err)
	// 	}
	// 	fmt.Println("9. Remove oldest folder in 2nd VM success : ", oldestFolder2ndVM, string(stdout))
	// 	// fmt.Println("9. Remove oldest folder in 2nd VM SKIPPED ( on comment ) : ", oldestFolder2ndVM)
	// 	// textResult += fmt.Sprintf("9. Remove oldest folder in 2nd VM SKIPPED ( on comment ) : %s<br>\n", oldestFolder2ndVM)
	// 	textResult += fmt.Sprintf("9. Remove oldest folder in 2nd VM success : %s<br>\n", oldestFolder2ndVM)

	// 	finishedTime := time.Now()
	// 	fmt.Printf("\n--------Program Backup Photo Finished at %v--------\n", finishedTime)
	// 	textResult += fmt.Sprintf("\n--------Program Backup Photo Finished at %v--------<br>\n", finishedTime)

	//====================================MONGODB================================================

	textResult += "------------------------------------------------------------------------------<br>"
	textResult += "<br><b>Continue Backup TaskList from MongoDB</b><br>"
	textResult += "------------------------------------------------------------------------------<br>"
	// MongoDB connection URI
	// uri := "mongodb://goodtime:HujanAir!2024!@10.10.86.142:27017/?authSource=admin"
	uri := "mongodb://mongoAdmin:&Mer4h&Mud4&@10.254.213.3:4949/?authSource=admin"
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		textResult += "10. Error backup tasklist mongodb : " + err.Error()
	}
	textResult += "<br>10. Success Connect to mongodb<br>"
	defer client.Disconnect(context.TODO())

	// Select database and collection
	db := client.Database("metaforce_prod")
	collection := db.Collection("tr_tasklists")

	// Define the query

	currentTime := time.Now()
	previousMonth := currentTime.AddDate(0, -1, 0)
	yearNow, monthNow, _ := currentTime.Date()
	yearPrev, monthPrev, _ := previousMonth.Date()
	startDate := time.Date(yearPrev, monthPrev, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(yearNow, monthNow, 1, 0, 0, 0, 0, time.UTC)
	filter := bson.D{
		{"ScheduleVisit",
			bson.D{
				{"$gte", startDate},
				{"$lt", endDate},
			},
		},
	}

	fmt.Println("filter:", filter)
	// panic(1)

	// Find documents
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		textResult += "<br>10. Error backup tasklist mongodb : " + err.Error()
	}
	defer cursor.Close(context.TODO())

	// Open output file
	mongoDBBackupFormatName := "MongoD_Tasklist_Backup_"
	mongoDBBackupFullName := mongoDBBackupFormatName + startDate.Format("2006-01-02") + ".json"
	FirstVMTargetDir := "/home/sysadmin/project/metaforce-api/public/photo/survey/"
	file, err := os.Create(FirstVMTargetDir + mongoDBBackupFullName)
	if err != nil {
		log.Fatal(err)
		textResult += "10. Error backup tasklist mongodb : " + err.Error()
	}
	defer file.Close()

	// Write documents to file
	encoder := json.NewEncoder(file)
	for cursor.Next(context.TODO()) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			log.Fatal(err)
			textResult += "10. Error backup tasklist mongodb : " + err.Error()
		}
		if err := encoder.Encode(document); err != nil {
			log.Fatal(err)
			textResult += "10. Error backup tasklist mongodb : " + err.Error()
		}
		// fmt.Println(cursor.Current)

		// var document bson.M
		// if err := cursor.Decode(&document); err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(document)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		textResult += "10. Error backup tasklist mongodb : " + err.Error()
	}

	fmt.Println("Export completed successfully!")

	textResult += "<br>10. Successfully extract tasklist from mongodb "

	entriesMongoDB, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	//check if there is 6 entries of mongodb backup, send the oldest one to 2nd VM
	// maxTotalMongoDBBackup := 6
	countMongoDBBackup := 0
	for _, e := range entriesMongoDB {
		fmt.Printf("%s ", e.Name())
		if strings.Contains(e.Name(), mongoDBBackupFormatName) {
			countMongoDBBackup++
		}
	}

	// if maxTotalMongoDBBackup == countMongoDBBackup {

	textResult += "<br>[Maximal save in this vm is 6, exceeding 6, processing to send to 2nd VM]"
	zipFolder := mongoDBBackupFormatName + ".zip"
	cmd := exec.Command("zip", "-r", zipFolder, mongoDBBackupFormatName)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("[*]", err.Error(), " : ", string(stdout))
	}

	// [*]. Send File using scp
	targetMachine := *host + ":/var/www/html/public/photo/survey/mongodb_tasklist_backup"
	cmd = exec.Command("sshpass", "-p", *password, "scp", "-P", "43210", zipFolder, targetMachine)
	stdout, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[*]", err)
		panic(1)
	}
	fmt.Println("[*]SCP ZIP success from 1st VM to 2nd VM: ", string(stdout))
	textResult += fmt.Sprintf("[*]SCP ZIP success from 1st VM to 2nd VM: %s<br>\n", string(stdout))

	// }
	// ====================================MONGODB================================================

	// 	// // 11. Send Email
	// 	smtpConfig := env.GetSMTPConfig()
	// 	smtpClient := mail.InitEmail(smtpConfig)
	// 	Email := []string{"frans.imanuel@visionet.co.id", "lishera.prihatni@visionet.co.id", "ari.darmawan@visionet.co.id", "azky.muhtarom@visionet.co.id"}
	// 	if err := smtpClient.Send(Email, nil, nil, "MetaForce Auto Backup", "text/html", textResult, []string{"program_log.txt"}); err != nil {
	// 		fmt.Println("11. Send Email Error: ", err)
	// 		panic(1)
	// 	}
	// 	fmt.Println("11. Send Email Success")

	// })
	// c.Start()
	// select {}

}
