package utils

import (
	"alfred/models"


	"gopkg.in/yaml.v2"
)

func ParseConfigString(config []byte) (models.RepoMap, error) {
	var availableRepo models.RepoMap

	err := yaml.Unmarshal(config, &availableRepo)

	return availableRepo, err
}

func SerializeConfig(config models.RepoMap) ([]byte, error) {
	return yaml.Marshal(config)
}
