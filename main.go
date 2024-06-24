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
)

//go:embed static/* projects/*.png templates/*.html
var assets embed.FS

type Page struct {
	ID   string
	Name string
	Data any
}

type View struct {
	Nav    []Page
	Title  string
	PageID string
	Data   any
}

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

	projects, _ := LoadProjects()

	navPages := []Page{
		{ID: "hello", Name: "Hello"},
		{ID: "skills", Name: "Skills"},
		{ID: "experience", Name: "Experience"},
		{ID: "projects", Name: "Projects", Data: projects},
	}

	render := func(w http.ResponseWriter, title, pageID string, data any) {
		view := View{
			Nav:    navPages,
			Title:  title,
			PageID: pageID,
			Data:   data,
		}
		w.Header().Add("Content-Type", "text/html")
		if err := t.ExecuteTemplate(w, "layout", view); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusSeeOther)
	})

	http.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[1:]
		serveAsset(w, filename)
	})

	http.HandleFunc("GET /static/projects/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[8:]
		serveAsset(w, filename)
	})

	for _, page := range navPages {
		http.HandleFunc("GET /"+page.ID, func(w http.ResponseWriter, r *http.Request) {
			render(w, page.Name, page.ID, page.Data)
		})
	}

	http.HandleFunc("GET /projects/tag/{tag}", func(w http.ResponseWriter, r *http.Request) {
		tag := r.PathValue("tag")
		taggedProjects := filterProjectsByTag(projects, tag)
		if len(taggedProjects) == 0 {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		title := "Projects (" + tag + ")"
		render(w, title, "projects", taggedProjects)
	})

	http.HandleFunc("GET /projects/id/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idx := findProjectByID(projects, id)
		if idx < 0 {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		render(w, projects[idx].Name, "project", projects[idx])
	})

	http.ListenAndServe(*addr, nil)
}

func serveAsset(w http.ResponseWriter, filename string) {
	ext := filepath.Ext(filename)
	f, err := assets.Open(filename)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", mime.TypeByExtension(ext))
	w.WriteHeader(http.StatusOK)
	io.Copy(w, f)
}
