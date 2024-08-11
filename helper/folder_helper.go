package helper

import (
	"fmt"
	"os"

	"github.com/kirsle/configdir"
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

// Getting directory where the files will be saved, the file structure will be as follows:-
// On Windows: C:/Users/username/AppData/Local/peril-Markdown-Notes
// On macOS: /Users/username/Library/Application Support/peril-Markdown-Notes
// On Linux: /home/username/.local/share/peril-Markdown-Notes
func GetHomeDirectory() (string, error) {
	appDataDir := configdir.LocalConfig("peril-Notes/")
	return appDataDir, nil
}
