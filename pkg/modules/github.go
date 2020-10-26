package modules

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/razzie/beepboop"
	"github.com/razzie/gorzsony.com/pkg/github"
	"github.com/razzie/gorzsony.com/pkg/layout"
)

type githubView struct {
	Tag         string
	GithubRepos []github.Repo
	GithubStars []github.Repo
}

var tagLangMap = map[string]string{
	"cpp":    "C++",
	"csharp": "C#",
	"go":     "Go",
}

func shuffleRepos(repos []github.Repo) []github.Repo {
	clone := append(repos[:0:0], repos...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(clone), func(i, j int) { clone[i], clone[j] = clone[j], clone[i] })
	return clone
}

func limitRepos(repos []github.Repo, maxRepos int) []github.Repo {
	return repos[:min(maxRepos, len(repos))]
}

func filterRepos(repos []github.Repo, tag string) (results []github.Repo) {
	tag = strings.ToLower(tag)
	lang, _ := tagLangMap[tag]
	for _, repo := range repos {
		if repo.Language == lang {
			results = append(results, repo)
			continue
		}
		for _, t := range repo.Tags {
			if t == tag {
				results = append(results, repo)
				continue
			}
		}
	}
	return
}

// Github returns the github repo and star modules
func Github(token string) (reposModule *layout.Module, starsModule *layout.Module) {
	var repos []github.Repo
	var stars []github.Repo

	go func() {
		ticker := time.NewTicker(time.Minute * 30)
		for ; true; <-ticker.C {
			tmpRepos, tmpStars, err := github.GetReposAndStars("razzie", token)
			if err != nil {
				log.Println(err)
				continue
			}

			repos, stars = tmpRepos, tmpStars
		}
	}()

	reposModule = &layout.Module{
		Name:            "Github Repos",
		ContentTemplate: getContentTemplate("github_repos"),
		Handler: func(pr *beepboop.PageRequest) interface{} {
			if len(repos) == 0 {
				return nil
			}
			var v *githubView
			tag := getTag(pr)
			if len(tag) > 0 {
				v = &githubView{
					Tag:         tag,
					GithubRepos: shuffleRepos(filterRepos(repos, tag)),
				}
			} else {
				v = &githubView{
					GithubRepos: limitRepos(shuffleRepos(repos), 8),
				}
			}
			if len(v.GithubRepos) == 0 {
				return nil
			}
			return v
		},
	}
	starsModule = &layout.Module{
		Name:            "Github Stars",
		ContentTemplate: getContentTemplate("github_stars"),
		Handler: func(pr *beepboop.PageRequest) interface{} {
			if len(stars) == 0 {
				return nil
			}
			v := &githubView{
				GithubStars: limitRepos(shuffleRepos(stars), 8),
			}
			if len(v.GithubStars) == 0 {
				return nil
			}
			return v
		},
	}
	return
}
