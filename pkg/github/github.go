package github

import (
	"context"
	"fmt"
	"time"

	gh "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	minCommits = 3
	maxCommits = 3
)

// Commit represents a github commit
type Commit struct {
	ID      string
	Message string
	User    string
	Date    time.Time
}

// Repo represents a github repository
type Repo struct {
	Name        string
	Description string
	Owner       string
	URL         string
	Language    string
	Tags        []string
	Commits     []Commit
}

func newCommit(commit *gh.RepositoryCommit) Commit {
	msg := *commit.Commit.Message
	if len(msg) > 100 {
		msg = msg[:100] + "..."
	}

	return Commit{
		ID:      (*commit.SHA)[:8],
		Message: msg,
		User:    *commit.Commit.Author.Name,
		Date:    *commit.Commit.Author.Date,
	}
}

func newRepo(ctx context.Context, client *gh.Client, repo *gh.Repository) Repo {
	opts := &gh.CommitsListOptions{}
	opts.Page = 1
	opts.PerPage = maxCommits
	commits, _, _ := client.Repositories.ListCommits(ctx, *repo.Owner.Login, *repo.Name, opts)

	if repo.Description == nil {
		repo.Description = new(string)
	}

	result := Repo{
		Name:        *repo.Name,
		Description: *repo.Description,
		Owner:       *repo.Owner.Login,
		URL:         *repo.HTMLURL,
		Tags:        repo.Topics,
		Language:    *repo.Language,
	}

	for _, commit := range commits {
		result.Commits = append(result.Commits, newCommit(commit))
	}

	return result
}

// GetReposAndStars returns the user's owned repos and starred repos as slices
func GetReposAndStars(user string, token string) (repos []Repo, stars []Repo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := gh.NewClient(tc)

	userRepos, _, err := client.Repositories.List(ctx, user, nil)
	if err != nil {
		return nil, nil, err
	}

	for _, repo := range userRepos {
		if *repo.Fork {
			continue
		}

		r := newRepo(ctx, client, repo)
		if len(r.Commits) >= minCommits {
			repos = append(repos, r)
		}
	}

	starredRepos, _, err := client.Activity.ListStarred(ctx, user, &gh.ActivityListStarredOptions{})
	if err != nil {
		return nil, nil, err
	}

	for _, repo := range starredRepos {
		r := newRepo(ctx, client, repo.Repository)
		stars = append(stars, r)
	}

	return repos, stars, nil
}
