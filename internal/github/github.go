package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v45/github"
)

type Github struct {
	Client *github.Client
}

type BranchPolicy struct {
	Name   string
	Policy *github.ProtectionRequest
}

func NewGithubClient(auth *http.Client) *Github {
	return &Github{
		Client: github.NewClient(auth),
	}
}

func (gh *Github) GetRepository(owner string, repositoryName string) (*github.Repository, error) {
	// Get repository
	repo, resp, err := gh.Client.Repositories.Get(context.Background(), owner, repositoryName)
	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("invalid credentials provided")
	} else if resp.StatusCode == 404 {
		return nil, fmt.Errorf("repository not found")
	} else if err != nil {
		return nil, err
	}

	return repo, nil
}

func (gh *Github) RepositoryExists(owner string, repositoryName string) bool {
	repo, _ := gh.GetRepository(owner, repositoryName)

	return repo != nil
}

// TODO: Add customiseable visibility options
func (gh *Github) CreateTeamRepository(owner string, name string, visibility string, teamId int64) (*github.Repository, error) {
	r := &github.Repository{
		Name:       &name,
		Visibility: &visibility,
		TeamID:     &teamId,
	}

	repo, resp, _ := gh.Client.Repositories.Create(context.Background(), owner, r)
	if resp.StatusCode == 422 {
		return nil, fmt.Errorf("403 Github repository creation: '%s' forbidden", name)
	}
	if resp.StatusCode == 422 {
		return nil, fmt.Errorf("422 validation for Github repository creation: '%s' has failed", name)
	}

	return repo, nil
}

func (gh *Github) CreateRepository(owner string, name string, visibility string) (*github.Repository, error) {
	r := &github.Repository{
		Name:       &name,
		Visibility: &visibility,
	}

	repo, resp, _ := gh.Client.Repositories.Create(context.Background(), owner, r)
	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("403 Github repository creation: '%s' forbidden", name)
	}
	if resp.StatusCode == 422 {
		return nil, fmt.Errorf("422 validation for Github repository creation: '%s' has failed", name)
	}

	return repo, nil
}

func (gh *Github) GetTeam(organisation string, slug string) (int64, error) {
	team, resp, err := gh.Client.Teams.GetTeamBySlug(context.Background(), organisation, slug)
	if resp.StatusCode == 404 {
		return 0, fmt.Errorf("404 team not found")
	} else if err != nil {
		return 0, err
	}
	return *team.ID, nil
}

func (gh *Github) StartImport(owner string, repositoryName string, URL string, username string, password string) (*github.Import, error) {
	vcs := "git"
	in := &github.Import{
		VCSURL:      &URL,
		VCS:         &vcs,
		VCSUsername: &username,
		VCSPassword: &password,
	}
	i, resp, _ := gh.Client.Migrations.StartImport(context.Background(), owner, repositoryName, in)
	if resp.StatusCode == 422 {
		return nil, fmt.Errorf("422 validation for Github repository creation: '%s' has failed", repositoryName)
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("404 resource not found")
	}

	return i, nil
}
