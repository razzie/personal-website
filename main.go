package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var (
	Projects []Project
	Repos    []Repo
	Stars    []Repo
)

func main() {
	index, err := Asset("index.html")
	if err != nil {
		panic(err)
	}

	Projects, err = LoadProjects()
	if err != nil {
		panic(err)
	}

	go func() {
		ticker := time.NewTicker(time.Minute * 30)
		for ; true; <-ticker.C {
			repos, stars, err := GetReposAndStars("razzie", "b37e88bef5b39e88dab437fc49351aff1c29d853")
			if err != nil {
				fmt.Println("error:", err)
				continue
			}

			Repos, Stars = repos, stars
		}
	}()

	tmpl := template.New("index")
	tmpl.Parse(string(index))

	fs := assetFS()
	fs.Prefix = ""

	http.Handle("/css/", http.FileServer(fs))
	http.Handle("/img/", http.FileServer(fs))
	http.Handle("/js/", http.FileServer(fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, NewView(Projects, Repos, Stars))
	})

	http.ListenAndServe("localhost:8080", nil)
}
