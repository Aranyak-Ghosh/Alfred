package utils

import (
	"testing"
)

func TestGetDir(t *testing.T) {
	dir, err := GetWorkingDirectory()

	if err != nil {
		t.Errorf("Error: %s\n", err)
	} else if dir != "c:\\Projects\\Personal\\go-code-generator\\utils" {
		t.Fail()
	}
}

func TestMakeDir(t *testing.T) {
	err := MakeDirectory("c:\\Projects\\Personal\\go-code-generator\\utils\\test\\sub")

	if err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestGitInstall(t *testing.T) {
	if err := ensureGitInstall(); err != nil {
		t.Errorf("Error: %s\n", err)
		t.Fail()
	}
}

func TestGoInstall(t *testing.T) {
	if err := ensureGoInstall(); err != nil {
		t.Errorf("Error: %s\n", err)
		t.Fail()
	}
}

func TestCloneProject(t *testing.T) {
	if err := CloneProject("https://github.com/Aranyak-Ghosh/gode-gen.git", "TestClone"); err != nil {
		t.Errorf("Error: %s\n", err)
		t.Fail()
	}
}
