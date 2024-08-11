package auth

import (
	"fmt"
	"os/exec"
	"strings"
)

func verifyRcloneConfig() (bool, error) {
	cmd := exec.Command("rclone", "listremotes")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to list remotes as %v", err)
	}

	remotes := strings.Split(string(output), "\n")
	for _, remote := range remotes {
		if strings.TrimSpace(remote) == "gdrive:" {
			return true, nil
		}
	}
	//Log that drive is already authenticated to work with Rclone
	return false, nil
}
