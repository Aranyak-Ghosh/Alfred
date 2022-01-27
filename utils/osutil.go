package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func GetWorkingDirectory() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return dir, nil
}

func GetConfigPath() (string, error) {
	configPath, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	configPath = path.Join(configPath, "alfred")

	return configPath, nil
}

func GetDirectoryContents(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var fileNames []string

	for _, file := range files {
		file_data := file.Name()
		if file.IsDir() {
			file_data = fmt.Sprintf("%-15s%4s", file_data+"/", "dir")
		} else {
			file_data = fmt.Sprintf("%-15s%4s", file_data+"/", "file")

		}
		fileNames = append(fileNames, file_data)
	}

	return fileNames, nil
}

func MakeDirectory(path string) error {
	err := os.MkdirAll(path, os.ModeDir)

	if err != nil {
		return err
	}

	return nil
}

func WriteFile(path string, data []byte) error {

	err := os.WriteFile(path, data, 0664)

	if err != nil {
		return err
	}
	return nil
}

func AppendToFile(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0664)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		return err
	}

	return nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ioutil.ReadAll(file)
}

func DeleteDir(path string) error {
	err := os.RemoveAll(path)
	return err
}

func EnsureDependencyInstall() error {
	if err := ensureGitInstall(); err != nil {
		return err
	}

	return nil
}

func CloneProject(url string, projectName string) error {
	gitPath, err := getGitPath()
	if err != nil {
		return err
	}

	args := []string{gitPath, "clone", url}

	if projectName != "" {
		args = append(args, projectName)
	}

	cmd := &exec.Cmd{
		Path:   gitPath,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func InitEmptyGitRepo(projectName string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	projectPath := path.Join(wd, projectName)

	gitPath, err := getGitPath()

	if err != nil {
		return err
	}

	args := []string{gitPath, "init"}

	cmd := &exec.Cmd{
		Path:   gitPath,
		Dir:    projectPath,
		Args:   args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func OpenInCode(path string) error {

	codePath, err := getCodePath()

	if err != nil {
		return err
	}

	cmd := exec.Cmd{
		Path:   codePath,
		Dir:    path,
		Args:   []string{codePath, path},
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

func getCodePath() (string, error) {
	codePath, err := exec.LookPath(
		"code",
	)

	if err != nil {
		return "", err
	}

	return codePath, nil
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
