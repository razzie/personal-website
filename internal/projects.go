package internal

import (
	"io/fs"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Tag string

func (tag Tag) Escape() string {
	return url.QueryEscape(string(tag))
}

type Project struct {
	ID          string   `yaml:"-"`
	Name        string   `yaml:"name"`
	Year        string   `yaml:"year"`
	Tags        []Tag    `yaml:"tags"`
	Description Markdown `yaml:"description"`
	LinkGroups  []struct {
		Name  string `yaml:"name"`
		Links []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		}
	} `yaml:"linkGroups"`
}

func (p *Project) containsTag(tag Tag) bool {
	return slices.Contains(p.Tags, tag)
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

func FilterProjectsByTag(projects []Project, tag Tag) []Project {
	taggedProjects := make([]Project, 0, len(projects))
	for _, p := range projects {
		if p.containsTag(tag) {
			taggedProjects = append(taggedProjects, p)
		}
	}
	return taggedProjects
}

func loadProjects(dir string) (projects []Project, tags []Tag) {
	tagMap := make(map[Tag]struct{})

	if err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		name := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(name, ".yaml") {
			pYaml, err := os.ReadFile(path)
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
			slices.Sort(p.Tags)
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
	slices.Sort(tags)
	return
}
