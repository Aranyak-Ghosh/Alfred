package services

import (
	"alfred/models"
	"alfred/utils"
	"fmt"
)

func InitializeRepoStore() error {
	configPath, err := utils.GetConfigPath()

	if err != nil {
		return err
	}

	err = utils.MakeDirectory(configPath)

	if err != nil {
		return err
	}

	err = utils.WriteFile(configPath+"/repositories.yaml", []byte(""))

	return err
}

func GetRepoStore() (models.RepoMap, error) {
	configPath, err := utils.GetConfigPath()

	if err != nil {
		return nil, err
	}

	config, err := utils.ReadFile(configPath + "/repositories.yaml")

	if err != nil {
		return nil, err
	}

	repos, err := utils.ParseConfigString(config)

	if err != nil {
		return nil, err
	}

	return repos, nil
}

func AddRepoToStore(repos map[string]string, overwrite bool) error {
	currentRepo, err := GetRepoStore()

	if err != nil {
		return err
	}

	if !overwrite {
		for tag, url := range repos {
			if val, ok := currentRepo[tag]; ok {
				return fmt.Errorf("Tag %s already exists in repository store. Current URL: %s, New URL: %s", tag, val, url)
			}
		}
	}

	txt, err := utils.SerializeConfig(repos)
	if err != nil {
		return err
	}

	configPath, err := utils.GetConfigPath()
	if err != nil {
		return err
	}

	err = utils.AppendToFile(configPath+"/repositories.yaml", txt)
	return err
}

func AddReposToStoreFromFile(filePath string, overwrite bool) error {
	txt, err := utils.ReadFile(filePath)

	if err != nil {
		return err
	}

	repos, err := utils.ParseConfigString(txt)

	if err != nil {
		return err
	}

	err = AddRepoToStore(repos, overwrite)

	return err
}
