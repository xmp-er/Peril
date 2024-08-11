package auth

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/xmp-er/peril/helper"
)

func DownloadFromGoogleDrive(fileName, localPath string) error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(localPath + fileName)
	isFileAlreadyExists, err := helper.IsFileExists(localPath + fileName)
	if err != nil {
		return fmt.Errorf("error trying to check if file %v already exists", fileName)
	}
	if isFileAlreadyExists {
		fmt.Println("🔴[Alert] File already exists in location, do you want to override local file?(y/n)")
	check:
		for {
			var temp string
			if scanner.Scan() {
				temp = scanner.Text()
			}
			if err := scanner.Err(); err != nil {
				//Log that there was a error reading input
				return fmt.Errorf("error reading input as %v", err)
			}
			switch temp {
			case "n", "No", "no", "nO":
				return fmt.Errorf("error as user has decided to abort the opertaion")
			case "y", "Yes", "yes", "yEs", "YeS", "yeS", "YEs", "yES":
				break check
			default:
				fmt.Println("🔴[Error] Please enter a valid option")
			}
		}
	}
	//Checking if there is a Authentication setup for Google Drive via Rclone
	err = configRCloneAuth()
	if err != nil {
		return err
	}
	fileName_raw := strings.Split(fileName, ".")
	var name string
	var format string
	if len(fileName_raw) >= 2 {
		name = fileName_raw[len(fileName_raw)-2]
		format = fileName_raw[len(fileName_raw)-1]
	} else {
		name = fileName_raw[len(fileName_raw)-1]
	}
	switch format {
	case "enc":
		format = "encoded"
	case "md":
		format = "md"
	default:
		format = "misc"
	}
	remoteFilePath := "peril-Markdown-Notes/" + format + "/" + name + "." + format
	fmt.Println(remoteFilePath)
	fmt.Println(localPath)
	cmd := exec.Command("rclone", "copy", "gdrive:"+remoteFilePath, localPath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to download file from google drive as %v", err)
	}
	return nil
}
