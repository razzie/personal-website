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
