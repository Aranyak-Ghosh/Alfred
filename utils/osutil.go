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

func EnsureDependencyInstall() error {
	if err := ensureGitInstall(); err != nil {
		return err
	}

	if err := ensureGoInstall(); err != nil {
		return err
	}

	return nil
}

func CloneProject(url string, projectName string) error {
	gitPath, err := getGitPath()
	if err != nil {
		return err
	}

	cmd := &exec.Cmd{
		Path:   gitPath,
		Args:   []string{gitPath, "clone", url, projectName},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func getGitPath() (string, error) {
	gitPath, err := exec.LookPath(
		"git",
	)

	if err != nil {
		return "", err
	}

	return gitPath, nil
}

func ensureGitInstall() error {
	_, err := exec.LookPath(
		"git",
	)
	return err
}

func ensureGoInstall() error {
	_, err := exec.LookPath(
		"go",
	)
	return err
}
