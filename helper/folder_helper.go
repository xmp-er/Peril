package helper

import (
	"fmt"
	"os"

	"github.com/sqweek/dialog"
)

func SelectFolder() (string, error) {
	selectedFolder, err := dialog.Directory().Title("Select a folder").Browse()
	if err != nil {
		if err == dialog.ErrCancelled {
			fmt.Println("Folder selection canceled")
		} else {
			fmt.Printf("Error selecting folder: %v\n", err)
		}
		return "", err
	}
	fmt.Printf("🟢[DONE] Selected folder: %s\n", selectedFolder)
	return selectedFolder, nil
}

func GetCurrentDirectory() (string, error) {
	//Log that we are getting the current directory
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory as %v", err)
	}
	return dir, nil
}
