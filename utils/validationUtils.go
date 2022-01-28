package utils

import "regexp"

func ValidateGitUrl(gitUrl string) bool {
	matched, err := regexp.Match(`^((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?`, []byte(gitUrl))

	if err != nil {
		return false
	} else {
		return matched
	}
}
