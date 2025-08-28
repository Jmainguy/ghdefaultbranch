package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v74/github"
	"golang.org/x/oauth2"
)

func getRepo(token, ownerRepo string) *github.Repository {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	ownerRepoSplit := strings.Split(ownerRepo, "/")
	if len(ownerRepoSplit) != 2 {
		fmt.Printf("%s is not a valid repository name, Should look like Jmainguy/k8sCapCity\n", ownerRepo)
		os.Exit(1)
	}
	owner := ownerRepoSplit[0]
	repoString := ownerRepoSplit[1]
	repo, _, err := client.Repositories.Get(ctx, owner, repoString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return repo
}

func getUserRepos(ctx context.Context, client *github.Client) ([]*github.Repository, error) {
	var repos []*github.Repository
	opts := &github.RepositoryListByAuthenticatedUserOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    1,
		},
		Affiliation: "owner,collaborator",
	}

	// loop through user repository list
	for opts.Page > 0 {
		list, resp, err := client.Repositories.ListByAuthenticatedUser(ctx, opts)
		if err != nil {
			return nil, err
		}
		repos = append(repos, list...)

		// increment the next page to retrieve
		opts.Page = resp.NextPage
	}

	return repos, nil
}

func getRepos(token string) []*github.Repository {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, err := getUserRepos(ctx, client)
	if err != nil {
		if strings.Contains(err.Error(), "401 Bad credentials") {
			fmt.Println("Token did not work, 401 Bad credentials")
		} else {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	return repos
}
