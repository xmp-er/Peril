package helper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sqweek/dialog"
)

func OpenOrCreateFile(fileName string, option string) error {
	if _, err := os.Stat(fileName + ".md"); option == "ofile" && os.IsNotExist(err) {
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

func DeleteFile(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		//File does not exist, log it.
		fmt.Printf("File %v not found at current location,aboring operation\n", fileName)
		return nil
	}
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func SelectFile() (string, error) {
	file, err := dialog.File().Title("Select a file to upload").Load()
	if err != nil {
		return "", fmt.Errorf("error selecting the file as %v", err)
	}
	return file, nil
}
