package internal

import (
	"bytes"
	"html/template"
	"io"
)

type Page struct {
	ID       string
	Title    string
	Template string
	Data     any
}

type TemplateRenderer func(w io.Writer, page Page) error

func LoadTemplateRenderer(navPages []Page) TemplateRenderer {
	var t *template.Template
	t = template.Must(template.New("").Funcs(map[string]interface{}{
		"CallTemplate": func(name string, data interface{}) (ret template.HTML, err error) {
			var buf bytes.Buffer
			err = t.ExecuteTemplate(&buf, name, data)
			ret = template.HTML(buf.String())
			return
		},
	}).ParseGlob("templates/*.html"))

	return func(w io.Writer, page Page) error {
		view := map[string]any{
			"Nav":      navPages,
			"Title":    page.Title,
			"PageID":   page.ID,
			"Template": page.Template,
			"Data":     page.Data,
		}
		return t.ExecuteTemplate(w, "layout", view)
	}
}
