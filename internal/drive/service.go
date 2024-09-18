package drive

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"google.golang.org/api/drive/v3"
)

type Service struct {
	srv *drive.Service
}

func NewService(client *http.Client) (*Service, error) {
	srv, err := drive.New(client)
	if err != nil {
		return nil, err
	}
	return &Service{srv: srv}, nil
}

func (s *Service) UploadFile(filePath string, folderID string) error {
	fmt.Printf("Path file: %s\n", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileName := filepath.Base(filePath)
	driveFile := &drive.File{
		Name: fileName,
	}

	if folderID != "" && folderID != "-" {
		driveFile.Parents = []string{folderID}
	} else if folderID == "-" {

		driveFile.Parents = []string{"root"}
	}

	_, err = s.srv.Files.Create(driveFile).Media(file).Do()
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	return nil
}
func (s *Service) CreateFolder(folderName string, parentID string) (string, error) {
	folder := &drive.File{
		Name:     folderName,
		MimeType: "application/vnd.google-apps.folder",
	}

	if parentID == "-" {

		folder.Parents = []string{"root"}
	} else if parentID != "" {

		folder.Parents = []string{parentID}
	}

	file, err := s.srv.Files.Create(folder).Do()
	if err != nil {
		return "", err
	}

	return file.Id, nil
}
