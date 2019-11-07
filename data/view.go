package data

import (
	"math/rand"
	"strings"
	"time"
)

// View contains data used by index.html template
type View struct {
	Base              string
	Tag               string
	ProjectsLoaded    bool
	Projects          []Project
	GithubReposLoaded bool
	GithubRepos       []Repo
	GithubStarsLoaded bool
	GithubStars       []Repo
}

var tagLangMap = map[string]string{
	"cpp":    "C++",
	"csharp": "C#",
	"go":     "Go",
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

func filterProjects(projects []Project, tag string) (results []Project) {
	for _, proj := range projects {
		tags := strings.Fields(proj.Tags)
		for _, t := range tags {
			if t == tag {
				results = append(results, proj)
				continue
			}
		}
	}
	return
}

func shuffleRepos(repos []Repo) []Repo {
	clone := append(repos[:0:0], repos...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(clone), func(i, j int) { clone[i], clone[j] = clone[j], clone[i] })
	return clone
}

func limitRepos(repos []Repo, maxRepos int) []Repo {
	return repos[:min(maxRepos, len(repos))]
}

func filterRepos(repos []Repo, tag string) (results []Repo) {
	lang, _ := tagLangMap[tag]
	for _, repo := range repos {
		if repo.Language == lang {
			results = append(results, repo)
		}
	}
	return
}

// NewView returns a new view using the given projects, owned repos and starred repos
func NewView(projects []Project, repos []Repo, stars []Repo) View {
	return View{
		Base:              "/",
		ProjectsLoaded:    len(projects) > 0,
		Projects:          shuffleProjects(projects),
		GithubReposLoaded: len(repos) > 0,
		GithubRepos:       limitRepos(shuffleRepos(repos), 6),
		GithubStarsLoaded: len(stars) > 0,
		GithubStars:       limitRepos(shuffleRepos(stars), 6),
	}
}

// NewTagView returns a new view that lacks intro and shows only tagged projects and repos
func NewTagView(projects []Project, repos []Repo, tag string) View {
	projects = filterProjects(projects, tag)
	repos = filterRepos(repos, tag)

	return View{
		Base:              "../",
		Tag:               tag,
		ProjectsLoaded:    len(projects) > 0,
		Projects:          shuffleProjects(projects),
		GithubReposLoaded: len(repos) > 0,
		GithubRepos:       shuffleRepos(repos),
	}
}
