package utils

import (
	"reflect"
	"testing"
)

func TestConfigParser(t *testing.T) {
	config := `
repo1: https://github.com/Aranyak-Ghosh/gode-gen.git
repo2: https://github.com/Aranyak-Ghosh/go-get-started.git`

	expected := map[string]string{
		"repo1": "https://github.com/Aranyak-Ghosh/gode-gen.git",
		"repo2": "https://github.com/Aranyak-Ghosh/go-get-started.git",
	}

	actual, err := ParseConfigString([]byte(config))

	if err != nil {
		t.Errorf("Error parsing config: %s", err)
		t.Fail()
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
		t.Fail()
	}
}

func TestConfigSerializer(t *testing.T) {
	config := map[string]string{
		"repo1": "https://github.com/Aranyak-Ghosh/gode-gen.git",
		"repo2": "https://github.com/Aranyak-Ghosh/go-get-started.git",
	}

	expected := `repo1: https://github.com/Aranyak-Ghosh/gode-gen.git
repo2: https://github.com/Aranyak-Ghosh/go-get-started.git
`

	actual, err := SerializeConfig(config)

	if err != nil {
		t.Errorf("Error parsing config: %s", err)
		t.Fail()
	}

	if string(actual[:]) != expected {
		t.Errorf("Expected \n%v \ngot \n%v", expected, string(actual[:]))
		t.Fail()
	}
}
