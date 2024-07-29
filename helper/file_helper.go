package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func OpenOrCreateFile(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		//File does not exist, log it.
		fmt.Printf("File %v.md not found at current location,creating it", fileName)
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
