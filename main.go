package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xmp-er/peril/helper"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Peril, what would you like to do?\n-ofold || 1 [Open folder]\n-mfold || 2 [Make folder]\n-ofile || 3 [Open file]\n-mfile || 4 [Make file]")
	var current_folder_location string
	for {
		var e string
		if scanner.Scan() {
			e = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			//Log that there was a error reading input
			fmt.Println("Error reading input:", err)
		}
		e_e := strings.Split(e, " ")
		fmt.Println(e_e)
		switch e_e[0] {
		case "ofold", "1":
			var err error
			current_folder_location, err = helper.SelectFolder()
			if err != nil {
				//Log that there was error in opening the folder
				fmt.Println("Error in opening folder", e)
			}
			//Log that folder was opened in this location
			fmt.Println("Current directory set to :", current_folder_location)
		case "mfold", "2":
			var err error
			current_folder_location, err = helper.SelectFolder()
			if err != nil {
				//log that there was a error in opening the folder
				fmt.Println("Error in opening folder", e)
			}
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("Please add the name of the folder as mfold <name> || 2 <name>")
				break
			}
			current_folder_location = filepath.Join(current_folder_location, e_e[1])
			err = os.Mkdir(current_folder_location, 0755)
			if err != nil {
				//log the error
				fmt.Println("Error creating new folder:", err)
				return
			}
			//log that file has been created
		case "ofile", "3":
			var err error
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("Please add the name of the file to open as ofile <name> || 3 <name>")
				break
			}
			err = helper.OpenOrCreateFile(e_e[1])
			if err != nil {
				//log the error
				fmt.Println("There was a error in opening the file:", err)
			}
			//Log that the file was opened at that location
		case "mfile", "4":
			if len(e_e) < 2 {
				//log that the name of the file was not provided
				fmt.Println("Please add the name of the file to open as mfile <name> || 4 <name>")
				break
			}
			err := helper.OpenOrCreateFile(e_e[1])
			if err != nil {
				//log the error
				fmt.Println("There was a error in making the file:", err)
			}
			//Log that file was made with the name at the location
		case "exit":
			return
		default:
			fmt.Println("Please enter input based on one of the following options:-\n-ofold || 1 [Open folder]\n-mfold || 2 [Make folder]\n-ofile || 3 [Open file]\n-mfile || 4 [Make file]")
		}
	}

}
