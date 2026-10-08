package internal

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ExperienceEntry struct {
	Content    Markdown `yaml:"content"`
	Sabbatical bool     `yaml:"sabbatical"`
}

type FeaturedProject struct {
	ID       string `yaml:"id"`
	Name     string `yaml:"name"`
	Category string `yaml:"category"`
	Summary  string `yaml:"summary"`
}

type HelloContent struct {
	Name               string            `yaml:"name"`
	Alias              string            `yaml:"alias"`
	Headline           []string          `yaml:"headline"`
	Summary            Markdown          `yaml:"summary"`
	About              Markdown          `yaml:"about"`
	FeaturedTitle      string            `yaml:"featuredTitle"`
	FeaturedProjects   []FeaturedProject `yaml:"featuredProjects"`
	CollaborationTitle string            `yaml:"collaborationTitle"`
	Collaboration      Markdown          `yaml:"collaboration"`
	ContactLabel       string            `yaml:"contactLabel"`
	Availability       Markdown          `yaml:"availability"`
}

type Content struct {
	Hello       HelloContent      `yaml:"hello"`
	Skills      Markdown          `yaml:"skills"`
	Experience  []ExperienceEntry `yaml:"experience"`
	Projects    []Project         `yaml:"-"`
	ProjectTags []Tag             `yaml:"-"`
}

func LoadContent(dir string) (content Content) {
	contentRaw, err := os.ReadFile(filepath.Join(dir, "content.yaml"))
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(contentRaw, &content); err != nil {
		panic(err)
	}
	content.Projects, content.ProjectTags = loadProjects(filepath.Join(dir, "projects"))
	return
}
