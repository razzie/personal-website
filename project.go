package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v2"
)

//go:embed projects/*.yaml
var projFS embed.FS

type Project struct {
	ID              string        `yaml:"-"`
	Name            string        `yaml:"name"`
	Year            string        `yaml:"year"`
	Tags            []string      `yaml:"tags"`
	Description     string        `yaml:"description"`
	DescriptionHTML template.HTML `yaml:"-"`
	LinkGroups      []struct {
		Name  string `yaml:"name"`
		Links []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		}
	} `yaml:"linkGroups"`
}

func (p *Project) containsTag(tag string) bool {
	for _, ptag := range p.Tags {
		if ptag == tag {
			return true
		}
	}
	return false
}

func convertMD(input string, output *template.HTML) {
	htmlRenderer := html.NewRenderer(html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank | html.UseXHTML,
	})
	mdParser := parser.NewWithExtensions(parser.CommonExtensions | parser.NoEmptyLineBeforeBlock)
	html := markdown.ToHTML([]byte(input), mdParser, htmlRenderer)
	*output = template.HTML(html)
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

func LoadProjects() (projects []Project, tags []string) {
	tagMap := make(map[string]struct{})

	if err := fs.WalkDir(projFS, "projects", func(path string, entry fs.DirEntry, err error) error {
		name := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(name, ".yaml") {
			pYaml, err := projFS.ReadFile(path)
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
			convertMD(p.Description, &p.DescriptionHTML)
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
