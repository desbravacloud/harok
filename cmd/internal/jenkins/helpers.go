package jenkins

import (
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

var home string = os.Getenv("HOME")
var path string = "/.harok/config.json"

func getJenkinsAddress() (jenkinsAddress string, err error) {

	content, err := ioutil.ReadFile(home + path) // the file is inside the local directory
	if err != nil {
		return jenkinsAddress, err
	}
	jenkinsAddress = gjson.Get(string(content), "credentials.jenkins_address").String()

	return jenkinsAddress, err
}

func getJenkinsUser() (jenkinsUser string, err error) {

	content, err := ioutil.ReadFile(home + path) // the file is inside the local directory
	if err != nil {
		return jenkinsUser, err
	}
	jenkinsUser = gjson.Get(string(content), "credentials.jenkins_user").String()

	return jenkinsUser, err
}

func getJenkinsToken() (jenkinsToken string, err error) {

	content, err := ioutil.ReadFile(home + path) // the file is inside the local directory
	if err != nil {
		return jenkinsToken, err
	}
	jenkinsToken = gjson.Get(string(content), "credentials.jenkins_token").String()

	return jenkinsToken, err
}

func getGitHubOrganization() (org string, err error) {

	content, err := ioutil.ReadFile(home + path) // the file is inside the local directory
	if err != nil {
		return org, err
	}
	org = gjson.Get(string(content), "credentials.github_org").String()

	return org, err
}
