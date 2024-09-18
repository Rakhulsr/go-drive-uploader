package main

import (
	"fmt"
	"log"

	"github.com/Rakhulsr/go-drive-uploader/internal/auth"
	"github.com/Rakhulsr/go-drive-uploader/internal/drive"
)

func main() {
	config := auth.GetClientConfig()
	client := auth.GetClient(config)

	srv, err := drive.NewService(client)
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	var choice int
	fmt.Println("1. Upload File")
	fmt.Println("2. Create Folder")
	fmt.Print("Select Option: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var filePath, folderID string

		fmt.Print("Inpu file path: ")
		fmt.Scan(&filePath)

		fmt.Print("Input folder ID (empty or - for main folder(root)): ")
		fmt.Scan(&folderID)

		if err := srv.UploadFile(filePath, folderID); err != nil {
			fmt.Printf("Failed  uploading file, error: %v\n", err)
		} else {
			fmt.Println("File Uploaded")
		}
	case 2:
		var folderName, parentID string
		fmt.Print("Input Folder Name: ")
		fmt.Scan(&folderName)

		fmt.Print("Input Folder ParentID (empty or - for main folder(root)): ")
		fmt.Scan(&parentID)

		if parentID == "-" {
			fmt.Println("Creating folder in Main directory...")
		} else if parentID != "" {
			fmt.Printf("Creating folder in: %s ...\n", parentID)
		} else {
			fmt.Println("Creating folder in Main directory...")
		}

		folderID, err := srv.CreateFolder(folderName, parentID)
		if err != nil {
			fmt.Printf("Error creating folder: %v\n", err)
		} else {
			fmt.Printf("Folder '%s' created within ID '%s'\n", folderName, folderID)
		}
	default:
		fmt.Println("Option is not valid")
	}
}
