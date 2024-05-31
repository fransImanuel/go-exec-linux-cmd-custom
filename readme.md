
# Metaforce Autobackup File

This code basically linux shell script that run using golang.

The Idea of this program is to move "certain" foldern to another vm.

in this case, metaforce have 2 vm, 1st vm is saving the last newest 6 month photo folder, and 2nd vm is saving the last oldest 6 month photo folder.

for example, the folder on 1st vm contains
```bash
2023-07 2023-08 2023-09 2023-10 2023-11 2023-12
```

and the folder on 2nd vm contains
```bash
2023-01 2023-02 2023-03 2023-04 2023-05 2023-06
```

in 1st VM: we find oldest folder in there and send it to 2nd VM after that we remove that oldest folder ( in example below in 2023-07 )

in 2nd VM: we accept the oldest folder ( 2023-07 ) through SCP command, unzip it and remove the oldest folder in 2nd VM ( 2023-01 )

the final result will look like this
```bash
2023-08 2023-09 2023-10 2023-11 2023-12 [NEW FOLDER WILL BE GENERATED NAMED 2024-01 LATER]
```

and the folder on 2nd vm contains
```bash
2023-02 2023-03 2023-04 2023-05 2023-06 2023-07
```


## Deployment

Linux command I used to deploy into the server:

Go inside the vm and run (you can use sudo or not depend on your case): 
```bash
  git clone https://github.com/fransImanuel/go-exec-linux-cmd-custom.git
```
repository I'm using is my private git account, just change it to this repo if someday change is needed.


You can change the folder permission so I have access to execute file
```bash
chmod -R 777 go-exec-linux-cmd-custom/
cd go-exec-linux-cmd-custom/
```

Command below is to build the go program
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o go-metaforce-auto-backup-media ./main.go
```

After the code compiled, compiler will generate file named "go-metaforce-auto-backup-media", you can run the code using this command

```bash
./go-metaforce-auto-backup-media -host@ip="sysadmin@10.254.212.4"  -password="R@ngerHi7au*" > program_log.txt 2>&1 &
```
Command above print the output in file named "program_log"

## Optional
Usually before I run 
```bash
./go-metaforce-auto-backup-media -host@ip="sysadmin@10.254.212.4"  -password="R@ngerHi7au*" > program_log.txt 2>&1 &
```

I test the program by creating the temporary folder such as 
```bash
mkdir 2025-01 2025-02 2025-03 2025-04 2025-05 2025-06
```