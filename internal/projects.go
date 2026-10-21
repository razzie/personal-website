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
	ID               string             `yaml:"-"`
	Name             string             `yaml:"name"`
	Year             string             `yaml:"year"`
	Tags             []Tag              `yaml:"tags"`
	Description      Markdown           `yaml:"description"`
	LinkGroups       []ProjectLinkGroup `yaml:"linkGroups"`
	DetailsLinkLabel string             `yaml:"detailsLabel"`
}

type ProjectLinkGroup struct {
	Name  string        `yaml:"name"`
	Links []ProjectLink `yaml:"links"`
}

type ProjectLink struct {
	Name      string `yaml:"name"`
	URL       string `yaml:"url"`
	CardLabel string `yaml:"cardLabel"`
}

func (p Project) DisplayYear() string {
	parts := strings.Split(p.Year, "-")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return strings.Join(parts, "–")
}

// CardLinks returns up to two explicitly labeled links in content order.
// Unlabeled links remain in the detail view only.
func (p Project) CardLinks() []ProjectLink {
	var links []ProjectLink
	for _, group := range p.LinkGroups {
		for _, link := range group.Links {
			if link.CardLabel == "" {
				continue
			}
			links = append(links, ProjectLink{Name: link.CardLabel, URL: link.URL})
			if len(links) == 2 {
				return links
			}
		}
	}
	return links
}

func (p Project) DetailsLabel() string {
	if p.DetailsLinkLabel != "" {
		return p.DetailsLinkLabel
	}
	return "Details"
}

func (p *Project) containsTag(tag Tag) bool {
	return slices.Contains(p.Tags, tag)
}

func (p Project) startingYear() int {
	start, _, _ := strings.Cut(p.Year, "-")
	year, _ := strconv.Atoi(strings.TrimSpace(start))
	return year
}

func orderProjectsByYear(projects []Project) {
	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].startingYear() > projects[j].startingYear()
	})
}

type ProjectYearGroup struct {
	Year     int
	Projects []Project
}

// GroupProjectsByYear groups projects by their starting year, newest first.
// Projects within the same year retain their original order.
func GroupProjectsByYear(projects []Project) []ProjectYearGroup {
	projects = slices.Clone(projects)
	orderProjectsByYear(projects)

	var groups []ProjectYearGroup
	for _, project := range projects {
		year := project.startingYear()
		if len(groups) == 0 || groups[len(groups)-1].Year != year {
			groups = append(groups, ProjectYearGroup{Year: year})
		}
		group := &groups[len(groups)-1]
		group.Projects = append(group.Projects, project)
	}
	return groups
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
