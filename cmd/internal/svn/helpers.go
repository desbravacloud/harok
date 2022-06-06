package svn

import (
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

var home string = os.Getenv("HOME")
var path string = "/.harok/config.json"

func getGitHubToken() (token string, err error) {

	content, err := ioutil.ReadFile(home + path) // the file is inside the local directory
	if err != nil {
		return token, err
	}
	token = gjson.Get(string(content), "credentials.github_token").String()

	return token, err
}

func getGitHubOrganization() (org string, err error) {

	content, err := ioutil.ReadFile(home + path) // the file is inside the local directory
	if err != nil {
		return org, err
	}
	org = gjson.Get(string(content), "credentials.github_org").String()

	return org, err
}

func getJenkinsAddress() (jenkinsAddress string, err error) {

	content, err := ioutil.ReadFile(home + path) // the file is inside the local directory
	if err != nil {
		return jenkinsAddress, err
	}
	jenkinsAddress = gjson.Get(string(content), "credentials.jenkins_address").String()

	return jenkinsAddress, err
}
