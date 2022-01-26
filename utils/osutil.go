package utils

import (
	"os"
	"os/exec"
)

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

func EnsureGitInstall() error {
	_, err := exec.LookPath(
		"git",
	)
	return err
}
