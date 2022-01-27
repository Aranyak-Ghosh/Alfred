package services

import "alfred/utils"

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
