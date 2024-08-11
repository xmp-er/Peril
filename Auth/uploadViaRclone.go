package auth

import (
	"fmt"
	"os/exec"
	"strings"
)

func UploadToGoogleDrive(localFilePath, current_folder_location string) error {
	err := configRCloneAuth()
	if err != nil {
		return err
	}
	//Getting file name
	fileName_raw := strings.Split(localFilePath, "/")
	fileName := fileName_raw[len(fileName_raw)-1]
	//Getting file format
	format_raw := strings.Split(fileName, ".")
	format := format_raw[len(format_raw)-1]
	switch format {
	case "enc":
		format = "encoded"
	case "md":
		format = "md"
	default:
		format = "misc"
	}
	remotePath := "peril-Markdown-Notes/" + format
	cmd := exec.Command("rclone", "copy", current_folder_location+localFilePath, "gdrive:"+remotePath)
	err = cmd.Run()
	if err != nil {
		//Log that we failed to upload file
		return fmt.Errorf("failed to upload file: %v", err)
	}
	//Log that file was uploaded
	return nil
}

func configRCloneAuth() error {
	isAuthConfigured, err := verifyRcloneConfig()
	if err != nil {
		return err
	}
	if isAuthConfigured {
		return nil
	}
	//Log that we are configuring Auth for Google drive
	cmd := exec.Command("rclone", "config", "create", "gdrive", "drive", "scope=drive")
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to create rclone config as %v", err)
	}
	//Log that configuration was created for uploading to google drive
	return nil
}
