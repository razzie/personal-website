package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/razzie/gorzsony.com/data"
)

var (
	// Projects contains my hobby projects
	Projects []data.Project
	// Repos contains my github owned repos
	Repos []data.Repo
	// Stars contains my github starred repos
	Stars []data.Repo
)

func main() {
	index, err := data.Asset("index.html")
	if err != nil {
		panic(err)
	}

	tmpl := template.New("index")
	_, err = tmpl.Parse(string(index))
	if err != nil {
		panic(err)
	}

	Projects, err = data.LoadProjects()
	if err != nil {
		panic(err)
	}

	go func() {
		token, _ := data.Asset("github.token")
		ticker := time.NewTicker(time.Minute * 30)
		for ; true; <-ticker.C {
			repos, stars, err := data.GetReposAndStars("razzie", string(token))
			if err != nil {
				fmt.Println("error:", err)
				continue
			}

			Repos, Stars = repos, stars
		}
	}()

	fs := http.FileServer(
		&assetfs.AssetFS{Asset: data.Asset, AssetDir: data.AssetDir, AssetInfo: nil, Prefix: ""})

	http.Handle("/css/", fs)
	http.Handle("/img/", fs)
	http.Handle("/js/", fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, data.NewView(Projects, Repos, Stars))
	})
	http.HandleFunc("/tag/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tag := r.URL.Path[5:]
		tmpl.Execute(w, data.NewTagView(Projects, Repos, tag))
	})

	http.ListenAndServe("localhost:8080", nil)
}
