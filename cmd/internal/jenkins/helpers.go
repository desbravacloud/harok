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
