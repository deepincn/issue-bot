package repository

import (
	"context"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

type Repository struct {
	Owner string
	Name string
	Client *github.Client
	context context.Context
	*github.Repository
}

func ListRepository(owner string, token string) ([]*Repository, error) {
	var result []*Repository
	var allRepos []*github.Repository
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	opts := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	for {
		repos, resp, err := client.Repositories.List(ctx, owner, opts)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	for _, repo := range allRepos {
		re, err := NewRepository(owner, repo.GetName(), token)
		if err != nil {
			continue
		}
		result = append(result, re)
	}
	return result, nil
}

func NewRepository(owner string, repo string, token string) (*Repository, error)  {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	_repo, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Owner: owner,
		Name: repo,
		context: ctx,
		Client: client,
		Repository: _repo,
	}, nil
}

// Transfer a repository transfer function
// owner is new owner of the repository
func (r *Repository)Transfer(owner string) error {
	request := github.TransferRequest{
		NewOwner: owner,
	}

	_, _, err := r.Client.Repositories.Transfer(r.context, r.Owner, r.Name, request)

	return err
}

func (r* Repository)Update() error {
	_, _, err := r.Client.Repositories.Edit(r.context, r.Owner, r.Name, r.Repository)
	return err
}
