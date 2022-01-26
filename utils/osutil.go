package utils

import "os"

func GetWorkingDirectory() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return dir, nil
}

func MakeDirectory(path string) error {
	err := os.MkdirAll(path, os.ModeDir)

	if err != nil {
		return err
	}

	return nil
}
