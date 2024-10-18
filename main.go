package main

import (
	"flag"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP listener address")
	flag.Parse()

	content := LoadContent()
	projectViews := make(map[string]Project)
	for _, p := range content.Projects {
		projectViews[p.ID] = p
	}
	tagViews := make(map[string]map[string]any)
	for _, tag := range content.ProjectTags {
		tagViews[tag] = map[string]any{
			"Projects": filterProjectsByTag(content.Projects, tag),
			"Tags":     content.ProjectTags,
			"Tag":      tag,
		}
	}

	navPages := []Page{
		{
			ID:       "hello",
			Title:    "Hello",
			Template: "columns",
			Data:     content.Hello.ToHTML(),
		},
		{
			ID:       "skills",
			Title:    "Skills",
			Template: "columns",
			Data:     content.Skills.ToHTML(),
		},
		{
			ID:       "experience",
			Title:    "Experience",
			Template: "timeline",
			Data:     content.Experience.ToHTML(),
		},
		{
			ID:       "projects",
			Title:    "Projects",
			Template: "projects",
			Data: map[string]any{
				"Projects": content.Projects,
				"Tags":     content.ProjectTags,
			},
		},
	}

	var mux http.ServeMux

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusSeeOther)
	})

	mux.HandleFunc("GET /robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/robots.txt", http.StatusMovedPermanently)
	})

	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[1:]
		serveAsset(w, filename)
	})

	mux.HandleFunc("GET /static/projects/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[8:]
		serveAsset(w, filename)
	})

	for _, page := range navPages {
		mux.HandleFunc("GET /"+page.ID, func(w http.ResponseWriter, r *http.Request) {
			page.Render(w, navPages)
		})
	}

	mux.HandleFunc("GET /projects/tag/{tag}", func(w http.ResponseWriter, r *http.Request) {
		tag := r.PathValue("tag")
		view := tagViews[tag]
		if view == nil {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		Page{
			Title:    "Projects (" + tag + ")",
			ID:       "projects",
			Template: "projects",
			Data:     view,
		}.Render(w, navPages)
	})

	mux.HandleFunc("GET /projects/id/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		p, ok := projectViews[id]
		if !ok {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		Page{
			Title:    p.Name,
			ID:       p.ID,
			Template: "project",
			Data:     p,
		}.Render(w, navPages)
	})

	http.ListenAndServe(*addr, GzipMiddleware(&mux))
}
