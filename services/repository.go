package services

import (
	"alfred/models"
	"alfred/utils"
	"fmt"
	"os"
	"path"
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

func AddRepoToStore(repos map[string]models.Repo, overwrite bool) error {
	currentRepo, err := GetRepoStore()

	if err != nil {
		return err
	}

	for _, repo := range repos {
		if ok := utils.ValidateGitUrl(repo.Url); !ok {
			return fmt.Errorf("Invalid git url %s", repo.Url)
		}
	}

	if !overwrite {
		for tag, repo := range repos {
			if val, ok := currentRepo[tag]; ok {
				return fmt.Errorf("Tag %s already exists in repository store. Current URL: %s, Current Branch: %s, New URL: %s, New Branch: %s", tag, val.Url, val.Branch, repo.Url, repo.Branch)
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

func UpdateRepoStore(repos map[string]models.Repo, create bool) error {
	currentRepo, err := GetRepoStore()

	if err != nil {
		return err
	}

	for tag, repo := range repos {
		if ok := utils.ValidateGitUrl(repo.Url); !ok {
			return fmt.Errorf("Invalid git url %s", repo.Url)
		}
		if val, ok := currentRepo[tag]; ok {
			fmt.Printf("Tag %s found with repository: \nUrl: %s\nBranch: %s\n", tag, val.Url, val.Branch)
			fmt.Printf("Updating tag to \nUrl: %s\nBranch: %s\n", repo.Url, repo.Branch)
			if val != repo {
				currentRepo[tag] = repo
			}
		} else if create {
			fmt.Printf("Tag %s not found! Creating tag\n", tag)

			currentRepo[tag] = repo
		} else {
			return fmt.Errorf("Tag %s not found in repository store", tag)
		}
	}

	txt, err := utils.SerializeConfig(currentRepo)
	if err != nil {
		return err
	}

	configPath, err := utils.GetConfigPath()
	if err != nil {
		return err
	}

	err = utils.WriteFile(configPath+"/repositories.yaml", txt)
	return err
}

func RemoveRepoFromStore(tags []string) error {
	currentRepo, err := GetRepoStore()

	if err != nil {
		return err
	}

	for _, tag := range tags {
		if _, ok := currentRepo[tag]; ok {
			delete(currentRepo, tag)
		} else {
			return fmt.Errorf("Tag %s not found in repository store", tag)
		}
	}

	txt, err := utils.SerializeConfig(currentRepo)
	if err != nil {
		return err
	}
	configPath, err := utils.GetConfigPath()
	if err != nil {
		return err
	}

	err = utils.WriteFile(configPath+"/repositories.yaml", txt)
	return err

}

func CreateProject(tag string, projectName string, gitInit bool, codeOpen bool, explorerOpen bool) error {
	wd, err := os.Getwd()
	if tag != "" {
		fmt.Printf("Creating project %s with tag %s\n", projectName, tag)
		repoStore, err := GetRepoStore()

		if err != nil {
			return err
		}
		if repo, ok := repoStore[tag]; ok {
			fmt.Printf("Cloning repository %s\n", repo)
			if projectName == "" {
				projectName = tag
			}
			err = utils.CloneProject(repo.Url, repo.Branch, projectName)
			if err != nil {
				return err
			}
			fmt.Println("Project cloned")

			if gitInit {
				fmt.Printf("Deleteing .git folder\n")
				err = utils.DeleteDir(path.Join(wd, projectName, ".git"))
				if err != nil {
					return err
				}
				fmt.Printf("Initializing git repository\n")
				err = utils.InitEmptyGitRepo(projectName)
				if err != nil {
					return err
				}
			}

		} else {
			return fmt.Errorf("Tag %s not found in repository store", tag)
		}
	} else {
		fmt.Printf("Creating project %s\n", projectName)

		if err != nil {
			return err
		}
		projectDirectory := path.Join(wd, projectName)
		err = utils.MakeDirectory(projectDirectory)
		if err != nil {
			return err
		}
		fmt.Println("Project directory created")
		if gitInit {
			fmt.Printf("Initializing git repository\n")
			err = utils.InitEmptyGitRepo(projectName)
			if err != nil {
				return err
			}
		}
	}

	if codeOpen {
		fmt.Printf("Opening project in code\n")
		projectDirectory := path.Join(wd, projectName)
		err = utils.OpenInCode(projectDirectory)
		if err != nil {
			return err
		}
	}

	if explorerOpen {
		projectDirectory := wd + string(os.PathSeparator) + projectName
		fmt.Printf("Opening project in explorer\n")

		err = utils.OpenInExplorer(projectDirectory)
		if err != nil {
			return err
		}
	}

	return nil
}
