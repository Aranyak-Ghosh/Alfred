package models

type RepoMap = map[string]Repo

type Repo struct {
	Url    string
	Branch string
}
