package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/google/uuid"
)

func UploadSettingType(fileType string) error {
	allowedTypes := []string{"image/jpeg", "image/jpg", "image/png"}

	for i := 0; i < len(allowedTypes); i++ {
		el := allowedTypes[i]
		if el == fileType {
			return nil
		}
	}

	return errors.New("unsupported file type")
}

func UploadSettingName(originalName string) (string, error) {
	removeExt := strings.Split(originalName, ".")[0]
	fileExt := strings.Split(originalName, ".")[1]
	lowerCase := strings.ToLower(removeExt)
	removeSpacing := strings.ReplaceAll(lowerCase, " ", "")
	fileName := uuid.New().String() + "-" + removeSpacing + "." + fileExt

	return fileName, nil
}

func DeleteImage(fileName string) error {
	if fileName != "" {
		destination := "./publics/images/" + fileName
		if err := os.Remove(destination); err != nil {
			return err
		}
	}
	return nil
}

func DeleteAllImage() error {
	folderPath := "./publics/images/"

	// START open folder
	directory, err1 := os.Open(folderPath)
	if err1 != nil {
		return err1
	}
	// END open folder

	// START read all files in the folder
	files, err2 := directory.Readdir(-1)
	if err2 != nil {
		return err2
	}
	// END read all files in the folder

	// START delete all image
	for _, file := range files {
		if err := os.Remove(folderPath + file.Name()); err != nil {
			return err
		}
	}
	// END delete all image

	return nil
}
