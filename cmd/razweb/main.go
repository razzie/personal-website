package main

//go:generate go run ../../tools/go-bindata/ -prefix ../../assets ../../assets/...

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/razzie/gorzsony.com/pkg/github"
)

var (
	// Projects contains my hobby projects
	Projects []Project
	// Repos contains my github owned repos
	Repos []github.Repo
	// Stars contains my github starred repos
	Stars []github.Repo
)

func main() {
	index, err := Asset("index.html")
	if err != nil {
		panic(err)
	}

	tmpl := template.New("index")
	_, err = tmpl.Parse(string(index))
	if err != nil {
		panic(err)
	}

	resume, err := Asset("resume.html")
	if err != nil {
		panic(err)
	}

	Projects, err = LoadProjects()
	if err != nil {
		panic(err)
	}

	go func() {
		token, _ := Asset("github.token")
		ticker := time.NewTicker(time.Minute * 30)
		for ; true; <-ticker.C {
			repos, stars, err := github.GetReposAndStars("razzie", string(token))
			if err != nil {
				fmt.Println("error:", err)
				continue
			}

			Repos, Stars = repos, stars
		}
	}()

	fs := http.FileServer(
		&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: nil, Prefix: ""})

	http.Handle("/css/", fs)
	http.Handle("/img/", fs)
	http.Handle("/js/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, NewView(Projects, Repos, Stars))
	})
	http.HandleFunc("/tag/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tag := r.URL.Path[5:]
		tmpl.Execute(w, NewTagView(Projects, Repos, tag))
	})
	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "%s", resume)
	})

	http.ListenAndServe("localhost:8080", nil)
}
