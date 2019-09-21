package main

import (
	"math/rand"
	"time"
)

type View struct {
	Projects          []Project
	GithubReposLoaded bool
	GithubRepos       []Repo
	GithubStarsLoaded bool
	GithubStars       []Repo
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shuffleProjects(projects []Project) []Project {
	clone := append(projects[:0:0], projects...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(clone), func(i, j int) { clone[i], clone[j] = clone[j], clone[i] })
	return clone
}

func shuffleRepos(repos []Repo, maxRepos int) []Repo {
	clone := append(repos[:0:0], repos...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(clone), func(i, j int) { clone[i], clone[j] = clone[j], clone[i] })
	return clone[:min(maxRepos, len(clone))]
}

func NewView(projects []Project, repos []Repo, stars []Repo) View {
	return View{
		Projects:          shuffleProjects(projects),
		GithubReposLoaded: len(repos) > 0,
		GithubRepos:       shuffleRepos(repos, 6),
		GithubStarsLoaded: len(stars) > 0,
		GithubStars:       shuffleRepos(stars, 6)}
}
