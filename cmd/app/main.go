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
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var filePath, folderID string

		fmt.Print("Masukkan path file: ")
		fmt.Scan(&filePath)

		fmt.Print("Masukkan folder ID (kosong untuk folder utama, '-' untuk root): ")
		fmt.Scan(&folderID)

		if err := srv.UploadFile(filePath, folderID); err != nil {
			fmt.Printf("Error uploading file: %v\n", err)
		} else {
			fmt.Println("File berhasil diunggah.")
		}
	case 2:
		var folderName, parentID string
		fmt.Print("Masukkan nama folder: ")
		fmt.Scan(&folderName)

		fmt.Print("Masukkan parent folder ID (kosong untuk folder utama): ")
		fmt.Scan(&parentID)

		if parentID == "-" {
			fmt.Println("Membuat folder di direktori utama...")
		} else if parentID != "" {
			fmt.Printf("Membuat folder di direktori: %s\n", parentID)
		} else {
			fmt.Println("Membuat folder di direktori utama...")
		}

		folderID, err := srv.CreateFolder(folderName, parentID)
		if err != nil {
			fmt.Printf("Error creating folder: %v\n", err)
		} else {
			fmt.Printf("Folder '%s' berhasil dibuat dengan ID '%s'\n", folderName, folderID)
		}
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
