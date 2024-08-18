package helper

import (
	"github.com/kirsle/configdir"
)

// Getting directory where the files will be saved, the file structure will be as follows:-
// On Windows: C:/Users/username/AppData/Local/peril-Markdown-Notes
// On macOS: /Users/username/Library/Application Support/peril-Markdown-Notes
// On Linux: /home/username/.local/share/peril-Markdown-Notes
func GetHomeDirectory() (string, error) {
	appDataDir := configdir.LocalConfig("peril-Notes/")
	return appDataDir, nil
}
