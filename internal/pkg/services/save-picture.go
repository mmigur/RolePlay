package services

import (
	"fmt"
	"os"
	"path/filepath"
)

func SavePicture(picture []byte, userName string) (string, error) {
	filename := fmt.Sprintf("%s.png", userName)
	directory := "./media"

	filePath := filepath.Join(directory, filename)

	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(picture); err != nil {
		return "", err
	}

	var link string
	if picture != nil {
		link = "http://80.87.110.246:3211/media/" + filename
	}
	return link, nil
}
