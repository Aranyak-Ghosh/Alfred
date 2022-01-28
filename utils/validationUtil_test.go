package utils

import "testing"

func TestGitUrlValidation(t *testing.T) {
	gitUrl := "http://github.com/user/project.git"
	ok := ValidateGitUrl(gitUrl)

	if !ok {
		t.Errorf("Git url %s should be valid", gitUrl)
	}
}
