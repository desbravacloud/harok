package svn

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var gitClient *github.Client = runGithub()
var gitOrg, _ = getGitHubOrganization()

func runGithub() *github.Client {
	token, _ := getGitHubToken()
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

/*
func ListRepositories() {
	repos, _, err := gitClient.Repositories.ListByOrg(setGithubCtx(), gitOrg, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repos)
}
*/
