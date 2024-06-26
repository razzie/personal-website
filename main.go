package main

import (
	"bytes"
	"embed"
	"flag"
	"html/template"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

//go:embed static/* projects/*.webp templates/*.html
var assets embed.FS

func main() {
	addr := flag.String("addr", ":8080", "HTTP listener address")
	flag.Parse()

	var t *template.Template
	t = template.Must(template.New("").Funcs(map[string]interface{}{
		"CallTemplate": func(name string, data interface{}) (ret template.HTML, err error) {
			var buf bytes.Buffer
			err = t.ExecuteTemplate(&buf, name, data)
			ret = template.HTML(buf.String())
			return
		},
	}).ParseFS(assets, "templates/*.html"))

	projects, tags := LoadProjects()
	projectViews := make(map[string]Project)
	for _, p := range projects {
		projectViews[p.ID] = p
	}
	tagViews := make(map[string]map[string]any)
	for _, tag := range tags {
		tagViews[tag] = map[string]any{
			"Projects": filterProjectsByTag(projects, tag),
			"Tags":     tags,
			"Tag":      tag,
		}
	}

	navPages := []struct {
		ID   string
		Name string
		Data any
	}{
		{
			ID:   "hello",
			Name: "Hello",
		},
		{
			ID:   "skills",
			Name: "Skills",
		},
		{
			ID:   "experience",
			Name: "Experience",
		},
		{
			ID:   "projects",
			Name: "Projects",
			Data: map[string]any{
				"Projects": projects,
				"Tags":     tags,
			},
		},
	}

	render := func(w http.ResponseWriter, title, pageID string, data any) {
		const maxAge = 24 * 60 * 60
		expires := time.Now().Add(time.Duration(maxAge) * time.Second).Format(http.TimeFormat)
		view := map[string]any{
			"Nav":    navPages,
			"Title":  title,
			"PageID": pageID,
			"Data":   data,
		}
		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("Cache-Control", "public, max-age="+strconv.Itoa(maxAge))
		w.Header().Add("Expires", expires)
		if err := t.ExecuteTemplate(w, "layout", view); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
			render(w, page.Name, page.ID, page.Data)
		})
	}

	mux.HandleFunc("GET /projects/tag/{tag}", func(w http.ResponseWriter, r *http.Request) {
		tag := r.PathValue("tag")
		view := tagViews[tag]
		if view == nil {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		title := "Projects (" + tag + ")"
		render(w, title, "projects", view)
	})

	mux.HandleFunc("GET /projects/id/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		p, ok := projectViews[id]
		if !ok {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		render(w, p.Name, "project", p)
	})

	http.ListenAndServe(*addr, GzipMiddleware(&mux))
}

func serveAsset(w http.ResponseWriter, filename string) {
	ext := filepath.Ext(filename)
	f, err := assets.Open(filename)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	const maxAge = 7 * 24 * 60 * 60
	expires := time.Now().Add(time.Duration(maxAge) * time.Second).Format(http.TimeFormat)
	w.Header().Add("Content-Type", mime.TypeByExtension(ext))
	w.Header().Add("Cache-Control", "public, max-age="+strconv.Itoa(maxAge))
	w.Header().Add("Expires", expires)
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, f)
}
