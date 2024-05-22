package services

import (
	"RolePlayModule/internal/utils/config"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

func SaveImage(cfg config.Config, base64Image string) (string, error) {
	name, _ := uuid.NewRandom()
	filename := fmt.Sprintf("%s.png", name)
	directory := "./media"

	filePath := filepath.Join(directory, filename)

	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", err
	}

	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(imageBytes); err != nil {
		return "", err
	}

	link := cfg.NgrokUrl + "/media/" + filename
	return link, nil
}
