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
