package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func OpenOrCreateFile(fileName string, current_folder_location string) error {
	fileName = current_folder_location + "/" + fileName
	if _, err := os.Stat(fileName + ".md"); os.IsNotExist(err) {
		//File does not exist, log it.
		fmt.Printf("File %v.md not found at current location,creating it\n", fileName)
	}
	cmd := exec.Command("vim", fileName+".md")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(fileName, current_folder_location string) error {
	fileName = current_folder_location + "/" + fileName
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		//File does not exist, log it.
		return err
	}
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func IsDirectoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}

func IsFileExists(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return !info.IsDir(), nil
}
