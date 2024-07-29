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
	//entry
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Peril, what would you like to do?")
	var current_folder_location string
	for {
		var e string
		if scanner.Scan() {
			e = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
		}
		//We might have 0,1,2,3,4 only as cmd, in that case certain ops have to be performed so splitting the string otherwise os command will be executed
		e_e := strings.Split(e, " ")
		fmt.Println(e_e)
		switch e_e[0] {
		//open folder
		case "ofold", "1":
			fmt.Println(1)
			// folder.openFolder()
		//make folder
		case "mfold", "2":
			var new_f_name string
			var err error
			//getting the location
			current_folder_location, err = helper.SelectFolder()
			if err != nil {
				fmt.Println("Error in opening folder", e)
			}
			fmt.Println("Please enter the name of the folder :")
			if scanner.Scan() {
				new_f_name = scanner.Text()
			}
			current_folder_location = filepath.Join(current_folder_location, new_f_name)
			err = os.Mkdir(current_folder_location, 0755)
			if err != nil {
				//log the error
				fmt.Println("Error creating new folder:", err)
				return
			} else {
				//log that file has been created
			}
			// helper.makeFile()
		//open file
		case "ofile", "3":
			fmt.Println(3)
			// file.openFile()
		//make file
		case "mfile", "4":
			fmt.Println(4)
			// file.makeFile()
		case "exit":
			break
		default:
		}
	}

}
