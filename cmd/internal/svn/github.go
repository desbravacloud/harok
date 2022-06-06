package svn

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var gitClient *github.Client = runGithub()
var gitOrg, _ = getGitHubOrganization()
var token, _ = getGitHubToken()
var jenkinsAddress, _ = getJenkinsAddress()

type Content struct {
	Name   string            `json:"name"`
	Active bool              `json:"active"`
	Events []string          `json:"events"`
	Config map[string]string `json:"config"`
}

func runGithub() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(setGithubCtx(), ts)

	client := github.NewClient(tc)

	return client
}

func setGithubCtx() context.Context {
	ctx := context.Background()

	return ctx
}

func CreateRepo(name string, private bool) error {
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(private),
	}
	_, _, err := gitClient.Repositories.Create(setGithubCtx(), gitOrg, repo)
	if err != nil {
		return err
	}
	return nil
}

func CreateWebHooks(repoName string) error {
	reqBody, err := json.Marshal(Content{
		Name:   "web",
		Active: true,
		Events: []string{"push", "pull_request"},
		Config: map[string]string{"url": jenkinsAddress + "/multibranch-webhook-trigger/invoke?token=" + repoName, "content_type": "json"},
	})
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+gitOrg+"/"+repoName+"/hooks", bytes.NewBuffer(reqBody))
	req.Header = http.Header{
		"Authorization": {"Bearer " + token},
		"Content-Type":  {"application/json"},
	}
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	defer req.Body.Close()

	// get reponse from creating webhooks
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(body))

	return nil
}
