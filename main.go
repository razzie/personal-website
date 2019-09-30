package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

var (
	// Projects contains my hobby projects
	Projects []Project
	// Repos contains my github owned repos
	Repos []Repo
	// Stars contains my github starred repos
	Stars []Repo
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

	Projects, err = LoadProjects()
	if err != nil {
		panic(err)
	}

	go func() {
		token, _ := Asset("github.token")
		ticker := time.NewTicker(time.Minute * 30)
		for ; true; <-ticker.C {
			repos, stars, err := GetReposAndStars("razzie", string(token))
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

	http.ListenAndServe("localhost:8080", nil)
}
