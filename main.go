package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	auth "github.com/xmp-er/peril/Auth"
	"github.com/xmp-er/peril/helper"
)

var options string = "\n⚪️ ofold <name> || 1 <name> [Open folder]\n⚪️ mfold <name> || 2 <name> [Make folder]\n⚪️ ofile <name> || 3 <name> [Open file]\n⚪️ mfile <name> || 4 <name> [Make file]\n⚪️ efile <name> || 5 <name> [Encrypt file]\n⚪️ dfile <name> || 6 <name> [Decrypt File]\n⚪️ del <name_with_extension> || 7 <name_with_extension> [Delete file]\n⚪️ up <Drive_Path> [Upload a file to Google Drive]\n⚪️ down <Drive_Path> [Download a file from Google Drive]"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Peril, what would you like to do?" + options)
	var current_folder_location string
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
		case "ofold", "1":
			var err error
			current_folder_location, err = helper.SelectFolder()
			if err != nil {
				//Log that there was error in opening the folder
				fmt.Println("🔴[ERROR] Error in opening folder", e)
				break
			}
			//Log that folder was opened in this location
			fmt.Println("🟢[DONE] Current directory set to :", current_folder_location)
		case "mfold", "2":
			var err error
			current_folder_location, err = helper.SelectFolder()
			if err != nil {
				//log that there was a error in opening the folder
				fmt.Println("🔴[ERROR] Error in opening folder", e)
			}
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("🔴[ERROR] Please add the name of the folder as mfold <name> || 2 <name>")
				break
			}
			current_folder_location = filepath.Join(current_folder_location, e_e[1])
			err = os.Mkdir(current_folder_location, 0755)
			if err != nil {
				//log the error
				fmt.Println("🔴[ERROR] Error creating new folder:", err)
				break
			}
			//log that file has been created
			fmt.Println("🟢[DONE] Folder created at :", current_folder_location)
		case "ofile", "3":
			var err error
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("🔴[ERROR] Please add the name of the file to open as ofile <name> || 3 <name>")
				break
			}
			err = helper.OpenOrCreateFile(e_e[1], "ofile")
			if err != nil {
				//log the error
				fmt.Println("🔴[ERROR] There was a error in opening the file:", err)
			}
			//Log that the file was opened at that location
		case "mfile", "4":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("🔴[ERROR] Please add the name of the file to open as mfile <name> || 4 <name>")
				break
			}
			err := helper.OpenOrCreateFile(e_e[1], "mfile")
			if err != nil {
				//log the error
				fmt.Println("🔴[ERROR] There was a error in making the file:", err)
				break
			}
			//Log that file was made with the name at the location
			msg := fmt.Sprintf("🟢[DONE] The file %v.md has been created", e_e[1])
			fmt.Println(msg)
		case "efile", "5":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("🔴[ERROR] Please add the name of the file to encrypt as efile <name> || 5 <name>")
				break
			}
			pass, err := helper.EncryptAndDeleteOriginal(e_e[1])
			if err != nil {
				fmt.Println(err)
				break
			}
			finalRes := fmt.Sprintf("🟢[DONE] Password for decrypting %v is %v", (e_e[1] + ".enc"), pass)
			fmt.Println(finalRes)
		case "dfile", "6":
			if len(e_e) < 3 {
				fmt.Println("🔴[ERROR] Please add the name of file and password as dfile <name> <pass> || 6 <name> <pass>")
				break
			}
			if len(e_e[2]) < 32 {
				fmt.Println("🔴[ERROR] Please add a password that is exactly 32 characters")
				break
			}
			err := helper.DecryptAndRecoverOriginal(e_e[1], e_e[2])
			if err != nil {
				fmt.Println(err)
			}
		case "del", "7":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("🔴[ERROR] Please add the name of the file to encrypt as del <name> || 7 <name>")
				break
			}
			err := helper.DeleteFile(e_e[1])
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
			cmd := exec.Command("vi", e_e[1])
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("🔴[ERROR] Error executing vi: %v\n", err)
			}
		case "up":
			fileLocation, err := helper.SelectFile()
			if err != nil {
				msg := fmt.Sprintf("🔴[ERROR] %v", err)
				fmt.Println(msg)
				break
			}
			err = auth.UploadToGoogleDrive(fileLocation)
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
			if current_folder_location == "" {
				location, err := helper.GetCurrentDirectory()
				if err != nil {
					fmt.Printf(`🔴[ERROR] %v\n"`, err)
					break
				}
				current_folder_location = location
			}
			err := auth.DownloadFromGoogleDrive(e_e[1], current_folder_location)
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
