package helper

import (
	"fmt"

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
	fmt.Printf("Selected folder: %s\n", selectedFolder)
	return selectedFolder, nil
}
