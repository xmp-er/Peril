package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	auth "github.com/xmp-er/peril/Auth"
	"github.com/xmp-er/peril/helper"
)

var options string = "\n⚪️ open <name> [Create/Open file]\n⚪️ efile <name> [Encrypt file]\n⚪️ dfile <name> [Decrypt File]\n⚪️ del <name_with_extension> [Delete file]\n⚪️ up <Drive_Path> [Upload a file to Google Drive]\n⚪️ down <Drive_Path> [Download a file from Google Drive]"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Peril, what would you like to do?" + options)
	current_folder_location, err := helper.GetHomeDirectory()
	if err != nil {
		fmt.Println("🔴[ERROR] Error unable to get home directory, aborting application")
		os.Exit(1)
	}

	isAppDirectoryExists, err := helper.IsDirectoryExists(current_folder_location)
	if err != nil {
		fmt.Println("🔴[ERROR] Error unable to check if application folder exists, aborting application")
		os.Exit(1)
	}
	if !isAppDirectoryExists {
		cmd := exec.Command("mkdir", current_folder_location)
		err = cmd.Run()
		if err != nil {
			fmt.Println("🔴[ERROR] Error unable to create the desired directory, aborting application")
			os.Exit(1)
		}
	}
	cmd := exec.Command("cd", current_folder_location)
	err = cmd.Run()
	if err != nil {
		fmt.Println("🔴[ERROR] Error unable to move to desired directory, aborting application")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("🔴[ERROR] Error getting home folder location, shutting application down")
		os.Exit(1)
	}
	for {
		var e string
		if scanner.Scan() {
			e = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			//Log that there was a error reading input
			fmt.Println("🔴[ERROR] Error reading input:", err)
		}
		e_e := strings.Split(e, " ")
		switch e_e[0] {
		case "open":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println(`🔴[ERROR] Please add the name of the file to create/open as "create <name>"`)
				break
			}
			err := helper.OpenOrCreateFile(e_e[1], current_folder_location)
			if err != nil {
				msg := fmt.Sprintf("🔴[ERROR] %v", err)
				fmt.Println(msg)
			}
		case "efile":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("🔴[ERROR] Please add the name of the file to encrypt as efile <name> || 5 <name>")
				break
			}
			pass, err := helper.EncryptAndDeleteOriginal(e_e[1], current_folder_location)
			if err != nil {
				fmt.Println(err)
				break
			}
			finalRes := fmt.Sprintf("🟢[DONE] Password for decrypting %v is %v", (e_e[1] + ".enc"), pass)
			fmt.Println(finalRes)
		case "dfile":
			if len(e_e) < 3 {
				fmt.Println("🔴[ERROR] Please add the name of file and password as dfile <name> <pass> || 6 <name> <pass>")
				break
			}
			if len(e_e[2]) < 32 {
				fmt.Println("🔴[ERROR] Please add a password that is exactly 32 characters")
				break
			}
			err := helper.DecryptAndRecoverOriginal(e_e[1], e_e[2], current_folder_location)
			if err != nil {
				fmt.Println(err)
			}
		case "del":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("🔴[ERROR] Please add the name of the file to encrypt as del <name> || 7 <name>")
				break
			}
			err := helper.DeleteFile(e_e[1], current_folder_location)
			if err != nil {
				fmt.Println("🔴[ERROR] Error in deleting file as ", err)
				break
			}
			msg := fmt.Sprintf("🟢[DONE] the file %v has been deleted", e_e[1])
			fmt.Println(msg)
		case "vi":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("🔴[ERROR] Please add the name of the file to open as vi <name> || 5 <name>")
				break
			}
			cmd := exec.Command("vi", current_folder_location+"/"+e_e[1])
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("🔴[ERROR] Error executing vi: %v\n", err)
			}
		case "up":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println(`🔴[ERROR] Please add the name of the file to be uploaded as "up <name_with_format>"`)
				break
			}
			err = auth.UploadToGoogleDrive(e_e[1], current_folder_location+"/")
			if err != nil {
				fmt.Printf("🔴[ERROR] %v", err)
				break
			}
			fmt.Println("🟢[DONE] the file has been uploaded to Google Drive")
		case "down":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println(`🔴[ERROR] Please add the name of the file to downloaded as "down <name_with_format>"`)
				break
			}
			err := auth.DownloadFromGoogleDrive(e_e[1], current_folder_location+"/")
			if err != nil {
				fmt.Printf("🔴[ERROR] %v\n", err)
				break
			}
			fmt.Println("🟢[DONE] the file has been downloaded from Google Drive")
		case "help", "h":
			fmt.Println(options)
		case "exit":
			return
		default:
			fmt.Println("📎 Please enter input based on one of the following options:-" + options)
		}
	}

}
