package auth

import (
	"fmt"
	"os/exec"
	"strings"
)

func DownloadFromGoogleDrive(fileName, localPath string) error {
	//Checking if there is a Authentication setup for Google Drive via Rclone
	err := configRCloneAuth()
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
	remoteFilePath := "peril-Markdown-Editor/Note-Taker/" + format + "/" + name
	cmd := exec.Command("rclone", "copy", "gdrive:"+remoteFilePath, " ", localPath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to download file from google drive as %v", err)
	}
	return nil
}
