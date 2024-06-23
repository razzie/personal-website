package main

import (
	"bytes"
	"embed"
	"html/template"
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
	Base   string
	PageID string
	Data   any
}

/*func getBase(r *http.Request) string {
	slashes := strings.Count(r.URL.Path, "/")
	if slashes > 1 {
		return strings.Repeat("../", slashes-1)
	}
	return "/"
}*/

func main() {
	var t *template.Template
	t = template.Must(template.New("").Funcs(map[string]interface{}{
		"CallTemplate": func(name string, data interface{}) (ret template.HTML, err error) {
			var buf bytes.Buffer
			err = t.ExecuteTemplate(&buf, name, data)
			ret = template.HTML(buf.String())
			return
		},
		"mod": func(a, b int) int { return a % b },
		"add": func(a, b int) int { return a + b },
	}).ParseFS(assets, "templates/*.html"))

	projects, _ := LoadProjects()

	navPages := []Page{
		{ID: "hello", Name: "Hello"},
		{ID: "skills", Name: "Skills"},
		{ID: "experience", Name: "Experience"},
		{ID: "projects", Name: "Projects", Data: projects},
	}

	var r http.ServeMux
	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/hello", http.StatusSeeOther)
	})
	r.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[1:]
		ext := filepath.Ext(filename)
		f, err := assets.ReadFile(filename)
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", mime.TypeByExtension(ext))
		w.WriteHeader(http.StatusOK)
		w.Write(f)
	})
	r.HandleFunc("GET /static/projects/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[8:]
		ext := filepath.Ext(filename)
		f, err := assets.ReadFile(filename)
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", mime.TypeByExtension(ext))
		w.WriteHeader(http.StatusOK)
		w.Write(f)
	})
	for _, page := range navPages {
		r.HandleFunc("GET /"+page.ID, func(w http.ResponseWriter, r *http.Request) {
			view := View{
				Nav:    navPages,
				Title:  page.Name,
				PageID: page.ID,
				Data:   page.Data,
			}
			w.Header().Add("Content-Type", "text/html")
			if err := t.ExecuteTemplate(w, "layout", view); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}
	r.HandleFunc("GET /projects/tag/{tag}", func(w http.ResponseWriter, r *http.Request) {
		tag := r.PathValue("tag")
		if len(tag) == 0 {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		taggedProjects := filterProjectsByTag(projects, tag)
		if len(taggedProjects) == 0 {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		view := View{
			Nav:    navPages,
			Title:  "Projects (" + tag + ")",
			PageID: "projects",
			Data:   taggedProjects,
		}
		if err := t.ExecuteTemplate(w, "layout", view); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	r.HandleFunc("GET /projects/id/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if len(id) == 0 {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		var proj *Project
		for _, p := range projects {
			if p.ID == id {
				proj = &p
				break
			}
		}
		if proj == nil {
			http.Redirect(w, r, "/projects", http.StatusSeeOther)
			return
		}
		view := View{
			Nav:    navPages,
			Title:  proj.Name,
			PageID: "project",
			Data:   proj,
		}
		if err := t.ExecuteTemplate(w, "layout", view); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", &r)
}
