package main

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"mime"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

//go:embed content.yaml projects/*.yaml templates/*.html
var contentFS embed.FS

//go:embed static/* projects/*.webp
var assetFS embed.FS

var t *template.Template

func init() {
	t = template.Must(template.New("").Funcs(map[string]interface{}{
		"CallTemplate": func(name string, data interface{}) (ret template.HTML, err error) {
			var buf bytes.Buffer
			err = t.ExecuteTemplate(&buf, name, data)
			ret = template.HTML(buf.String())
			return
		},
	}).ParseFS(contentFS, "templates/*.html"))
}

type Content struct {
	Hello       Markdown      `yaml:"hello"`
	Skills      Markdown      `yaml:"skills"`
	Experience  MarkdownSlice `yaml:"experience"`
	Projects    []Project     `yaml:"-"`
	ProjectTags []string      `yaml:"-"`
}

type Page struct {
	ID       string
	Title    string
	Template string
	Data     any
}

func (page Page) Render(w http.ResponseWriter, navPages []Page) {
	const maxAge = 24 * 60 * 60
	expires := time.Now().Add(time.Duration(maxAge) * time.Second).Format(http.TimeFormat)
	view := map[string]any{
		"Nav":      navPages,
		"Title":    page.Title,
		"PageID":   page.ID,
		"Template": page.Template,
		"Data":     page.Data,
	}
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Cache-Control", "public, max-age="+strconv.Itoa(maxAge))
	w.Header().Add("Expires", expires)
	if err := t.ExecuteTemplate(w, "layout", view); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Project struct {
	ID          string   `yaml:"-"`
	Name        string   `yaml:"name"`
	Year        string   `yaml:"year"`
	Tags        []string `yaml:"tags"`
	Description Markdown `yaml:"description"`
	LinkGroups  []struct {
		Name  string `yaml:"name"`
		Links []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		}
	} `yaml:"linkGroups"`
}

func LoadContent() (content Content) {
	contentRaw, err := contentFS.ReadFile("content.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(contentRaw, &content); err != nil {
		panic(err)
	}
	content.Projects, content.ProjectTags = loadProjects()
	return
}

func (p *Project) containsTag(tag string) bool {
	for _, ptag := range p.Tags {
		if ptag == tag {
			return true
		}
	}
	return false
}

func orderProjectsByYear(projects []Project) {
	getYear := func(i int) int {
		years := strings.Split(projects[i].Year, "-")
		year, _ := strconv.Atoi(strings.TrimSpace(years[len(years)-1]))
		return year
	}
	sort.SliceStable(projects, func(i, j int) bool {
		return getYear(i) > getYear(j)
	})
}

func filterProjectsByTag(projects []Project, tag string) []Project {
	taggedProjects := make([]Project, 0, len(projects))
	for _, p := range projects {
		if p.containsTag(tag) {
			taggedProjects = append(taggedProjects, p)
		}
	}
	return taggedProjects
}

func loadProjects() (projects []Project, tags []string) {
	tagMap := make(map[string]struct{})

	if err := fs.WalkDir(contentFS, "projects", func(path string, entry fs.DirEntry, err error) error {
		name := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(name, ".yaml") {
			pYaml, err := contentFS.ReadFile(path)
			if err != nil {
				slog.Error("failed to open project yaml", slog.Any("err", err), slog.String("name", name))
				return nil
			}
			var p Project
			if err := yaml.Unmarshal(pYaml, &p); err != nil {
				slog.Error("failed to unmarshal project yaml", slog.Any("err", err), slog.String("name", name))
				return nil
			}
			p.ID = name[:len(name)-5]
			for _, tag := range p.Tags {
				tagMap[tag] = struct{}{}
			}
			sort.Strings(p.Tags)
			projects = append(projects, p)
		}
		return nil
	}); err != nil {
		slog.Error("failed to walk projects", slog.Any("err", err))
		return
	}

	for tag := range tagMap {
		tags = append(tags, tag)
	}
	orderProjectsByYear(projects)
	sort.Strings(tags)
	return
}

func serveAsset(w http.ResponseWriter, filename string) {
	ext := filepath.Ext(filename)
	f, err := assetFS.Open(filename)
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
